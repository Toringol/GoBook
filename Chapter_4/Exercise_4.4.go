/*
* Напишите версию функции rotate, которая работает
* в один проход. (циклический сдвиг влево на один)
 */

package main

import "fmt"

func rotate(values []int) {
	if len(values) <= 1 {
		return
	}

	first := values[0]
	copy(values, values[1:])
	values[len(values)-1] = first
}

func main() {
	values := []int{0, 1, 2, 3, 4, 5}
	rotate(values)
	fmt.Println(values)
}
