/*
* Добавьте ограничение по глубине в параллельный сканер. Иначе
* говоря, если пользователь устанавливает -depth=3, то
* выбираются только те URL, которые достижимы через цепочку
* не более чем из трех ссылок.
 */

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/adonovan/gopl.io/ch5/links"
)

var maxdepth int
var args []string

type Work struct {
	url   string
	depth int
}

func init() {
	flag.IntVar(&maxdepth, "depth", 3, "max depth to crawl")
	flag.Parse()
	args = flag.Args()
}

var tokens = make(chan struct{}, 20)

func crawl(work Work) []Work {
	fmt.Printf("%d\t%s\n", work.depth, work.url)

	if work.depth >= maxdepth {
		return nil
	}

	tokens <- struct{}{}
	list, err := links.Extract(work.url)
	<-tokens

	if err != nil {
		log.Print(err)
	}

	works := []Work{}
	for _, link := range list {
		works = append(works, Work{link, work.depth + 1})
	}
	return works
}

func main() {
	worklist := make(chan []Work)
	var n int

	n++
	go func() {
		works := []Work{}
		for _, url := range args {
			works = append(works, Work{url, 1})
		}
		worklist <- works
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				n++
				go func(link Work) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
