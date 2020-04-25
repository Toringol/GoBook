/*
* Напишите вариативную версию функции strings.Join.
 */

package main

import "fmt"

func JoinStrs(sep string, strs ...string) string {
	finalString := ""
	for i, val := range strs {
		if i == len(strs)-1 {
			finalString += val
			break
		}
		finalString += val + sep
	}
	return finalString
}

func main() {
	fmt.Println(JoinStrs(", ", "foo", "bar", "baz"))
}
