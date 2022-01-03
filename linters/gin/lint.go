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

	if err := converter(fieldTypes, tagValues); err != nil {
		pass.Reportf(tag.Pos(), "%s(%s): %s", key, convName, err.Error())
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

func converter(fieldTypes []filedtype.FiledType, tagValues tags) error {
	for _, fType := range fieldTypes {
		switch fType {
		case filedtype.Invalid:
			continue
		case filedtype.Bool:
		case filedtype.Int, filedtype.Int8, filedtype.Int16, filedtype.Int32, filedtype.Int64,
			filedtype.Uint, filedtype.Uint8, filedtype.Uint16, filedtype.Uint32, filedtype.Uint64,
			filedtype.Float32, filedtype.Float64:
			if err := number(tagValues); err != nil {
				return err
			}
		case filedtype.Complex64, filedtype.Complex128:
			continue
		case filedtype.Array, filedtype.Slice, filedtype.Map:
			if err := sliceOrArrayOrMap(tagValues); err != nil {
				return err
			}
		case filedtype.Ptr:
			continue
		case filedtype.String:
			if err := str(tagValues); err != nil {
				return err
			}
		case filedtype.Time:
			if err := datetime(tagValues); err != nil {
				return err
			}
		case filedtype.Other:
			continue
		}
	}
	return nil
}

func common(tagValues tags) error {
	for key, value := range tagValues {
		switch key {
		case "required", "structonly", "nostructlevel", "omitempty",
			"alpha", "alphanum", "alphanumunicode", "alphaunicode",
			"numeric", "number", "boolean":
			if value.valueType == valueTypeNONE {
				return nil
			}
			return errors.New("no value required")
		case "eqfield", "eqcsfield", "necsfield", "gtcsfield", "gtecsfield", "ltcsfield", "ltecsfield",
			"nefield", "gtfield", "gtefield", "ltfield", "ltefield":
			if value.valueType == valueTypeSTRING {
				continue
			}
			return errors.New("value must be a string")
		case "required_if", "required_unless",
			"required_with", "required_with_all", "required_without", "required_without_all",
			"excluded_with", "excluded_with_all", "excluded_without", "excluded_without_all":
			// TODO
		}
	}
	return nil
}

func number(tagValues tags) error {
	if err := common(tagValues); err != nil {
		return err
	}
	for key, value := range tagValues {
		switch key {
		case "len", "min", "max", "eq", "ne", "lt", "lte", "gt", "gte":
			if value.valueType == valueTypeINT || value.valueType == valueTypeFLOAT {
				continue
			}
			return errors.New("value must be a number")
		default:
			return errors.New("validation characters are not supported")
		}
	}
	return nil
}

func str(tagValues tags) error {
	if err := common(tagValues); err != nil {
		return err
	}
	for key, value := range tagValues {
		switch key {
		case "len", "min", "max", "eq", "ne", "lt", "lte", "gt", "gte":
			if value.valueType == valueTypeINT {
				continue
			}
			return errors.New("value must be a integer")
		case
			// color
			"hexadecimal", "hexcolor", "rgb", "rgba", "hsl", "hsla", "iscolor",
			// email
			"email",
			// url
			"url", "uri", "base64", "base64url", "url_encoded", "urn_rfc2141",
			// International Standard Book Number
			"isbn", "isbn10", "isbn13",
			// Universally Unique Identifier UUID
			"uuid", "uuid3", "uuid4", "uuid5", "ulid",
			// character
			"ascii", "printascii", "multibyte",
			// phone number
			"e164",
			// country code
			"iso3166_1_alpha2", "iso3166_1_alpha3", "iso3166_1_alpha_numeric", "country_code",
			// datauri
			"datauri",
			// Latitude Longitude
			"latitude", "longitude",
			// Social Security Number SSN
			"ssn",
			// Timezone
			"timezone",
			// network
			"ip", "ipv4", "ipv6", "ip_addr", "ip4_addr", "ip6_addr",
			"cidr", "cidrv4", "cidrv6",
			"tcp_addr", "tcp4_addr", "tcp6_addr",
			"udp_addr", "udp4_addr", "udp6_addr",
			"unix_addr", "mac", "fqdn",
			"hostname", "hostname_port", "hostname_rfc1123",
			// bcp
			"bic", "bcp47_language_tag",
			// btc
			"btc_addr", "btc_addr_bech32",
			// html tags
			"html", "html_encoded":
			if value.valueType == valueTypeNONE {
				continue
			}
			return errors.New("no value required")
		case "json", "jwt", "lowercase", "uppercase":
			if value.valueType == valueTypeSTRING {
				continue
			}
			return errors.New("field must be a string")
		case "datetime":
			if value.valueType == valueTypeDATETIME {
				continue
			}
			return errors.New("value must be a datetime format")
		case "postcode_iso3166_alpha2", "postcode_iso3166_alpha2_field":
			// Postcode
			if value.valueType == valueTypeSTRING {
				continue
			}
			return errors.New("field must be a string")
		default:
			return errors.New("validation characters are not supported")
		}
	}
	return nil
}

func sliceOrArrayOrMap(tagValues tags) error {
	if err := common(tagValues); err != nil {
		return err
	}
	for key, value := range tagValues {
		if key == "dive" {
			return nil
		}
		switch key {
		case "len", "min", "max", "eq", "ne", "lt", "lte", "gt", "gte":
			if value.valueType == valueTypeINT {
				continue
			}
			return errors.New("value must be a integer")
		case "unique":
			if value.valueType == valueTypeNONE {
				continue
			}
			return errors.New("no value required")
		default:
			return errors.New("validation characters are not supported")
		}
	}
	return nil
}

func datetime(tagValues tags) error {
	for key, value := range tagValues {
		switch key {
		case "len", "min", "max", "eq", "ne", "lt", "lte", "gt", "gte":
			if value.valueType == valueTypeNONE {
				continue
			}
			return errors.New("no value required")
		}
	}
	return nil
}
