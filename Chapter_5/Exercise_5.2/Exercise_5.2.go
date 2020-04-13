/*
* Напишите функцию для заполнения отображения, ключами которого
* являются имена элементов (p, div, span и тп), а значения -
* количество элементов с таким именем в дереве HTML-документа.
 */

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func CountElements(elements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	if n.FirstChild != nil {
		elements = CountElements(elements, n.FirstChild)
	}

	if n.NextSibling != nil {
		elements = CountElements(elements, n.NextSibling)
	}

	return elements
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountElements: %v\n", err)
		os.Exit(1)
	}

	for key, value := range CountElements(make(map[string]int), doc) {
		fmt.Println(key, " : ", value)
	}
}
