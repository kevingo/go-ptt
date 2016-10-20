// file: crawl.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"flag"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify board")
		os.Exit(1)
	}

	url := "https://www.ptt.cc/bbs/" + args[0] + "/index.html"
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
		fmt.Println(strings.TrimSpace(s.Text()) + "\thttps://www.ptt.cc" + qHref)
	})
}
