package filedtype

type FiledType int

const (
	Invalid FiledType = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Complex64
	Complex128
	Array
	Map
	Ptr
	Slice
	String
	Time
	Other
)

func Parse(typeName string) FiledType {
	switch typeName {
	case "":
		return Invalid
	case "bool":
		return Bool
	case "int":
		return Int
	case "int8":
		return Int8
	case "int16":
		return Int16
	case "int32":
		return Int32
	case "int64":
		return Int64
	case "uint":
		return Uint
	case "uint8":
		return Uint8
	case "uint16":
		return Uint16
	case "uint32":
		return Uint32
	case "uint64":
		return Uint64
	case "float32":
		return Float32
	case "float64":
		return Float64
	case "complex64":
		return Complex64
	case "complex128":
		return Complex128
	case "array":
		return Array
	case "map":
		return Map
	case "ptr":
		return Ptr
	case "slice":
		return Slice
	case "string":
		return String
	case "time":
		return Time
	default:
		return Other
	}
}
