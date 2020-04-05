/*
* Напишите программу wordfreq для подсчета частоты каждого
* слова во входном текстовом файле. Вызовите input.Split(bufio.ScanWords)
* до первого вызова Scan для разбивки текста на слова, а не на строки.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	wordfreq := make(map[string]int)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordfreq[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for word, freq := range wordfreq {
		fmt.Printf("%q\t%d\n", word, freq)
	}
}
