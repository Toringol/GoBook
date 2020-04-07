/*
* Создайте веб-серве, однократно запрашивающий информацию
* у Github, а затем позволяющий выполнять навигацию по списку сообщений
* об ошибках, контрольных точек и пользователей.
 */

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

var issueList = template.Must(template.New("issuelist").Parse(`
 <h1>{{.TotalCount}} issues</h1>
 <table>
 <tr style='text-align: left'>
   <th>#</th>
   <th>State</th>
   <th>User</th>
   <th>Title</th>
 </tr>
 {{range .Items}}
 <tr>
   <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
   <td>{{.State}}</td>
   <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
   <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
 </tr>
 {{end}}
 </table>
 `))

func issueHandler(w http.ResponseWriter, r *http.Request) {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", issueHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
