# go ptt 

Side project for crawling and parsing ptt webpages.

## Install

- go build

```
$ go build -o ptt *.go
```

- go install

```
go install github.com/kevingo/go-ptt
```

## Usage

- Fetch default `car` board data

```
$ go run *.go
```

- Fetch specific board

```
$ go run *.go -board {board}
```

- Fetch number of pages on specific board

```
$ go run *.go -board {board} -page {number}
```

## References
- [zhihu-go 源碼解析：用 goquery 解析 HTML](http://liyangliang.me/posts/2016/03/zhihu-go-insight-parsing-html-with-goquery/)
- [multiconfig](https://github.com/koding/multiconfig)
