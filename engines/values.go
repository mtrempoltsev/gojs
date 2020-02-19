package engines

type Value interface {
	Dispose()

	IsUndefined() bool
	IsBoolean() bool
	IsNull() bool
	IsNumber() bool
	IsDouble() bool
	IsInteger() bool
	IsString() bool
	IsObject() bool
	IsArray() bool
	IsSet() bool
	IsMap() bool

	ToBool() (bool, error)
	ToInt() (int64, error)
	ToUint() (uint64, error)
	ToFloat() (float64, error)
	ToString() (string, error)

	ToObject(obj interface{}) error

	ToBoolArray() ([]bool, error)
	ToIntArray() ([]int64, error)
	ToUintArray() ([]uint64, error)
	ToFloatArray() ([]float64, error)
	ToStringArray() ([]string, error)
	ToArray() ([]interface{}, error)
}
