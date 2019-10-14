package main

/*
* Измените программу echo так, чтобы она выводила также os.Args[0],
* имя выполняемой программы.
*/

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
