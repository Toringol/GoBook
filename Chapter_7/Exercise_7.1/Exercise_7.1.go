/*
* Используя идеи из ByteCounter, реализуйте счетчики для слов
* и строк. Вам пригодится функция bufio.ScanWords.
 */

package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)

	counter := 0
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	*wc += WordCounter(counter)

	return counter, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))

	counter := 0
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	*lc += LineCounter(counter)

	return counter, nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("hello i am Sergey"))
	fmt.Println(wc)

	wc = 0
	var text = "asighas sdg l aighak a aisghasi"
	fmt.Fprintf(&wc, "hello i am Sergey %s", text)
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("hello i am Sergey\nqwe"))
	fmt.Println(lc)

	lc = 0
	var text1 = "asighas sdg \nl aighak a \naisghasi"
	fmt.Fprintf(&lc, "hello i am \nSergey\n %s", text1)
	fmt.Println(lc)
}
