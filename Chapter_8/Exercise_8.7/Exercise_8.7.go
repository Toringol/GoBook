/*
* Напишите параллельную программу, которая создает локальное
* зеркало веб-сайта, загружая все доступные страницы и записывая их
* в каталог на локальном диске. Выбираться должны только страницы
* в переделах исходного домена (например, golang.org). URL в страницах
* зеркала должны при необходимости быть изменены таким образом, чтобы
* они ссылались на зеркальную страницу, а не на оригинал.
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adonovan/gopl.io/ch5/links"
)

var (
	tokens   = make(chan struct{}, 20)
	filePath = "./out/"
)

func crawl(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath + strings.ReplaceAll(url, "/", ".") + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func sameDomain(domain string, url string) bool {
	return strings.HasPrefix(url, domain)
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		startLink := list[0]
		for _, link := range list {
			if !seen[link] && sameDomain(startLink, link) {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
