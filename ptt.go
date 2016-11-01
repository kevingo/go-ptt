package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koding/multiconfig"
)

type DefaultConf struct {
	Board       string
	BaseUrl     string
	CarBoardUrl string
}

func main() {
	m := multiconfig.NewWithPath("config.toml")
	defaultConf := new(DefaultConf)
	m.MustLoad(defaultConf)

	flag.Parse()
	args := flag.Args()
	var url string

	if len(args) == 0 {
		url = defaultConf.CarBoardUrl
	} else {
		board := args[0]
		page := args[1]
		url = defaultConf.BaseUrl + board + "/index" + page + ".html"
	}

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
