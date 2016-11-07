package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
		pages := make(chan int)
		go fetchPages(url, pages)
		result := <-pages
		fmt.Println(result)
	} else if len(args) == 1 && args[0] != "allpages" {
		board = args[0]
		url = defaultConf.BaseUrl + board + "/index.html"
		fmt.Println("Fetch Single Pages " + url)
		fetchSingle(url)
	} else if len(args) == 3 {
		board = args[0]
		url = defaultConf.BaseUrl + board + "/index.html"
		fmt.Println("Fetch " + args[2] + " Pages " + url)
		pre, _ := strconv.Atoi(args[2])
		fetchMultiPages(defaultConf.BaseUrl, board, pre)
	} else {
		board = args[0]
		page := args[1]
		url = defaultConf.BaseUrl + board + "/index" + page + ".html"
		fmt.Println("Fetch single pages with page " + url)
		fetchSingle(url)
	}
}

func fetchSingle(url string) {
	resp := fetch(url)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	doc.Find("div.title").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		qHref, _ := a.Attr("href")
		title := strings.TrimSpace(s.Text())
		//ch <- title + "\t" + "https://www.ptt.cc" + qHref
		fmt.Println(title + "\t" + "https://www.ptt.cc" + qHref)
	})
}

func fetchPages(url string, ch chan int) {
	resp := fetch(url)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	href, _ := doc.Find("div.action-bar").Find("a.btn").Eq(3).Attr("href")

	pages, _ := strconv.Atoi(strings.Trim(strings.Split(strings.Split(href, "/")[3], ".")[0], "index"))
	ch <- pages + 1
}

func fetchMultiPages(baseUrl string, board string, pre int) {
	fmt.Println("Do fetch multiple pages")
	ch := make(chan int)
	url := baseUrl + board + "/index.html"
	go fetchPages(url, ch)
	p := <-ch

	// p, _ := strconv.Atoi(strings.Trim(strings.Split(pages, ".")[0], "index"))
	fmt.Println(p)
	var pagesURL = make([]string, pre+1)
	for i := pre; i >= 0; i-- {
		pagesURL[i] = baseUrl + board + "/index" + strconv.Itoa(p-i) + ".html"
		fmt.Println("\n" + pagesURL[i] + "\n")
		fetchSingle(pagesURL[i])
	}
}

func fetch(url string) *http.Response {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	return resp
}
