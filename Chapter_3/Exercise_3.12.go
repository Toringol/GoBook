/*
* Напишите функцию, которая сообщает, являются ли две
* строки анаграммами одна другой, т.е. состоят ли они из одних и
* тех же букв в другом порядке.
 */

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Можно сделать через map, но здесь глава про фундаментальные типы данных
// Думаю имелось ввиду решение с библиотекой strings
func anagram(s1 string, s2 string) bool {
	firstStr := strings.Split(s1, "")
	secondStr := strings.Split(s2, "")

	sort.Strings(firstStr)
	sort.Strings(secondStr)

	return strings.Join(firstStr, "") == strings.Join(secondStr, "")
}

func main() {
	s1 := "hello"
	s2 := "llohe"

	if anagram(s1, s2) {
		fmt.Println("Correct solution")
	} else {
		fmt.Println("Incorrect")
	}

}
