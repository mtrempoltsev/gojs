package benchmarks

import "github.com/mtrempoltsev/gojs/engines"

func runScript(id, code string) (engines.Value, error) {
	err := _jsExecutor.Compile(id, code)
	if err != nil {
		return nil, err
	}

	return _jsExecutor.Run(id)
}
