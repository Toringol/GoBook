/*
* Выражение x&(x-1) сбрасывает крайний справа ненужный ненулевой бит x.
* Напишите версию PopCount, которая подсчитывает биты с использованием
* этого факта, и оцените ее производительность.
 */

package main

import (
	"fmt"
	"time"
)

// var pc [256]byte

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
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

func main() {
	start := time.Now()
	// 94.5422ms -------- Variant with init
	// 54.3685ms -------- Variant with drop right bit x & (x - 1)
	for i := 0; i < 100; i++ {
		fmt.Println(PopCount(378923689236823867))
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
