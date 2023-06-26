package main

import (
	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	cfg := tagliatelle.Config{
		Base: tagliatelle.Base{
			Rules: map[string]string{
				"json":      "camel",
				"yaml":      "camel",
				"xml":       "camel",
				"bson":      "camel",
				"avro":      "snake",
				"header":    "header",
				"envconfig": "upperSnake",
				"env":       "upperSnake",
			},
			UseFieldName: true,
		},
	}

	singlechecker.Main(tagliatelle.New(cfg))
}
