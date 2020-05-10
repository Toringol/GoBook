/*
* Функция LimitReader из пакета принимает переменную r
* типа io.Reader и количество байтов n и возвращает другой
* объект Reader, который читает из r, но после чтения n
* сообщает о достижении конца файла. Реализуйте его.
* func LimitReader(r io.Reader, n int64) io.Reader
 */

package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r                       io.Reader
	readAlready, limitBytes int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	n, err = lr.r.Read(p[:lr.limitBytes])
	lr.readAlready += int64(n)
	if lr.readAlready >= lr.limitBytes {
		fmt.Println("Limit is over")
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, limitBytes: n}
}

func main() {
	s := "Hello"
	sr := strings.NewReader(s)
	limit := int64(3)
	nsr := LimitReader(sr, limit)
	p := make([]byte, 10)
	nsr.Read(p)
}
