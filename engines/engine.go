package engines

type Function interface {
	Call(args ...Value) (Value, error)
	Terminate()
	Dispose()
}

type Script interface {
	Run() (Value, error)
	Terminate()
	GetFunction(funcName string) (Function, error)
	Dispose()
}

type Runner interface {
	Compile(id, code string) (Script, error)
	Dispose()
}

type Engine interface {
	NewRunner() (Runner, error)
	Dispose()
}
