package gin

import (
	"errors"
	"go/ast"
	"reflect"
	"strings"

	"github.com/ldez/tagliatelle/filedtype"

	"golang.org/x/tools/go/analysis"
)

const tagName = "binding"

type tags map[string]tagValue

type tagValue struct {
	value     string
	valueType valueType
}

func Lint(pass *analysis.Pass, tag *ast.BasicLit, fieldTypes []filedtype.FiledType, key, convName string) {
	if key != tagName {
		return
	}

	tagValues, ok := lookupTagValue(tag, key)
	if !ok {
		// skip when no struct tag for the key
		return
	}

	if _, ok := tagValues["-"]; ok {
		// skip when skipped :)
		return
	}

	for k, v := range tagValues {
		if err := converter(fieldTypes, k, v); err != nil {
			pass.Reportf(tag.Pos(), "%s(%s): %s", key, convName, err.Error())
		}
	}
}

func lookupTagValue(tag *ast.BasicLit, key string) (tags, bool) {
	raw := strings.Trim(tag.Value, "`")

	value, ok := reflect.StructTag(raw).Lookup(key)
	if !ok {
		return nil, false
	}

	values := strings.Split(value, ",")
	tagValues := make(tags, len(values))
	for _, v := range values {
		vsp := strings.Split(v, "=")
		if len(vsp) < 1 {
			continue
		}
		var tagK, tagV string
		tagK = vsp[0]
		if len(vsp) > 1 {
			tagV = vsp[1]
		}
		tagValues[tagK] = tagValue{
			value:     tagV,
			valueType: Parse(tagV),
		}
	}
	return tagValues, true
}

func converter(fieldTypes []filedtype.FiledType, key string, value tagValue) error {
	for _, fType := range fieldTypes {
		switch fType {
		case filedtype.Invalid:
			continue
		case filedtype.Bool:
		case filedtype.Int, filedtype.Int8, filedtype.Int16, filedtype.Int32, filedtype.Int64,
			filedtype.Uint, filedtype.Uint8, filedtype.Uint16, filedtype.Uint32, filedtype.Uint64,
			filedtype.Float32, filedtype.Float64:
			if err := number(key, value); err != nil {
				return err
			}
		case filedtype.Complex64, filedtype.Complex128:
			continue
		case filedtype.Array, filedtype.Slice:
		case filedtype.Map:
		case filedtype.Ptr:
			continue
		case filedtype.String:
		case filedtype.Time:

		case filedtype.Other:
			continue
		}
	}
	return nil
}

func number(key string, value tagValue) error {
	switch key {
	case "required":
		if value.valueType == valueTypeNONE {
			return nil
		}
		return errors.New("no value required")
	case "len", "min", "max", "eq", "ne", "lt", "lte", "gt", "gte":
		if value.valueType == valueTypeINT || value.valueType == valueTypeFLOAT {
			return nil
		}
		return errors.New("value must be a number")
	default:
		return errors.New("validation characters are not supported")
	}
}
