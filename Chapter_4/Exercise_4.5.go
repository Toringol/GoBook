/*
* Напишите функцию, которая без выделения дополнительной
* памяти удаляет все смежные дубликаты в срезе []string.
 */

package main

import "fmt"

func removeAdjacentDup(baseStr []string) []string {
	if len(baseStr) <= 1 {
		return baseStr
	}

	indexPreviousStr := 0
	for _, str := range baseStr {
		if baseStr[indexPreviousStr] == str {
			continue
		}

		indexPreviousStr++
		baseStr[indexPreviousStr] = str
	}

	return baseStr[:indexPreviousStr+1]
}

func main() {
	str := []string{"aa", "aa", "bb", "bb", "cc"}
	str = removeAdjacentDup(str)
	fmt.Println(str)
}
