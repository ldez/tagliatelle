package tagliatelle_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc     string
		dir      string
		patterns []string
		cfg      tagliatelle.Config
	}{
		{
			desc:     "simple",
			dir:      "one",
			patterns: []string{"one"},
			cfg: tagliatelle.Config{
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
			},
		},
		{
			desc:     "with non-applicable overrides",
			dir:      "one",
			patterns: []string{"one/..."},
			cfg: tagliatelle.Config{
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
						Package: "one/b/c",
						Base: tagliatelle.Base{
							Rules: map[string]string{
								"json": "upperSnake",
								"yaml": "upperSnake",
							},
							UseFieldName: false,
						},
					},
				},
			},
		},
		{
			desc:     "with applicable overrides",
			dir:      "two",
			patterns: []string{"two/..."},
			cfg: tagliatelle.Config{
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
						Package: "two/b",
						Base: tagliatelle.Base{
							Rules: map[string]string{
								"json": "upperSnake",
								"yaml": "upperSnake",
							},
							UseFieldName: false,
						},
					},
				},
			},
		},
		{
			desc:     "ignore",
			dir:      "three",
			patterns: []string{"three/..."},
			cfg: tagliatelle.Config{
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
						Package: "three/b",
						Base: tagliatelle.Base{
							Ignore: true,
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			runWithSuggestedFixes(t, tagliatelle.New(test.cfg), test.dir, test.patterns...)
		})
	}
}

func runWithSuggestedFixes(t *testing.T, a *analysis.Analyzer, dir string, patterns ...string) []*analysistest.Result {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = os.Chdir(wd) }()

	testdata := analysistest.TestData()

	// NOTE: analysistest does not yet support modules;
	// see https://github.com/golang/go/issues/37054 for details.

	err = os.Chdir(filepath.Join(testdata, "src", filepath.FromSlash(dir)))
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
