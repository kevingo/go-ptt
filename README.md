# go ptt 

Side project for crawling and parsing ptt webpages.

## Usage

- Fetch default car board data

```
$ go run ptt.go
```

- Fetch specific board

```
$ go run ptt.go {board}
```

- Fetch specific board with pages

```
$ go run ptt.go {board} allpages {index}
```

- Fetch number of pages on specific board

```
$ go run ptt.go {board} allapges
```

## References
- [zhihu-go 源碼解析：用 goquery 解析 HTML](http://liyangliang.me/posts/2016/03/zhihu-go-insight-parsing-html-with-goquery/)
- [multiconfig](https://github.com/koding/multiconfig)
