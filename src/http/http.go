package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"create_at"`
	Body     string
}


type AiocRequest struct{
	 RequestUrl        string
	 RequestType       string
	 MaxResponseTime   int8
	 NumberTests       int8
	 TestInterNal      int8
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	urls := IssueURL + "?q=" + q
	resp, err := http.Get(urls)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("start encode ...")

	/*web, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("web:%s\n==============\n", web)
	*/
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issue: \n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#<%-5d> <%9.9s> <%.55s>\n\n", item.Number, item.User.Login, item.Title)
	}
}

