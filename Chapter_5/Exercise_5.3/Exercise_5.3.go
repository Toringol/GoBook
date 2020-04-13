/*
* Напишите функцию для вывода содержимого всех текстовых
* узлов в дереве документа HTML. Не входите в элементы <script>
* и <style>, поскольку их содержимое в веб браузере не является видимым.
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func TextElements(textElements []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		text := strings.Trim(n.Data, " \n")
		if text != "" {
			textElements = append(textElements, text)
		}
	}

	if n.FirstChild != nil && n.Data != "script" && n.Data != "style" {
		textElements = TextElements(textElements, n.FirstChild)
	}

	if n.NextSibling != nil && n.Data != "script" && n.Data != "style" {
		textElements = TextElements(textElements, n.NextSibling)
	}

	return textElements
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "TextElements: %v\n", err)
		os.Exit(1)
	}

	for key, value := range TextElements(nil, doc) {
		fmt.Println(key, " : ", value)
	}
}
