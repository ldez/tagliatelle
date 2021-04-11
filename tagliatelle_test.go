package tagliatelle_test

import (
	"testing"

	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	rules := map[string]string{
		"json": "camel",
		"yaml": "camel",
		"xml":  "camel",
		"bson": "camel",
		"avro": "snake",
	}

	analysistest.RunWithSuggestedFixes(t, testdata, tagliatelle.New(rules), "a")
}
