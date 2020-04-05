/*
* Измените charcount так, чтобы программа подсчитывала
* количество букв, цифр и прочих категорий Unicode с использованием
* функций наподобие unicode.IsLetter.
 */

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	NUMBER = iota
	LETTER
	MARK
	SPACE
)

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	categoryCount := [4]int{}
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			categoryCount[LETTER]++
		} else if unicode.IsMark(r) {
			categoryCount[MARK]++
		} else if unicode.IsSpace(r) {
			categoryCount[SPACE]++
		} else if unicode.IsNumber(r) {
			categoryCount[NUMBER]++
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("category\tcount\n")
	for i, n := range categoryCount {
		fmt.Printf("%d\t%d\n", i, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
