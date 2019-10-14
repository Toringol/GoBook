/*
* Измените программу dup2 так, чтобы она выводила
* имена всех файлов, в которых найдены повторяющиеся строки.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filesWithDup := []string{}

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)

			for _, n := range counts {
				if n > 1 {
					filesWithDup = append(filesWithDup, f.Name())
				}
			}

			f.Close()
		}
		for _, fileName := range filesWithDup {
			fmt.Println(fileName)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
