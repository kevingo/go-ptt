# go ptt 

Side project for crawling and parsing ptt webpages.

## Usage

- Fetch default `car` board data

```
$ go run ptt.go
```

- Fetch specific board

```
$ go run ptt.go -board {board}
```

- Fetch number of pages on specific board

```
$ go run ptt.go -board {board} -page {number}
```

## References
- [zhihu-go 源碼解析：用 goquery 解析 HTML](http://liyangliang.me/posts/2016/03/zhihu-go-insight-parsing-html-with-goquery/)
- [multiconfig](https://github.com/koding/multiconfig)
