package main

type DefaultConf struct {
	Board       string `default:"car"`
	BaseUrl     string `default:"https://www.ptt.cc/bbs/"`
	CarBoardUrl string `default:"https://www.ptt.cc/bbs/car/index.html"`
	Page        int    `default:"2"`
	Mode        string `default:"view"`
}
