package benchmarks

import (
	"fmt"
	"testing"
)

func BenchmarkV8Run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := runScript("my.js",
			"var x = {"+
				"	b: true,"+
				"	i: -1,"+
				"	u: 1, "+
				"	f: 0.5, "+
				"	a1: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], "+
				"	a2: [1., 2., 3., 4., 5., 6., 7., 8., 9., 10.], "+
				"	s1: 'ok', "+
				"	s2: 'Select executes a select operation described by the list of cases. Like the Go select statement, it blocks until at least one of the cases can proceed, makes a uniform pseudo-random choice, and then executes that case. It returns the index of the chosen case and, if that case was a receive operation, the value received and a boolean indicating whether the value corresponds to a send on the channel (as opposed to a zero value received because the channel is closed).',"+
				"	o: { x: 2, y: false }, "+
				"}; x")

		if err != nil {
			fmt.Println(err)
			return
		}

		res.Dispose()
	}
}
