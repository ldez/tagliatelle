package main

import (
	"github.com/ldez/tagliatelle"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	rules := map[string]string{
		"json": "camel",
		"yaml": "camel",
		"xml":  "camel",
		"bson": "camel",
		"avro": "snake",
	}

	singlechecker.Main(tagliatelle.New(rules))
}
