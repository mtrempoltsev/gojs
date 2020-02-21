package benchmarks

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkV8Values(b *testing.B) {
	const id = "my.js"
	const code = "var x = {" +
		"	b: true," +
		"	i: -1," +
		"	u: 1, " +
		"	f: 0.5, " +
		"	a1: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], " +
		"	a2: [1., 2., 3., 4., 5., 6., 7., 8., 9., 10.], " +
		"	s1: 'ok', " +
		"	s2: 'Select executes a select operation described by the list of cases. Like the Go select statement, it blocks until at least one of the cases can proceed, makes a uniform pseudo-random choice, and then executes that case. It returns the index of the chosen case and, if that case was a receive operation, the value received and a boolean indicating whether the value corresponds to a send on the channel (as opposed to a zero value received because the channel is closed).'," +
		"	o: { x: 2, y: false }, " +
		"}; x"

	err := _jsExecutor.Compile(id, code)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := _jsExecutor.Run(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Dispose()

	type nestedObj struct {
		x int
		y bool
	}

	type testObj struct {
		b  bool
		i  int
		u  uint
		f  float64
		a1 []int64
		a2 []float64
		s1 string
		s2 string
		o  nestedObj
	}

	var obj testObj

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = res.ToObject(&obj)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
