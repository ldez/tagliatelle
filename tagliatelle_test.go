package tagliatelle_test

import (
	"testing"

	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	cfg := tagliatelle.Config{
		Rules: map[string]string{
			"json":         "camel",
			"yaml":         "camel",
			"xml":          "camel",
			"bson":         "camel",
			"avro":         "snake",
			"mapstructure": "kebab",
			"header":       "header",
			"envconfig":    "upperSnake",
		},
		UseFieldName: true,
	}

	analysistest.RunWithSuggestedFixes(t, testdata, tagliatelle.New(cfg), "a")
}
