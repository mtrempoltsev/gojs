package gojs

import "github.com/mtrempoltsev/gojs"

func runScript(id, code string) (gojs.Value, error) {
	err := _js.Compile(id, code)
	if err != nil {
		return nil, err
	}

	future, err := _js.Run(id)
	if err != nil {
		return nil, err
	}

	res := <-future
	if res.Err != nil {
		return nil, err
	}

	return res.Val, nil
}
