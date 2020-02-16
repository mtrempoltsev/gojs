package gojs

import (
	"errors"
	"fmt"
	"runtime"
	"sync"

	"github.com/mtrempoltsev/gojs/engines"
	"github.com/mtrempoltsev/gojs/engines/v8"
)

type Result struct {
	Val engines.Value
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
	args []engines.Value
	res  ResultChannel
}

type taskChannel chan *task

type scriptCtx struct {
	script    engines.Script
	functions map[string]engines.Function
}

func (ctx *scriptCtx) run() (engines.Value, error) {
	return ctx.script.Run()
}

func (ctx *scriptCtx) dispose() {
	for _, function := range ctx.functions {
		function.Dispose()
	}
	ctx.script.Dispose()
}

type runnerCtx struct {
	runner       engines.Runner
	scripts      map[string]*scriptCtx
	pendingTasks taskChannel
	mutex        sync.RWMutex
}

func (ctx *runnerCtx) start() {
	for {
		task := <-ctx.pendingTasks
		if task == nil {
			continue
		}

		defer close(task.res)

		switch task.cmd {
		case run:
			ctx.mutex.RLock()
			script := ctx.scripts[task.name]
			ctx.mutex.RUnlock()

			if script == nil {
				task.res <- &Result{
					Val: nil,
					Err: fmt.Errorf("gojs.Executor: can't find script '%s'", task.name),
				}
			} else {
				res, err := script.run()
				task.res <- &Result{
					Val: res,
					Err: err,
				}
			}
		case callFunction:
			break
		}
	}
}

func (ctx *runnerCtx) compile(scriptName, code string) (*scriptCtx, error) {
	script, err := ctx.runner.Compile(scriptName, code)
	if err != nil {
		return nil, err
	}

	return &scriptCtx{
		script:    script,
		functions: make(map[string]engines.Function),
	}, nil
}

func (ctx *runnerCtx) dispose() {
	for _, script := range ctx.scripts {
		script.dispose()
	}
	ctx.runner.Dispose()
}

type Executor struct {
	engine       engines.Engine
	pendingTasks taskChannel
	runners      []*runnerCtx
}

func (executor *Executor) newRunner() (*runnerCtx, error) {
	runner, err := executor.engine.NewRunner()
	if err != nil {
		return nil, err
	}

	instance := &runnerCtx{
		runner:       runner,
		pendingTasks: executor.pendingTasks,
		scripts:      make(map[string]*scriptCtx),
	}

	return instance, nil
}

func New(runnersNum int) (*Executor, error) {
	if runnersNum < 0 {
		return nil, errors.New(
			"gojs.Executor.New: number of runners must be a positive number " +
				"or zero to use a size equal to the number of cores")
	}

	if runnersNum == 0 {
		runnersNum = runtime.NumCPU()
	}

	engine, err := v8.New()
	if err != nil {
		return nil, err
	}

	instance := Executor{
		engine:       engine,
		pendingTasks: make(taskChannel),
		runners:      make([]*runnerCtx, runnersNum),
	}

	for i := 0; i < runnersNum; i++ {
		instance.runners[i], err = instance.newRunner()
		if err != nil {
			for j := 0; j < i; j++ {
				instance.runners[j].dispose()
			}
			instance.engine.Dispose()
			return nil, err
		}
	}

	for _, runner := range instance.runners {
		go runner.start()
	}

	return &instance, nil
}

func (executor *Executor) Compile(scriptName, code string) error {
	if len(scriptName) == 0 {
		return errors.New("gojs.Executor.Compile: you must specify scriptID")
	}

	if len(code) == 0 {
		return errors.New("gojs.Executor.Compile: code is empty, nothing to compile")
	}

	type results struct {
		index  int
		script *scriptCtx
		err    error
	}

	n := len(executor.runners)

	channel := make(chan results, n)

	for i := 0; i < n; i++ {
		go func(i int) {
			script, err := executor.runners[i].compile(scriptName, code)
			channel <- results{i, script, err}
		}(i)
	}

	scripts := make([]*scriptCtx, n)

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
		runner := executor.runners[i]
		runner.mutex.Lock()
		runner.scripts[scriptName] = scripts[i]
		runner.mutex.Unlock()
	}

	return nil
}

func (executor *Executor) Run(scriptName string) (ResultChannel, error) {
	if len(scriptName) == 0 {
		return nil, errors.New("gojs.Executor.Run: you must specify scriptID")
	}

	res := make(ResultChannel)

	executor.pendingTasks <- &task{
		cmd:  run,
		name: scriptName,
		res:  res,
	}

	return res, nil
}

func (executor *Executor) Dispose() {
	for _, runner := range executor.runners {
		runner.dispose()
	}
	executor.engine.Dispose()
}
