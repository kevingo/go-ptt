package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/koding/multiconfig"
	"github.com/olekukonko/tablewriter"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	doc.Find("div.r-ent").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		title := s.Find("a").Text()
		push := s.Find("span").Text()
		if push == "" {
			push = "X"
		}

		if len(link) != 0 {
			tmp += push + "\t" + title + "\t" + "https://www.ptt.cc" + link + "\n"
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
	ch := make(chan int)
	url := conf.BaseUrl + board + "/index.html"
	go fetchPages(url, ch)
	p := <-ch

	var output string

	var pagesURL = make([]string, pre+1)
	for i := pre; i >= 0; i-- {
		pagesURL[i] = conf.BaseUrl + board + "/index" + strconv.Itoa(p+1-i) + ".html"
		s := make(chan string)
		go fetchSingle(pagesURL[i], s)
		result := <-s
		output += result
	}

	printOutput(output)
}

func printOutput(output string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Push", "Title", "URL"})
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	o := [][]string{}

	rows := strings.Split(output, "\n")
	for i := 0; i < len(rows); i++ {
		row := strings.Split(rows[i], "\t") // row[0], row[1], row[2]
		if len(row) == 3 {
			arr := []string{row[0], row[1], row[2]}
			o = append(o, arr)
		}
	}

	for _, v := range o {
		table.Append(v)
	}

	table.Render() // Send output
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
