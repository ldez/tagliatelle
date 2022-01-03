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
		},
		UseFieldName: true,
	}

	analysistest.RunWithSuggestedFixes(t, testdata, tagliatelle.New(cfg), "a")
}

func TestAnalyzerGin(t *testing.T) {
	testdata := analysistest.TestData()
	t.Run("Gin-camel", func(t *testing.T) {
		cfg := tagliatelle.Config{
			Rules: map[string]string{
				"binding": "camel",
			},
			UseFieldName: true,
		}
		analysistest.RunWithSuggestedFixes(t, testdata, tagliatelle.New(cfg), "gin_camel")
	})
	t.Run("Gin-snake", func(t *testing.T) {
		cfg := tagliatelle.Config{
			Rules: map[string]string{
				"binding": "snake",
			},
			UseFieldName: true,
		}
		analysistest.RunWithSuggestedFixes(t, testdata, tagliatelle.New(cfg), "gin_snake")
	})
	t.Run("Gin-kebab", func(t *testing.T) {
		cfg := tagliatelle.Config{
			Rules: map[string]string{
				"binding": "kebab",
			},
			UseFieldName: true,
		}
		analysistest.RunWithSuggestedFixes(t, testdata, tagliatelle.New(cfg), "gin_kebab")
	})
}
