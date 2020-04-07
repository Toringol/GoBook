/*
* Создайте инструмет, который позволит пользователю создавать,
* читать, обновлять и закрывать темы GitHub из командной строки, вызывая
* предпочитаемый пользователем текстовый редактор, когда требуется ввести
* текст значительного размера.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

const APIURL = "https://api.github.com"

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

func create(owner, repo, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}
}

func read(owner, repo, number string) {
	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}
	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func edit(owner, repo, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	_, err = EditIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func close(owner, repo, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "closed"})
	if err != nil {
		log.Fatal(err)
	}
}

func HowToUse() {
	fmt.Println("Usage: cmd{create|read|edit|close} owner repo number")
}

func main() {
	if len(os.Args) < 4 {
		HowToUse()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if len(args) != 3 {
		HowToUse()
		return
	}

	owner, repo, number := args[0], args[1], args[2]

	switch cmd {
	case "create":
		create(owner, repo, number)
	case "read":
		read(owner, repo, number)
	case "edit":
		edit(owner, repo, number)
	case "close":
		close(owner, repo, number)
	}

}
