// Package tagliatelle a linter that handle struct tags.
package tagliatelle

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ettle/strcase"
	"github.com/golangci/modinfo"
	iradix "github.com/hashicorp/go-immutable-radix/v2"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Config the tagliatelle configuration.
type Config struct {
	Base
	Overrides []Overrides
}

// Overrides applies configuration overrides by package.
type Overrides struct {
	Base
	Package string
}

// Base shared configuration between rules.
type Base struct {
	Rules         map[string]string
	UseFieldName  bool
	IgnoredFields []string
	Ignore        bool
}

// New creates an analyzer.
func New(config Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "tagliatelle",
		Doc:  "Checks the struct tags.",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			if len(config.Rules) == 0 && len(config.Overrides) == 0 {
				return nil, nil
			}

			return run(pass, config)
		},
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
			modinfo.Analyzer,
		},
	}
}

func run(pass *analysis.Pass, config Config) (interface{}, error) {
	info, err := modinfo.FindModuleFromPass(pass)
	if err != nil {
		return nil, err
	}

	isp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("missing inspect analyser")
	}

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	r := createRadixTree(config, info)

	isp.Preorder(nodeFilter, func(n ast.Node) {
		node, ok := n.(*ast.StructType)
		if !ok {
			return
		}

		for _, field := range node.Fields.List {
			_, v, _ := r.Root().LongestPrefix([]byte(path.Join(pass.Pkg.Path(), filepath.Base(pass.Fset.File(node.Pos()).Name()))))
			if v.Ignore {
				continue
			}

			analyze(pass, v, node, field)
		}
	})

	return nil, nil
}

func analyze(pass *analysis.Pass, config Base, n *ast.StructType, field *ast.Field) {
	if n.Fields == nil || n.Fields.NumFields() < 1 {
		// skip empty structs
		return
	}

	if field.Tag == nil {
		// skip when no struct tag
		return
	}

	fieldName, err := getFieldName(field)
	if err != nil {
		pass.Reportf(n.Pos(), "unable to get field name: %v", err)
		return
	}

	if contains(config.IgnoredFields, fieldName) {
		return
	}

	for key, convName := range config.Rules {
		if convName == "" {
			continue
		}

		value, flags, ok := lookupTagValue(field.Tag, key)
		if !ok {
			// skip when no struct tag for the key
			continue
		}

		if value == "-" {
			// skip when skipped :)
			continue
		}

		// TODO(ldez): need to be rethink.
		// tagliatelle should try to remain neutral in terms of format.
		if key == "xml" && strings.ContainsAny(value, ">:") {
			// ignore XML names than contains path
			continue
		}

		// TODO(ldez): need to be rethink.
		// This is an exception because of a bug.
		// https://github.com/ldez/tagliatelle/issues/8
		// For now, tagliatelle should try to remain neutral in terms of format.
		if hasTagFlag(flags, "inline") {
			// skip for inline children (no name to lint)
			continue
		}

		if value == "" {
			value = fieldName
		}

		converter, err := getConverter(convName)
		if err != nil {
			pass.Reportf(n.Pos(), "%s(%s): %v", key, convName, err)
			continue
		}

		expected := value
		if config.UseFieldName {
			expected = fieldName
		}

		if value != converter(expected) {
			pass.Reportf(field.Tag.Pos(), "%s(%s): got '%s' want '%s'", key, convName, value, converter(expected))
		}
	}
}

func getFieldName(field *ast.Field) (string, error) {
	var name string
	for _, n := range field.Names {
		if n.Name != "" {
			name = n.Name
		}
	}

	if name != "" {
		return name, nil
	}

	return getTypeName(field.Type)
}

func getTypeName(exp ast.Expr) (string, error) {
	switch typ := exp.(type) {
	case *ast.Ident:
		return typ.Name, nil
	case *ast.StarExpr:
		return getTypeName(typ.X)
	case *ast.SelectorExpr:
		return getTypeName(typ.Sel)
	default:
		bytes, _ := json.Marshal(exp)
		return "", fmt.Errorf("unexpected error: type %T: %s", typ, string(bytes))
	}
}

func lookupTagValue(tag *ast.BasicLit, key string) (name string, flags []string, ok bool) {
	raw := strings.Trim(tag.Value, "`")

	value, ok := reflect.StructTag(raw).Lookup(key)
	if !ok {
		return value, nil, ok
	}

	values := strings.Split(value, ",")

	if len(values) < 1 {
		return "", nil, true
	}

	return values[0], values[1:], true
}

func hasTagFlag(flags []string, query string) bool {
	for _, flag := range flags {
		if flag == query {
			return true
		}
	}

	return false
}

func getConverter(c string) (func(s string) string, error) {
	switch c {
	case "camel":
		return strcase.ToCamel, nil
	case "pascal":
		return strcase.ToPascal, nil
	case "kebab":
		return strcase.ToKebab, nil
	case "snake":
		return strcase.ToSnake, nil
	case "goCamel":
		return strcase.ToGoCamel, nil
	case "goPascal":
		return strcase.ToGoPascal, nil
	case "goKebab":
		return strcase.ToGoKebab, nil
	case "goSnake":
		return strcase.ToGoSnake, nil
	case "header":
		return toHeader, nil
	case "upper":
		return strings.ToUpper, nil
	case "upperSnake":
		return strcase.ToSNAKE, nil
	case "lower":
		return strings.ToLower, nil
	default:
		return nil, fmt.Errorf("unsupported case: %s", c)
	}
}

func toHeader(s string) string {
	return strcase.ToCase(s, strcase.TitleCase, '-')
}

func createRadixTree(config Config, info modinfo.ModInfo) *iradix.Tree[Base] {
	r := iradix.New[Base]()

	defaultRule := Base{
		Rules:        copyMap(config.Rules),
		UseFieldName: config.UseFieldName,
		Ignore:       config.Ignore,
	}
	defaultRule.IgnoredFields = append(defaultRule.IgnoredFields, config.IgnoredFields...)

	r, _, _ = r.Insert([]byte(""), defaultRule)

	for _, override := range config.Overrides {
		c := Base{
			Rules:        copyMap(config.Rules),
			UseFieldName: override.UseFieldName,
			Ignore:       override.Ignore,
		}
		c.IgnoredFields = append(c.IgnoredFields, config.IgnoredFields...)
		c.IgnoredFields = append(c.IgnoredFields, override.IgnoredFields...)

		for k, v := range override.Rules {
			c.Rules[k] = v
		}

		r, _, _ = r.Insert([]byte(path.Join(info.Path, override.Package)), c)
	}

	return r
}

func copyMap[K, V comparable](m map[K]V) map[K]V {
	c := make(map[K]V)
	for k, v := range m {
		c[k] = v
	}
	return c
}

func contains[T comparable](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
