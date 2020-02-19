package v8

// #include <v8capi.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Value struct {
	data C.struct_v8_value
}

func (val Value) Dispose() {
	C.v8_delete_value(&val.data)
}

func typeToString(data C.struct_v8_value) string {
	switch C.v8_get_value_type(data) {
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
	return bool(C.v8_is_double(val.data))
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

func toBool(data C.struct_v8_value) (bool, error) {
	if bool(C.v8_is_boolean(data)) {
		return bool(C.v8_to_bool(data)), nil
	}
	return false, fmt.Errorf("Can't convert %s to bool", typeToString(data))
}

func (val Value) ToBool() (bool, error) {
	return toBool(val.data)
}

func toInt(data C.struct_v8_value) (int64, error) {
	if bool(C.v8_is_integer(data)) {
		return int64(C.v8_to_int64(data)), nil
	}
	return 0, fmt.Errorf("Can't convert %s to int64", typeToString(data))
}

func (val Value) ToInt() (int64, error) {
	return toInt(val.data)
}

func toUint(data C.struct_v8_value) (uint64, error) {
	if bool(C.v8_is_integer(data)) {
		i := int64(C.v8_to_int64(data))
		if i < 0 {
			return 0, fmt.Errorf("Can't cast negative value %d to unsigned value", i)
		}
		return uint64(i), nil
	}
	return 0, fmt.Errorf("Can't convert %s to uint64", typeToString(data))
}

func (val Value) ToUint() (uint64, error) {
	return toUint(val.data)
}

func toFloat(data C.struct_v8_value) (float64, error) {
	if bool(C.v8_is_number(data)) {
		return float64(C.v8_to_double(data)), nil
	}
	return 0., fmt.Errorf("Can't convert %s to float64", typeToString(data))
}

func (val Value) ToFloat() (float64, error) {
	return toFloat(val.data)
}

func toString(data C.struct_v8_value) (string, error) {
	if bool(C.v8_is_string(data)) {
		str := C.v8_to_string(&data)
		return C.GoStringN(str.data, str.size), nil
	}
	return "", fmt.Errorf("Can't convert %s to string", typeToString(data))
}

func (val Value) ToString() (string, error) {
	return toString(val.data)
}

func toObject(typeName string, data C.struct_v8_value, value reflect.Value) error {
	if !bool(C.v8_is_object(data)) {
		return fmt.Errorf("Can't convert %q to %q", typeToString(data), typeName)
	}

	obj := C.v8_to_object(data)

	size := int(obj.size)
	ptr := unsafe.Pointer(obj.data)
	elemSize := unsafe.Sizeof(*obj.data)

	applier := reflect.Indirect(value)

	for i := 0; i < size; i++ {
		data := (*C.struct_v8_pair_value)(ptr)

		fieldName, err := toString(data.first)
		if err != nil {
			return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
		}

		field := applier.FieldByName(fieldName)
		if !field.IsValid() {
			return fmt.Errorf("Type %q does not has field %q", typeName, fieldName)
		}

		fieldPtr := unsafe.Pointer(field.UnsafeAddr())

		switch field.Kind() {
		case reflect.Bool:
			val, err := toBool(data.second)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
			valPtr := (*bool)(fieldPtr)
			*valPtr = val
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			val, err := toInt(data.second)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
			valPtr := (*int64)(fieldPtr)
			*valPtr = val
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			val, err := toUint(data.second)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
			valPtr := (*uint64)(fieldPtr)
			*valPtr = val
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			val, err := toFloat(data.second)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
			valPtr := (*float64)(fieldPtr)
			*valPtr = val
		case reflect.Array:
			err := toArray(data.second, field.Elem().Kind(), fieldPtr)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
		case reflect.Map:
		case reflect.Slice:
			err := toArray(data.second, field.Type().Elem().Kind(), fieldPtr)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
		case reflect.String:
			val, err := toString(data.second)
			if err != nil {
				return fmt.Errorf("At %s.%s: %s", typeName, fieldName, err)
			}
			valPtr := (*string)(fieldPtr)
			*valPtr = val
		case reflect.Struct:
			err := toObject(typeName+"."+fieldName, data.second, reflect.NewAt(field.Type(), fieldPtr))
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Field %q of type %q has unsupported type", fieldName, typeName)
		}

		ptr = unsafe.Pointer(uintptr(ptr) + elemSize)
	}

	return nil
}

func (val Value) ToObject(res interface{}) error {
	typeName := reflect.TypeOf(res).Elem().Name()
	value := reflect.ValueOf(res)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("You must pass type %q by pointer", typeName)
	}
	return toObject(typeName, val.data, value)
}

func toBoolArray(data C.struct_v8_value) ([]bool, error) {
	if !bool(C.v8_is_array(data)) {
		return nil, fmt.Errorf("Can't convert %s to []bool", typeToString(data))
	}

	arr := C.v8_to_array(data)

	arrSize := int(arr.size)
	ptr := unsafe.Pointer(arr.data)
	elemSize := unsafe.Sizeof(*arr.data)

	res := make([]bool, arrSize)

	for i := 0; i < arrSize; i++ {
		val, err := toBool(*(*C.struct_v8_value)(ptr))
		if err != nil {
			return nil, fmt.Errorf("[%d]: %s", i, err)
		}
		res[i] = val
		ptr = unsafe.Pointer(uintptr(ptr) + elemSize)
	}

	return res, nil
}

func (val Value) ToBoolArray() ([]bool, error) {
	return toBoolArray(val.data)
}

func toIntArray(data C.struct_v8_value) ([]int64, error) {
	if !bool(C.v8_is_array(data)) {
		return nil, fmt.Errorf("Can't convert %s to []int64", typeToString(data))
	}

	arr := C.v8_to_array(data)

	arrSize := int(arr.size)
	ptr := unsafe.Pointer(arr.data)
	elemSize := unsafe.Sizeof(*arr.data)

	res := make([]int64, arrSize)

	for i := 0; i < arrSize; i++ {
		val, err := toInt(*(*C.struct_v8_value)(ptr))
		if err != nil {
			return nil, fmt.Errorf("[%d]: %s", i, err)
		}
		res[i] = val
		ptr = unsafe.Pointer(uintptr(ptr) + elemSize)
	}

	return res, nil
}

func (val Value) ToIntArray() ([]int64, error) {
	return toIntArray(val.data)
}

func toUintArray(data C.struct_v8_value) ([]uint64, error) {
	if !bool(C.v8_is_array(data)) {
		return nil, fmt.Errorf("Can't convert %s to []uint64", typeToString(data))
	}

	arr := C.v8_to_array(data)

	arrSize := int(arr.size)
	ptr := unsafe.Pointer(arr.data)
	elemSize := unsafe.Sizeof(*arr.data)

	res := make([]uint64, arrSize)

	for i := 0; i < arrSize; i++ {
		val, err := toUint(*(*C.struct_v8_value)(ptr))
		if err != nil {
			return nil, fmt.Errorf("[%d]: %s", i, err)
		}
		res[i] = val
		ptr = unsafe.Pointer(uintptr(ptr) + elemSize)
	}

	return res, nil
}
func (val Value) ToUintArray() ([]uint64, error) {
	return toUintArray(val.data)
}

func toFloatArray(data C.struct_v8_value) ([]float64, error) {
	if !bool(C.v8_is_array(data)) {
		return nil, fmt.Errorf("Can't convert %s to []float64", typeToString(data))
	}

	arr := C.v8_to_array(data)

	arrSize := int(arr.size)
	ptr := unsafe.Pointer(arr.data)
	elemSize := unsafe.Sizeof(*arr.data)

	res := make([]float64, arrSize)

	for i := 0; i < arrSize; i++ {
		val, err := toFloat(*(*C.struct_v8_value)(ptr))
		if err != nil {
			return nil, fmt.Errorf("[%d]: %s", i, err)
		}
		res[i] = val
		ptr = unsafe.Pointer(uintptr(ptr) + elemSize)
	}

	return res, nil
}

func (val Value) ToFloatArray() ([]float64, error) {
	return toFloatArray(val.data)
}

func toStringArray(data C.struct_v8_value) ([]string, error) {
	if !bool(C.v8_is_array(data)) {
		return nil, fmt.Errorf("Can't convert %s to []string", typeToString(data))
	}

	arr := C.v8_to_array(data)

	arrSize := int(arr.size)
	ptr := unsafe.Pointer(arr.data)
	elemSize := unsafe.Sizeof(*arr.data)

	res := make([]string, arrSize)

	for i := 0; i < arrSize; i++ {
		val, err := toString(*(*C.struct_v8_value)(ptr))
		if err != nil {
			return nil, fmt.Errorf("[%d]: %s", i, err)
		}
		res[i] = val
		ptr = unsafe.Pointer(uintptr(ptr) + elemSize)
	}

	return res, nil
}

func (val Value) ToStringArray() ([]string, error) {
	return toStringArray(val.data)
}

func (val Value) ToArray() ([]interface{}, error) {
	return nil, nil
}

func toArray(data C.struct_v8_value, typeKind reflect.Kind, ptr unsafe.Pointer) error {
	switch typeKind {
	case reflect.Bool:
		val, err := toBoolArray(data)
		if err != nil {
			return err
		}
		valPtr := (*[]bool)(ptr)
		*valPtr = val
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		val, err := toIntArray(data)
		if err != nil {
			return err
		}
		valPtr := (*[]int64)(ptr)
		*valPtr = val
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		val, err := toUintArray(data)
		if err != nil {
			return err
		}
		valPtr := (*[]uint64)(ptr)
		*valPtr = val
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		val, err := toFloatArray(data)
		if err != nil {
			return err
		}
		valPtr := (*[]float64)(ptr)
		*valPtr = val
	case reflect.Array:
	case reflect.Map:
	case reflect.String:
		val, err := toStringArray(data)
		if err != nil {
			return err
		}
		valPtr := (*[]string)(ptr)
		*valPtr = val
	case reflect.Struct:
	}
	return nil
}
