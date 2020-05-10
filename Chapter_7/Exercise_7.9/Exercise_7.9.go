/*
* Воспользуйтесь пакетом html/template (раздел 4.6) для замены
* printTracks функцией, которая выводит дорожки в виде таблицы HTML.
* Используйте решение предыдущего упражнения, чтобы каждый щелчок на
* заголовке столбца генерировал HTTP-запрос на сортировку таблицы.
 */

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(writer io.Writer, tracks []*Track) {
	if err := trackTable.Execute(writer, tracks); err != nil {
		log.Fatal(err)
	}
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var trackTable = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
	  <meta charset="utf-8">
		<style media="screen" type="text/css">
		  table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
		  table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<h1>Tracks</h1>
		<table>
		  <thead>
				<tr>
					<th><a href="/?sort=Title">Title</a></th>
					<th><a href="/?sort=Artist">Artist</a></th>
					<th><a href="/?sort=Album">Album</a></th>
					<th><a href="/?sort=Year">Year</a></th>
					<th><a href="/?sort=Length">Length</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	sortBy := r.FormValue("sort")

	sort.Sort(customSort{tracks, func(x, y *Track) bool {

		switch sortBy {
		case "Title":
			if x.Title != y.Title {
				return x.Title < y.Title
			}
		case "Artist":
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
		case "Album":
			if x.Album != y.Album {
				return x.Album < y.Album
			}
		case "Year":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		case "Length":
			if x.Length != y.Length {
				return x.Length < y.Length
			}
		}

		return false
	}})

	printTracks(w, tracks)
}
