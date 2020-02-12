package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommon(t *testing.T) {
	res, err := runScript("my.js", "2 + 2")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	defer res.Dispose()

	val, err := res.ToInt()

	assert.NoError(t, err)
	assert.Equal(t, val, int64(4))
}
