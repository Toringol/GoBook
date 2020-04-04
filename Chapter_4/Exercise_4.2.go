/*
* Напишите программу, которая по умолчанию выводит дайджест
* SHA256 для входных данных, но при использовании соответствующих флагов
* командной строки выводит SHA384 или SHA512
 */

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var (
	shaFlag = flag.String("shaFlag", "sha256", "flag for sha256, sha384, sha512")
)

func main() {
	flag.Parse()

	var inputValue string
	fmt.Print("Input value that you want to crypt: ")
	fmt.Scan(&inputValue)

	if *shaFlag == "sha256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(inputValue)))
	} else if *shaFlag == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(inputValue)))
	} else if *shaFlag == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(inputValue)))
	} else {
		fmt.Println("Incorrect flag")
	}
}
