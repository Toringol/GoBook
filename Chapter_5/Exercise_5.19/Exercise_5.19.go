/*
* Воспользуйтесь функциями panic и recover для написания
* функции, которая не содержит инструкцию return, но
* возвращает ненулевое значение.
 */

package main

import "fmt"

func resultReturn() (result string) {
	defer func() {
		recover()
		result = "error"
	}()

	panic("Error")
}

func main() {
	fmt.Println(resultReturn())
}
