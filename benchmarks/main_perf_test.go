package benchmarks

import (
	"os"
	"testing"

	"github.com/mtrempoltsev/gojs"
)

var _jsExecutor *gojs.Executor

func TestMain(m *testing.M) {
	var err error

	_jsExecutor, err = gojs.New(1)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	code := m.Run()

	_jsExecutor.Dispose()

	os.Exit(code)
}
