package test

import "github.com/mtrempoltsev/gojs/internal/abstract"

func runScript(id, code string) (abstract.Value, error) {
	err := _jsExecutor.Compile(id, code)
	if err != nil {
		return nil, err
	}

	future, err := _jsExecutor.Run(id)
	if err != nil {
		return nil, err
	}

	res := <-future
	if res.Err != nil {
		return nil, res.Err
	}

	return res.Val, nil
}
