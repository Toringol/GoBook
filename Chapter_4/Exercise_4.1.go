/*
* Напишите функцию, которая подсчитывает количество битов,
* различных в двух дайджестах SHA256.
 */

package main

import (
	"crypto/sha256"
	"fmt"
)

func PopCount(x byte) int {
	var result int
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

func VariousCountBits(firstValue [32]byte, secondValue [32]byte) int {
	result := 0
	for i := 0; i < 32; i++ {
		result += PopCount(firstValue[i] ^ secondValue[i])
	}
	return result
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(VariousCountBits(c1, c2))
}
