// file: crawl.go
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

	queue := make(chan string)

	go func() {
		queue <- board
	}()

	for uri := range queue {
		enqueue(uri, queue)
	}

	// url := getIndex(board)
	// fetcher(url)
}

func enqueue(uri string, queue chan string) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := []string{"https://www.ptt.cc/bbs/car/index3263.html", "https://www.ptt.cc/bbs/car/index3262.html"}

	for _, link := range links {
		go func() { queue <- link }()
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
		a := s.Find("a")
		qHref, _ := a.Attr("href")
		fmt.Println(strings.TrimSpace(s.Text()) + "\t" + getHome() + qHref)
	})
}

func getTestIndex() string {
	return
}

func getIndex(board string) string {
	return "https://www.ptt.cc/bbs/" + board + "/index.html"
}

func getHome() string {
	return "https://www.ptt.cc/"
}
