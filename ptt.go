package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koding/multiconfig"
)

var conf = loadConfig()

func loadConfig() *DefaultConf {
	m := multiconfig.New()
	conf := new(DefaultConf)
	m.MustLoad(conf)

	return conf
}

func main() {

	switch conf.Mode {
	case "view":
		fetchMultiPages(conf.Board, conf.Page)
	case "crawl":
		fmt.Println("Will support crawl later on")
	}
}

func fetchSingle(url string) {
	resp := fetch(url)
	defer resp.Body.Close()

	doc := getDocument(resp)
	doc.Find("div.title").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		title := strings.TrimSpace(s.Text())
		if len(link) != 0 {
			fmt.Println(title + "\t" + "https://www.ptt.cc" + link)
		}
	})
}

func fetchPages(url string, ch chan int) {
	resp := fetch(url)
	defer resp.Body.Close()

	doc := getDocument(resp)
	href, _ := doc.Find("div.action-bar").Find("a.btn").Eq(3).Attr("href")
	pages, _ := strconv.Atoi(strings.Trim(strings.Split(strings.Split(href, "/")[3], ".")[0], "index"))
	ch <- pages + 1
}

func fetchMultiPages(board string, pre int) {
	fmt.Println("Do fetch multiple pages")
	ch := make(chan int)
	url := conf.BaseUrl + board + "/index.html"
	go fetchPages(url, ch)
	p := <-ch

	var pagesURL = make([]string, pre+1)
	for i := pre; i >= 0; i-- {
		pagesURL[i] = conf.BaseUrl + board + "/index" + strconv.Itoa(p-i) + ".html"
		fmt.Println("\n" + pagesURL[i] + "\n")
		fetchSingle(pagesURL[i])
	}
}

func getDocument(resp *http.Response) *goquery.Document {
	r := io.Reader(resp.Body)
	doc, _ := goquery.NewDocumentFromReader(r)
	return doc
}

func fetch(url string) *http.Response {
	resp, _ := http.Get(url)
	return resp
}
