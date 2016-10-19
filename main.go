// file: crawl.go
package main

import (
	"fmt"
	// "github.com/jackdanger/collectlinks"
	"net/http"
	"io"
	// "io/ioutil"
  	"github.com/PuerkitoBio/goquery"
	// "golang.org/x/net/html"
	"strings"
)

func main() {

	boards := []string{"car", "nba"}

	for _, e := range boards {
		url := "https://www.ptt.cc/bbs/" + e + "/index.html"
		fetcher(url)
	}
}

func fetcher(url string) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()   

	doc, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	doc.Find("div.title").Each(func(i int, s *goquery.Selection) {
		fmt.Println(strings.TrimSpace(s.Text()))
	})

}
