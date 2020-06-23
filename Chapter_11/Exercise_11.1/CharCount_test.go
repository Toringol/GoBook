/*
* Напишите тесты для программы charcount из раздела 4.3.
 */

package main

import (
	"bytes"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
		want  struct {
			wantCounts  map[rune]int
			wantUtfLen  [utf8.UTFMax + 1]int
			wantInvalid int
		}
	}{
		{
			"asdasdasd", struct {
				wantCounts  map[rune]int
				wantUtfLen  [utf8.UTFMax + 1]int
				wantInvalid int
			}{
				map[rune]int{'a': 3, 's': 3, 'd': 3},
				[utf8.UTFMax + 1]int{0, 9, 0, 0, 0},
				0,
			},
		},
	}

	for _, test := range tests {
		counts, utflen, invalid := CharCount(bytes.NewReader([]byte(test.input)))

		if !reflect.DeepEqual(test.want.wantCounts, counts) {
			t.Errorf("Counts failed - input: %s, got: %v, want: %v", test.input, counts, test.want.wantCounts)
		}

		if !reflect.DeepEqual(test.want.wantUtfLen, utflen) {
			t.Errorf("UtfLen failed - input: %s, got: %v, want: %v", test.input, utflen, test.want.wantUtfLen)
		}

		if test.want.wantInvalid != invalid {
			t.Errorf("Invalid failed - input: %s, got: %v, want: %v", test.input, invalid, test.want.wantInvalid)
		}

	}
}
