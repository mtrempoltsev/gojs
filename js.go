package js

import (
	"errors"
	"fmt"
	"runtime"
	"sync"

	v8 "github.com/mtrempoltsev/gojs/v8"
)

type script struct {
	ptr       *v8.Script
	functions map[string]*v8.Function
}

func (runner *runner) newScript(id, code string) (*script, error) {
	ptr, err := runner.isolate.Compile(code, id)
	if err != nil {
		return nil, err
	}

	return &script{
		ptr:       ptr,
		functions: make(map[string]*v8.Function),
	}, nil
}

func (script *script) dispose() {
	for _, function := range script.functions {
		function.Dispose()
	}
	script.ptr.Dispose()
}

func (script *script) run() (Value, error) {
	_, err := script.ptr.Run()
	return nil, err
}

type Value interface{}

type Result struct {
	Val Value
	Err error
}

type ResultChannel chan *Result

type command int

const (
	run command = iota
	callFunction
)

type task struct {
	cmd  command
	name string
	args []Value
	res  ResultChannel
}

type taskChannel chan *task

type runner struct {
	isolate      *v8.Isolate
	pendingTasks taskChannel
	scripts      map[string]*script
	mutex        sync.RWMutex
}

func (runner *runner) start() {
	for {
		task := <-runner.pendingTasks
		if task == nil {
			continue
		}

		switch task.cmd {
		case run:
			runner.mutex.RLock()
			script := runner.scripts[task.name]
			runner.mutex.RUnlock()
			if script == nil {
				task.res <- &Result{
					Val: nil,
					Err: fmt.Errorf("js.Runner: can't find script '%s'", task.name),
				}
			} else {
				res, err := script.run()
				if res != nil && err != nil {

				}
				task.res <- &Result{
					Val: nil,
					Err: nil,
				}
				close(task.res)
			}
			break
		case callFunction:
			break
		}
	}
}

func (engine *Engine) newRunner() *runner {
	instance := &runner{
		isolate:      engine.v8.NewIsolate(),
		pendingTasks: engine.pendingTasks,
		scripts:      make(map[string]*script),
	}

	go instance.start()

	return instance
}

func (runner *runner) dispose() {
	for _, script := range runner.scripts {
		script.dispose()
	}
	runner.isolate.Dispose()
}

type Engine struct {
	v8           *v8.Instance
	pendingTasks taskChannel
	runners      []*runner
}

func New(poolSize int) (*Engine, error) {
	if poolSize < 0 {
		return nil, errors.New(
			"js.New: poolSize must be a positive number or zero to use a size equal to thr number of cores")
	}

	if poolSize == 0 {
		poolSize = runtime.NumCPU()
	}

	instance := &Engine{
		v8:           v8.New(),
		pendingTasks: make(taskChannel),
		runners:      make([]*runner, poolSize),
	}

	for i := 0; i < poolSize; i++ {
		instance.runners[i] = instance.newRunner()
	}

	return instance, nil
}

func (engine *Engine) Dispose() {
	for _, runner := range engine.runners {
		runner.dispose()
	}
	engine.v8.Dispose()
}

func (engine *Engine) Compile(scriptID, code string) error {
	if len(scriptID) == 0 {
		return errors.New("js.Engine.Compile: you must specify scriptID")
	}

	if len(code) == 0 {
		return errors.New("js.Engine.Compile: code is empty, nothing to compile")
	}

	type results struct {
		index  int
		script *script
		err    error
	}

	n := len(engine.runners)

	channel := make(chan results, n)

	for i := 0; i < n; i++ {
		go func(i int) {
			script, err := engine.runners[i].newScript(scriptID, code)
			channel <- results{i, script, err}
		}(i)
	}

	scripts := make([]*script, n)

	var err error

	for i := 0; i < n; i++ {
		res := <-channel
		if res.err != nil {
			err = res.err
		} else {
			scripts[res.index] = res.script
		}
	}

	if err != nil {
		for i := 0; i < n; i++ {
			if scripts[i] != nil {
				scripts[i].dispose()
			}
		}
		return err
	}

	for i := 0; i < n; i++ {
		runner := engine.runners[i]
		runner.mutex.Lock()
		runner.scripts[scriptID] = scripts[i]
		runner.mutex.Unlock()
	}

	return nil
}

func (engine *Engine) Run(scriptID string) (ResultChannel, error) {
	if len(scriptID) == 0 {
		return nil, errors.New("js.Engine.Run: you must specify scriptID")
	}

	res := make(ResultChannel)

	engine.pendingTasks <- &task{
		cmd:  run,
		name: scriptID,
		res:  res,
	}

	return res, nil
}

func (engine *Engine) CallFunction(scriptID, funcName string, args ...Value) {

}
