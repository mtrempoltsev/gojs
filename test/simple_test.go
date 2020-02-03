package js

import (
	"fmt"

	"testing"

	js "github.com/mtrempoltsev/go-with-js"
)

func TestCommon(t *testing.T) {
	js, err := js.New(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = js.Compile("my.js", "2 + 2")
	if err != nil {
		fmt.Println(err)
		return
	}

	future, err := js.Run("my.js")
	if err != nil {
		fmt.Println(err)
		return
	}

	res := <-future
	if res.Err != nil {
		fmt.Println(res.Err)
		return
	}

	js.Dispose()
}
