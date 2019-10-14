package main

/*
* Измените программу echo так, чтобы она выводила индекс и
* значение каждого аргумента по одному аргументу в строке.
*/

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Fprintf(os.Stdout, "%d %s\n", index, value)
	}
}
