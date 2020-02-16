package v8

// #include <v8capi.h>
import "C"

import "fmt"

type Value struct {
	data C.struct_v8_value
}

func (val Value) Dispose() {
	C.v8_delete_value(&val.data)
}

func (val Value) typeAsString() string {
	switch C.v8_get_value_type(val.data) {
	case C.v8_undefined:
		return "undefined"
	case C.v8_boolean:
		return "boolean"
	case C.v8_null:
		return "null"
	case C.v8_number:
		return "number"
	case C.v8_string:
		return "string"
	case C.v8_big_int:
		return "big_int"
	case C.v8_symbol:
		return "symbol"
	case C.v8_object:
		return "object"
	case C.v8_array:
		return "array"
	case C.v8_set:
		return "set"
	case C.v8_map:
		return "map"
	case C.v8_function:
		return "function"
	case C.v8_date:
		return "date"
	}
	return "[unknown type]"
}

func makeUndefined() Value {
	return Value{data: C.v8_new_undefined()}
}

func (val Value) IsUndefined() bool {
	return bool(C.v8_is_undefined(val.data))
}

func (val Value) IsBoolean() bool {
	return bool(C.v8_is_boolean(val.data))
}

func (val Value) IsNull() bool {
	return bool(C.v8_is_null(val.data))
}

func (val Value) IsNumber() bool {
	return bool(C.v8_is_number(val.data))
}

func (val Value) IsDouble() bool {
	return bool(C.v8_is_number(val.data))
}

func (val Value) IsInteger() bool {
	return bool(C.v8_is_integer(val.data))
}

func (val Value) IsString() bool {
	return bool(C.v8_is_string(val.data))
}

func (val Value) IsObject() bool {
	return bool(C.v8_is_object(val.data))
}

func (val Value) IsArray() bool {
	return bool(C.v8_is_array(val.data))
}

func (val Value) IsSet() bool {
	return bool(C.v8_is_set(val.data))
}

func (val Value) IsMap() bool {
	return bool(C.v8_is_map(val.data))
}

func (val Value) ToBool() (bool, error) {
	if val.IsBoolean() {
		return bool(C.v8_to_bool(val.data)), nil
	}
	return false, fmt.Errorf("Can't convert %s to bool", val.typeAsString())
}

func (val Value) ToInt() (int64, error) {
	if val.IsInteger() {
		return int64(C.v8_to_int64(val.data)), nil
	}
	return 0, fmt.Errorf("Can't convert %s to int64", val.typeAsString())
}

func (val Value) ToUint() (uint64, error) {
	if val.IsInteger() {
		i := int64(C.v8_to_int64(val.data))
		if i < 0 {
			return 0, fmt.Errorf("Can't cast negative value %d to unsigned value", i)
		}
		return uint64(i), nil
	}
	return 0, fmt.Errorf("Can't convert %s to uint64", val.typeAsString())
}

func (val Value) ToFloat() (float64, error) {
	if val.IsDouble() {
		return float64(C.v8_to_double(val.data)), nil
	}
	return 0., fmt.Errorf("Can't convert %s to float64", val.typeAsString())
}

func (val Value) ToString() (string, error) {
	if val.IsString() {
		str := C.v8_to_string(&val.data)
		return C.GoStringN(str.data, str.size), nil
	}
	return "", fmt.Errorf("Can't convert %s to string", val.typeAsString())
}
