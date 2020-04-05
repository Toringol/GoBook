/*
* Популярный веб-ресурс с комиксами xkcd имеет интерфейс JSON.
* Например, запрос https://xkcd.com/571/info.0.json возвращает
* детальное описание комикса 571, одного из многочисленных фаворитов сайта.
* Загрузите каждый URL (по одному разу!) и постройте автономный список комиксов.
* Напишите программу xkcd, которая, используя этот список, будет выводить URL и
* описание каждого комикса, соответствующего условию поиска, заданному в командной строке.
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const urls = "urls.txt"

type ComicsInfo struct {
	ID         int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
}

func SearchComics(comics []ComicsInfo, url string) ([]ComicsInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result ComicsInfo

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	comics = append(comics, result)

	return comics, nil
}

func Print(comics []ComicsInfo) {
	for _, item := range comics {
		fmt.Printf("Comics number: %d\nTitle: %v\nAlt: %v\nPreview:\n%v\n",
			item.ID, item.Title, item.Alt, item.Transcript)
	}
}

func main() {

	comics := make([]ComicsInfo, 0)

	file, err := os.Open(urls)
	if err != nil {
		log.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		comics, err = SearchComics(comics, scanner.Text())
		if err != nil {
			log.Println(err)
			return
		}
	}

	Print(comics)
}
