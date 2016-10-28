package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify board")
		os.Exit(1)
	}

	board := args[0]
	page := args[1]

	url := getIndex(board, page)
	fetcher(url)
}

func fetcher(url string) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	doc.Find("div.title").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		qHref, _ := a.Attr("href")
		fmt.Println(strings.TrimSpace(s.Text()) + "\t" + getHome() + qHref)
	})
}

func getIndex(board string, page string) string {
		
	url := "https://www.ptt.cc/bbs/" + board + "/index.html"

	if page != "0" {
		url = "https://www.ptt.cc/bbs/" + board + "/index" + page + ".html"
	}

	return url;
}

func getHome() string {
	return "https://www.ptt.cc/"
}
