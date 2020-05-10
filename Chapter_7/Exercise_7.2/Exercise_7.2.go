/*
* Напишите функцию CountingWriter с приведенной ниже
* сигнатурой, которая для данного io.Writer возращает
* новый Writer, являющиеся оболочкой исходного, и указатель
* на переменную int64, которая в любой момент содержит количество
* байтов записанных в новый Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64)
*/

package main

import (
	"bytes"
	"fmt"
	"io"
)

type ByteCounter struct {
	w       io.Writer
	written int64
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	n, err := bc.w.Write(p)
	bc.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := &ByteCounter{w, 0}
	return bc, &bc.written
}

func main() {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := []byte("skdhgdsk")
	c.Write(data)
	fmt.Println(*n)
}
