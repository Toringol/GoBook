/*
* Усовершенствуйте функцию comma так, чтобы она корректно
* работала с числами с плавающей точкой и необязательным знаком.
 */

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)

	if s[0] == '-' {
		s = s[1:]
		buf.WriteByte('-')
		n--
	}

	dot := strings.Index(s, ".")
	if dot != strings.LastIndex(s, ".") || dot == len(s)-1 {
		return "Incorrect string"
	} else if dot != -1 {
		firstPartStr := s[:dot]
		secondPartStr := s[dot+1:]
		buf.WriteString(comma(firstPartStr))
		buf.WriteByte('.')
		buf.WriteString(comma(secondPartStr))
		return buf.String()
	}

	if n <= 3 {
		return s
	}

	for i, letter := range s {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(letter)
	}

	return buf.String()
}

func main() {
	s := "-12345.12345"
	fmt.Println(comma(s))
}
