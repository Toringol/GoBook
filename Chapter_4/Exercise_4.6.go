/*
* Напишите функцию, которая без выделения дополнительной памяти
* преобразует последовательность смежных пробельных символов Unicode
* в срезе []byte в кодировке UTF-8 в один пробел ASCII.
 */

package main

import (
	"fmt"
	"unicode"
)

func removeSymbol(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeAdjacentSpaces(seq []byte) []byte {
	if len(seq) <= 1 {
		return seq
	}

	firstSpace := false
	spaceSeq := false
	indexEndSeq := 0
	for i, indexStartSeq := 0, 0; i < len(seq); i++ {

		if unicode.IsSpace(rune(seq[i])) && !firstSpace && !spaceSeq { // First WhiteSpace
			firstSpace = true
			continue
		} else if unicode.IsSpace(rune(seq[i])) && firstSpace && !spaceSeq { // Starting Sequence of WhiteSpaces
			spaceSeq = true
			firstSpace = false
			indexStartSeq = i
			continue
		}

		if unicode.IsSpace(rune(seq[i])) && spaceSeq { // Continue Sequence of WhiteSpaces
			continue
		} else if !unicode.IsSpace(rune(seq[i])) && spaceSeq { // End Sequence of WhiteSpaces
			spaceSeq = false
			indexEndSeq = i
			for j := indexStartSeq; j < indexEndSeq; j++ {
				seq = removeSymbol(seq, indexStartSeq)
			}
		}

		firstSpace = false
	}
	return seq
}

func main() {
	str := []byte{'b', 'a', ' ', ' ', 'a', ' '}
	str = removeAdjacentSpaces(str)
	fmt.Println(str)
}
