/*
* Функция strings.NewReader возвращает значение, соответствующее
* интерфейсу io.Reader (и другим), путем чтения из своего аргумента,
* который представляет собой строку. Реализуйте простую версию NewReader
* и используйте ее для создания синтаксического анализатора HTML (раздел 5.2),
* принимающего входные данные из строки.
 */

package main

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type stringReader struct {
	s string
}

func (sr *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, sr.s)
	sr.s = sr.s[n:]
	if len(sr.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}

func main() {
	s := "<html><body><p>hi</p></body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
}
