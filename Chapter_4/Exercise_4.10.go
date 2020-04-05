/*
* Измените программу issues так, чтобы она сообщала о результатах
* с учетом их давности, деля на категории, например, поданные менее месяца назад,
* менее года назад и более года.
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
	"time"
)

const (
	IssuesURL = "https://api.github.com/search/issues"
)

const (
	LessThanMonth = "LessThanMonth"
	LessThanYear  = "LessThanYear"
	MoreThanYear  = "MoreThanYear"
)

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

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	issuesTimeCategories := make(map[string][]*Issue)

	now := time.Now()
	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.Before(day) && item.CreatedAt.After(month):
			issuesTimeCategories[LessThanMonth] = append(issuesTimeCategories[LessThanMonth], item)
		case item.CreatedAt.Before(month) && item.CreatedAt.After(year):
			issuesTimeCategories[LessThanYear] = append(issuesTimeCategories[LessThanYear], item)
		case item.CreatedAt.Before(year):
			issuesTimeCategories[MoreThanYear] = append(issuesTimeCategories[MoreThanYear], item)
		}
	}

	for category, issues := range issuesTimeCategories {
		fmt.Println(strings.ToUpper(category) + ":")
		for _, item := range issues {
			fmt.Printf("#%-5d %9.9s %.55s\t %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}
