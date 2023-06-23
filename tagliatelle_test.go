package tagliatelle_test

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"testing"

	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	cfg := tagliatelle.Config{
		Base: tagliatelle.Base{
			Rules: map[string]string{
				"json":         "camel",
				"yaml":         "camel",
				"xml":          "camel",
				"bson":         "camel",
				"avro":         "snake",
				"mapstructure": "kebab",
				"header":       "header",
				"envconfig":    "upperSnake",
				"env":          "upperSnake",
			},
			UseFieldName: true,
		},
		Overrides: []tagliatelle.Overrides{
			{
				Package: "a/b/c",
				Base: tagliatelle.Base{
					Rules: map[string]string{
						"json": "upperSnake",
						"yaml": "upperSnake",
					},
					UseFieldName: false,
				},
			},
		},
	}

	runWithSuggestedFixes(t, tagliatelle.New(cfg), "a", "a")
}

func runWithSuggestedFixes(t *testing.T, a *analysis.Analyzer, dir string, patterns ...string) []*analysistest.Result {
	t.Helper()

	testdata := analysistest.TestData()

	// NOTE: analysistest does not yet support modules;
	// see https://github.com/golang/go/issues/37054 for details.

	err := os.Chdir(filepath.Join(testdata, "src", path.Join(dir)))
	if err != nil {
		t.Fatal(err)
	}

	output, err := exec.Command("go", "mod", "vendor").CombinedOutput()
	if err != nil {
		t.Log(string(output))
		t.Fatal(err)
	}

	return analysistest.RunWithSuggestedFixes(t, testdata, a, patterns...)
}
