/*
* Перепишите функцию PopCount так, чтобы она использовала цикл вместо единого выражения.
* Сравните производительность двух версий. (В разделе 11.4 показано, как правильно
* сравнивать производительность различных реализаций.)
 */

package main

import (
	"fmt"
	"time"
)

var pc [256]byte

// Example variant
// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

// func PopCount(x uint64) int {
// 	return int(pc[byte(x>>(0*8))] +
// 		pc[byte(x>>(1*8))] +
// 		pc[byte(x>>(2*8))] +
// 		pc[byte(x>>(3*8))] +
// 		pc[byte(x>>(4*8))] +
// 		pc[byte(x>>(5*8))] +
// 		pc[byte(x>>(6*8))] +
// 		pc[byte(x>>(7*8))])
// }

func PopCount(x uint64) int {
	var result int
	var i uint
	for i = 0; i < 8; i++ {
		for j := range pc {
			pc[j] = pc[j/2] + byte(j&1)
		}
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func main() {
	start := time.Now()
	// 98.4883ms -------- Bad Variant without init
	// 94.5422ms -------- Variant with init
	for i := 0; i < 100; i++ {
		fmt.Println(PopCount(378923689236823867))
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
