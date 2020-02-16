package v8

// #cgo CFLAGS: -I../../thirdparty/v8capi/include -O0 -g
// #cgo LDFLAGS: -L../../out -lv8capi -lv8_monolith
// #cgo LDFLAGS: -lstdc++ -lm
// #include <stdlib.h>
// #include <v8capi.h>
import "C"

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/mtrempoltsev/gojs/engines"
)

func makeError(err C.struct_v8_error) string {
	if err.location == nil {
		return C.GoString(err.message)
	}

	buf := strings.Builder{}

	fmt.Fprintf(&buf,
		"%s:%d: %s\n%s",
		C.GoString(err.location),
		err.line_number,
		C.GoString(err.message),
		C.GoString(err.wavy_underline))

	if err.stack_trace != nil {
		fmt.Fprintf(&buf,
			"\nstack trace:\n%s",
			C.GoString(err.stack_trace))
	}

	return buf.String()
}

type Function struct {
	ptr *C.struct_v8_callable
}

type Script struct {
	ptr *C.struct_v8_script
}

type Runner struct {
	ptr *C.struct_v8_isolate
}

type Engine struct {
	ptr *C.struct_v8_instance
}

func New() (engines.Engine, error) {
	path := C.CString(os.Args[0])
	defer C.free(unsafe.Pointer(path))

	return &Engine{
		ptr: C.v8_new_instance(path),
	}, nil
}

func (*Engine) NewRunner() (engines.Runner, error) {
	return &Runner{
		ptr: C.v8_new_isolate(),
	}, nil
}

func (engine *Engine) Dispose() {
	C.v8_delete_instance(engine.ptr)
}

func (runner *Runner) Compile(name, code string) (engines.Script, error) {
	codePtr := C.CString(code)
	defer C.free(unsafe.Pointer(codePtr))

	namePtr := C.CString(name)
	defer C.free(unsafe.Pointer(namePtr))

	var err C.struct_v8_error
	defer C.v8_delete_error(&err)

	script := C.v8_compile_script(runner.ptr, codePtr, namePtr, &err)

	if script == nil {
		return nil, errors.New(makeError(err))
	}

	return &Script{script}, nil
}

func (runner *Runner) Dispose() {
	C.v8_delete_isolate(runner.ptr)
}

func (script *Script) Run() (engines.Value, error) {
	var res C.struct_v8_value
	defer C.v8_delete_value(&res)

	var err C.struct_v8_error
	defer C.v8_delete_error(&err)

	if !C.v8_run_script(script.ptr, &res, &err) {
		return nil, errors.New(makeError(err))
	}

	return Value{data: res}, nil
}

func (script *Script) Terminate() {
	// TODO: implement this
}

func (script *Script) GetFunction(funcName string) (engines.Function, error) {
	// TODO: implement this
	return nil, nil
}

func (script *Script) Dispose() {
	C.v8_delete_script(script.ptr)
}

func (function *Function) Call(args ...engines.Value) (engines.Value, error) {
	// TODO: implement this
	return makeUndefined(), nil
}

func (function *Function) Terminate() {
	// TODO: implement this
}

func (function *Function) Dispose() {
	C.v8_delete_function(function.ptr)
}
