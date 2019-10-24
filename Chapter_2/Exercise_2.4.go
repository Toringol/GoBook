/*
* Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента
* по всем 64 позициям, проверяя при каждом сдвиге крайний справа бит. Сравните
* произовдительность этой версии с выборкой из таблицы.
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
	var i uint
	for i = 0; i < 64; i++ {
		result += int(byte(x>>i) & 1)
	}
	return result
}

func main() {
	start := time.Now()
	// 94.5422ms -------- Variant with init
	// 50.8635ms -------- Variant with checking every bit
	for i := 0; i < 100; i++ {
		fmt.Println(PopCount(378923689236823867))
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
