/*
* Напишите вариативную функцию ElementByTagName, которая
* для данного дерева узла HTML и нуля или нескольких имен
* возвращает все элементы, которые соответствуют одному из этих
* имен. Вот два примера вызова такой функции:
* func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
* images := ElementsByTagName(doc, "img")
* headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
 */

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			return
		}

		names := []string{"img", "a"}

		items := ElementsByTagName(doc, names...)

		for _, value := range items {
			fmt.Println(value)
		}
	}
}

func ElementsByTagName(doc *html.Node, name ...string) (elements []*html.Node) {

	element := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, item := range name {
				if n.Data == item {
					elements = append(elements, n)
				}
			}
		}
	}

	forEachNode(doc, element, nil)

	return elements
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
