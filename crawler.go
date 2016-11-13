package main

import (
	"fmt"
	"github.com/koding/multiconfig"
)

type Config struct {
	Board string
}

var conf = loadConfig()

func loadConfig() *Config {
	m := multiconfig.New()
	conf := new(Config)
	m.MustLoad(conf)

	return conf
}

func main() {
	fmt.Println("Run ptt crawler with board - ", conf.Board)
}

func fetch(url string) *http.Response {
	resp, _ := http.Get(url)
	return resp
}