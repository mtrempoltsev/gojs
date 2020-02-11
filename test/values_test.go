package gojs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndefined(t *testing.T) {
	res, err := runScript("my.js", "undefined")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsUndefined())
}

func TestBool(t *testing.T) {
	res, err := runScript("my.js", "true")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsBoolean())

	val, err := res.ToBool()

	assert.NoError(t, err)
	assert.Equal(t, val, true)
}

func TestNull(t *testing.T) {
	res, err := runScript("my.js", "null")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsNull())
}

func TestInt(t *testing.T) {
	res, err := runScript("my.js", "-2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsInteger())

	val, err := res.ToInt()

	assert.NoError(t, err)
	assert.Equal(t, val, int64(-2))
}

func TestUint(t *testing.T) {
	res1, err := runScript("my.js", "2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res1.Dispose()

	assert.True(t, res1.IsInteger())

	val, err := res1.ToUint()

	assert.NoError(t, err)
	assert.Equal(t, val, uint64(2))

	res2, err := runScript("my.js", "-2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res2.Dispose()

	assert.True(t, res2.IsInteger())

	val, err = res2.ToUint()

	assert.Error(t, err)
}

func TestFloat(t *testing.T) {
	res, err := runScript("my.js", "2.5")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsDouble())

	val, err := res.ToFloat()

	assert.NoError(t, err)
	assert.Equal(t, val, float64(2.5))
}

func TestString(t *testing.T) {
	res, err := runScript("my.js", "'ok'")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	assert.True(t, res.IsString())

	val, err := res.ToString()

	assert.NoError(t, err)
	assert.Equal(t, val, "ok")
}
