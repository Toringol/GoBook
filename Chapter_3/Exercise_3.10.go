/*
* Напишите нерекурсивную версию функции comma,
* использующую bytes.Buffer вместо конкатенации строк.
 */

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	for i, letter := range s {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(letter)
	}

	return buf.String()
}

func main() {
	s := "12345"
	fmt.Println(comma(s))
}
