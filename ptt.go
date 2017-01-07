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
	}
}

func fetchSingle(url string, str chan string) {
	resp := fetch(url)
	defer resp.Body.Close()

	doc := getDocument(resp)
	tmp := ""
	doc.Find("div.title").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		title := strings.TrimSpace(s.Text())
		if len(link) != 0 {
			tmp += title + "\t" + "https://www.ptt.cc" + link + "\n"
		}
	})
	str <- tmp
}

func fetchPages(url string, ch chan int) {
	resp := fetch(url)
	defer resp.Body.Close()

	doc := getDocument(resp)
	href, _ := doc.Find("div.action-bar").Find("a.btn").Eq(3).Attr("href")
	pages, _ := strconv.Atoi(strings.Trim(strings.Split(strings.Split(href, "/")[3], ".")[0], "index"))
	ch <- pages
}

func fetchMultiPages(board string, pre int) {
	fmt.Println("Do fetch multiple pages")
	ch := make(chan int)
	url := conf.BaseUrl + board + "/index.html"
	go fetchPages(url, ch)
	p := <-ch

	var pagesURL = make([]string, pre+1)
	for i := pre; i >= 0; i-- {
		pagesURL[i] = conf.BaseUrl + board + "/index" + strconv.Itoa(p+1-i) + ".html"
		fmt.Println("\n" + pagesURL[i] + "\n")
		s := make(chan string)
		go fetchSingle(pagesURL[i], s)
		result := <-s
		fmt.Println(result)
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
