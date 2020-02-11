package gojs

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
}
