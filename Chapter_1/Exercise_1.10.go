/*
* Найдите веб-сайт, который содержит большое количество данных.
* Исследуйте работу кэширования путем двукратного запуска fetchall и
* сравнение времени запросов. Получаете ли вы каждый раз одно и то же содержимое?
* Измените fetchall так, чтобы вывод осуществлялся в файл и чтобы затем можно
* было его изучить.
 */

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0600)
		defer file.Close()
		if err != nil {
			fmt.Println("Error opening file")
		}
		file.WriteString(<-ch + "\n")
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
