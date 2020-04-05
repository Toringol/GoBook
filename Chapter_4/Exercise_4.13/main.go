/*
* Веб-служба Open Movie Database https://omdbapi.com/ на
* базе JSON позволяет выполнять поиск фильма по названию и
* загружать его афишу. Напишите инструмент poster, который
* загружает афишу фильма по переданному в командной строке названию.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	URL    = "http://www.omdbapi.com/"
	apikey = "6df36eee"
)

type FilmInfo struct {
	Title    string
	Year     string
	Director string
	Writer   string
	Actors   string
	Language string
	Country  string
	Plot     string
}

func Poster(terms []string) (*FilmInfo, error) {
	filmName := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(URL + "?t=" + filmName + "&apikey=" + apikey)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result FilmInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	return &result, nil
}

func PrintFilmInfo(filmInfo *FilmInfo) {
	fmt.Printf("Title: %v\nYear: %v\nDirector: %v\nWriter: %v\nActors: %v\nLanguage: %v\nCountry: %v\nPlot: %v\n",
		filmInfo.Title, filmInfo.Year, filmInfo.Director, filmInfo.Writer, filmInfo.Actors, filmInfo.Language, filmInfo.Country, filmInfo.Plot)
}

func main() {
	filmInfo, err := Poster(os.Args[1:])
	if err != nil {
		log.Println(err)
		return
	}

	PrintFilmInfo(filmInfo)

}
