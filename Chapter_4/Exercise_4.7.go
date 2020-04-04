/*
* Перепишите функцию reverse так, чтобы она без выделения
* дополнительной памяти обращала последовательность символов среза
* []byte, который представляет строку в кодировке UTF-8. Сможете ли
* вы обойтись без выделения новой памяти?
 */

package main

import "fmt"

func reverse(s []byte) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	s := "asdasdas"
	s = reverse([]byte(s))
	fmt.Println(s)
}
