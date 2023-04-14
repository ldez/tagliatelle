package main

import (
	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	cfg := tagliatelle.Config{
		Rules: map[string]string{
			"json":      "camel",
			"yaml":      "camel",
			"xml":       "camel",
			"bson":      "camel",
			"avro":      "snake",
			"header":    "header",
			"envconfig": "upperSnake",
		},
		UseFieldName: true,
	}

	singlechecker.Main(tagliatelle.New(cfg))
}
