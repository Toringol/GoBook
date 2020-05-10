/*
* Тип sort.Interface можно адаптировать для других
* применений. Напишите функцию IsPalindrome(s sort.Interface) bool,
* которая сообщает, является ли последовательность s палиндромом (другими словами,
* что обращение последовательности не изменяет ее). Считайте, что элементы
* с индексами i и j равны, если !s.Less(i, j) && !s.Less(j, i).
 */

package main

import (
	"fmt"
	"sort"
)

type palidromeCheck []byte

func (p palidromeCheck) Len() int           { return len(p) }
func (p palidromeCheck) Less(i, j int) bool { return p[i] < p[j] }
func (p palidromeCheck) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(palidromeCheck([]byte("abbaa"))))
}
