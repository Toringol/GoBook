/*
* Расширьте TestSplit так, чтобы она использовала
* таблицу входных и ожидаемых выходных данных.
 */

package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		inputStr, inputSep string
		want               []string
	}{
		{
			"a:b:c",
			":",
			[]string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		outPut := strings.Split(test.inputStr, test.inputSep)
		if !reflect.DeepEqual(outPut, test.want) {
			t.Errorf("Split(%q, %q) - got: %v, want: %v",
				test.inputStr, test.inputSep, outPut, test.want)
		}
	}
}
