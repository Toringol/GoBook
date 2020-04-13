/*
* Измените функцию forEachNode так, чтобы функции pre и
* post возвращали булево значение, указывающее, следует ли продолжать обход
* дерева. Воспользуйтесь ими для написания функции ElementID с приведенной
* ниже сигнатурой, которая находит первый HTML-элемент с указанным атрибутом
* id. Функция должна прекращать обход дерева, как только соответствующий
* элемент найден:
* func ElementByID(doc *html.Node, id string) *html.Node
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(os.Stdin)
		id, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", ElementByID(doc, id))
	}
}

func ElementByID(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}

		}
		return true
	}
	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0) // unvisited
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}
