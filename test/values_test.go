package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUndefined(t *testing.T) {
	res, err := runScript("my.js", "undefined")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())
}

func TestBool(t *testing.T) {
	res, err := runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.True(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	val, err := res.ToBool()

	assert.NoError(t, err)
	assert.Equal(t, true, val)

	res, err = runScript("my.js", "1")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	val, err = res.ToBool()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert number to bool", err.Error())
}

func TestNull(t *testing.T) {
	res, err := runScript("my.js", "null")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.True(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())
}

func TestInt(t *testing.T) {
	res, err := runScript("my.js", "-2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.True(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.True(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	val, err := res.ToInt()

	assert.NoError(t, err)
	assert.Equal(t, int64(-2), val)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	val, err = res.ToInt()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to int64", err.Error())
}

func TestUint(t *testing.T) {
	res, err := runScript("my.js", "2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.True(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.True(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	val, err := res.ToUint()

	assert.NoError(t, err)
	assert.Equal(t, uint64(2), val)

	res, err = runScript("my.js", "-2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.True(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.True(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	val, err = res.ToUint()

	assert.Error(t, err)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	val, err = res.ToUint()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to uint64", err.Error())
}

func TestFloat(t *testing.T) {
	res, err := runScript("my.js", "2.5")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.True(t, res.IsNumber())
	assert.True(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	val, err := res.ToFloat()

	assert.NoError(t, err)
	assert.Equal(t, float64(2.5), val)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	val, err = res.ToFloat()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to float64", err.Error())
}

func TestString(t *testing.T) {
	res, err := runScript("my.js", "'ok'")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.True(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	val, err := res.ToString()

	assert.NoError(t, err)
	assert.Equal(t, "ok", val)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	val, err = res.ToString()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to string", err.Error())
}

func TestObject(t *testing.T) {
	res, err := runScript("my.js", "let x = {"+
		"b: true,"+
		"i: -1,"+
		"u: 1, "+
		"f: 0.5, "+
		"a: [1, 2], "+
		//"m: {}, "+
		"s: 'ok', "+
		"o: { x: 2 }, "+
		"}; x")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.True(t, res.IsObject())
	assert.False(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	type nestedObj struct {
		x int
	}

	type testObj struct {
		b bool
		i int
		u uint
		f float64
		a []int64
		s string
		o nestedObj
	}

	var obj testObj

	err = res.ToObject(&obj)

	assert.NoError(t, err)
	assert.Equal(t, true, obj.b)
	assert.Equal(t, -1, obj.i)
	assert.Equal(t, uint(1), obj.u)
	assert.Equal(t, 0.5, obj.f)
	assert.Equal(t, []int64{1, 2}, obj.a)
	assert.Equal(t, "ok", obj.s)
	assert.Equal(t, 2, obj.o.x)
}

func TestBoolArray(t *testing.T) {
	res, err := runScript("my.js", "[true, false, true]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.True(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	arr, err := res.ToBoolArray()

	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false, true}, arr)

	res, err = runScript("my.js", "1")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToBoolArray()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert number to []bool", err.Error())

	res, err = runScript("my.js", "[true, 1, true]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToBoolArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert number to bool", err.Error())
}

func TestIntArray(t *testing.T) {
	res, err := runScript("my.js", "[-1, 0, 1]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.True(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	arr, err := res.ToIntArray()

	assert.NoError(t, err)
	assert.Equal(t, []int64{-1, 0, 1}, arr)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToIntArray()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to []int64", err.Error())

	res, err = runScript("my.js", "[1, true, 3]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToIntArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert boolean to int64", err.Error())

	res, err = runScript("my.js", "[1, 1.5]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToIntArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert number to int64", err.Error())
}

func TestUintArray(t *testing.T) {
	res, err := runScript("my.js", "[0, 1, 2]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.True(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	arr, err := res.ToUintArray()

	assert.NoError(t, err)
	assert.Equal(t, []uint64{0, 1, 2}, arr)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToUintArray()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to []uint64", err.Error())

	res, err = runScript("my.js", "[0, true]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToUintArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert boolean to uint64", err.Error())

	res, err = runScript("my.js", "[1, 1.5]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToUintArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert number to uint64", err.Error())

	res, err = runScript("my.js", "[0, -1]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToUintArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't cast negative value -1 to unsigned value", err.Error())
}

func TestFloatArray(t *testing.T) {
	res, err := runScript("my.js", "[-1.5, 0., 1]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.True(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	arr, err := res.ToFloatArray()

	assert.NoError(t, err)
	assert.Equal(t, []float64{-1.5, 0., 1.}, arr)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToFloatArray()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to []float64", err.Error())

	res, err = runScript("my.js", "[1, true, 3]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToFloatArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert boolean to float64", err.Error())
}

func TestStringArray(t *testing.T) {
	res, err := runScript("my.js", "['one', 'two']")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.False(t, res.IsUndefined())
	assert.False(t, res.IsBoolean())
	assert.False(t, res.IsNull())
	assert.False(t, res.IsNumber())
	assert.False(t, res.IsDouble())
	assert.False(t, res.IsInteger())
	assert.False(t, res.IsString())
	assert.False(t, res.IsObject())
	assert.True(t, res.IsArray())
	assert.False(t, res.IsSet())
	assert.False(t, res.IsMap())

	arr, err := res.ToStringArray()

	assert.NoError(t, err)
	assert.Equal(t, []string{"one", "two"}, arr)

	res, err = runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToStringArray()

	assert.Error(t, err)
	assert.Equal(t, "Can't convert boolean to []string", err.Error())

	res, err = runScript("my.js", "['one', 2, 3]")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	arr, err = res.ToStringArray()

	assert.Error(t, err)
	assert.Equal(t, "[1]: Can't convert number to string", err.Error())
}
