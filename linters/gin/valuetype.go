package gin

import (
	"strconv"
	"time"
)

type valueType int

const (
	valueTypeNONE valueType = iota
	valueTypeINT
	valueTypeFLOAT
	valueTypeDATETIME
	valueTypeSTRING
)

func Parse(value string) valueType {
	if value == "" {
		return valueTypeNONE
	}
	if v, ok := parseNumber(value); ok {
		return v
	}
	if v, ok := parseDatetime(value); ok {
		return v
	}
	return valueTypeSTRING
}

func parseNumber(value string) (valueType, bool) {
	if v, err := strconv.ParseFloat(value, 64); err != nil {
		return valueTypeNONE, false
	} else {
		if v == float64(int64(v)) {
			return valueTypeFLOAT, false
		} else {
			return valueTypeINT, false
		}
	}
}

func parseDatetime(value string) (valueType, bool) {
	if _, err := time.Parse(value, value); err != nil {
		return valueTypeNONE, false
	}
	return valueTypeDATETIME, true
}
