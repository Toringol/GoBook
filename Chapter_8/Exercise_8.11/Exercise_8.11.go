/*
* Следуя подходу mirroredQuery из раздела 8.4.4, реализуйте
* вариант программы fetch, который параллельно запрашивает несколько URL.
* Как только получен первый ответ, прочие запросы отменяются.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var cancel = make(chan struct{})

func mirroredFetch(urls chan string) (err error) {
	responses := make(chan *http.Response, len(urls))

	go func() {
		for url := range urls {
			go func(url string) {
				req, err := http.NewRequest("GET", url, nil)
				req.Cancel = cancel
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return
				}
				responses <- resp
			}(url)
		}
	}()

	resp, err := ioutil.ReadAll((<-responses).Body)
	if err != nil {
		return err
	}
	close(cancel)

	err = ioutil.WriteFile("index.html", resp, 0644)

	return err
}

func main() {
	urls := os.Args[1:]
	urlsCh := make(chan string, len(urls))

	for _, url := range urls {
		urlsCh <- url
	}
	err := mirroredFetch(urlsCh)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
	}
}
