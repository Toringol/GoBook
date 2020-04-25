/*
* Напишите вариативные функции max и min, аналогичные
* функции sum. Что должны делать эти функци, будучи
* вызванными без аргументов? Напишите варианты функций,
* требующие как минимум одного аргумента.
 */

package main

import (
	"errors"
	"fmt"
)

func max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, errors.New("Zero length")
	}

	max := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}

	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, errors.New("Zero length")
	}

	min := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}

	return min, nil
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
}
