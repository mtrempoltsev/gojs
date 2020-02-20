package main

import (
	"fmt"
	"github.com/mtrempoltsev/gojs"
)

func main() {
	js, err := gojs.New(1) // use 1 thread
	if err != nil {
		fmt.Println(err)
		return
	}

	defer js.Dispose()

	err = js.Compile("my.js", "2 + 2")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := js.Run("my.js")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Dispose()

	val, err := res.ToInt()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T, %d\n", val, val) // int64, 4
}
