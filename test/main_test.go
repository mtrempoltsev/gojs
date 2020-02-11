package gojs

import (
	"os"
	"testing"

	"github.com/mtrempoltsev/gojs"
)

var _js *gojs.Engine

func TestMain(m *testing.M) {
	var err error

	_js, err = gojs.New(0)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	code := m.Run()

	_js.Dispose()

	os.Exit(code)
}
