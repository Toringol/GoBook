/*
* Напишите набор тестов для множества IntSet (раздел 6.5),
* который проверяет, эквивалентно ли его поведение множеству
* на основе встроенный отображений. Сохраните реализацию для
* проверок производительности в упражнении 11.7.
 */

package main

import (
	"testing"
)

func TestIntSetAdd(t *testing.T) {
	var tests = []struct {
		input int
		add   int
		want  string
	}{
		{
			1,
			3,
			"{1 3}",
		},
	}

	for _, test := range tests {
		intSet := IntSet{}
		intSet.Add(test.input)
		intSet.Add(test.add)

		if got := intSet.String(); got != test.want {
			t.Errorf("Test failed - got: %v, want: %v", intSet.String(), test.want)
		}
	}
}
