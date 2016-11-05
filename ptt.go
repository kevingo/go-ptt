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
	var board string

	if len(args) == 0 {
		url = defaultConf.CarBoardUrl
		fmt.Println("Fetch default car board")
		fetchSingle(url)
	} else if len(args) == 2 && args[1] == "allpages" {
		board = args[0]
		url = defaultConf.BaseUrl + board + "/index.html"
		fmt.Println("Fetch allpages " + url)
		fetchPages(url)
	} else if len(args) == 1 && args[0] != "allpages" {
		board = args[0]
		url = defaultConf.BaseUrl + board + "/index.html"
		fmt.Println("Fetch Single Pages " + url)
		fetchSingle(url)
	} else {
		board = args[0]
		page := args[1]
		url = defaultConf.BaseUrl + board + "/index" + page + ".html"
		fmt.Println("Fetch single pages with page " + url)
		fetchSingle(url)
	}
}

func fetchSingle(url string) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	doc.Find("div.title").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		qHref, _ := a.Attr("href")
		fmt.Println(strings.TrimSpace(s.Text()) + "\t" + "https://www.ptt.cc" + qHref)
	})
}

func fetchPages(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	href, _ := doc.Find("div.action-bar").Find("a.btn").Eq(3).Attr("href")
	pages := strings.Split(href, "/")[3]
	fmt.Println(pages)
}
