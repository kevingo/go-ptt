# go ptt 

Read ptt in your console.

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
$ ptt
```

- Fetch specific board

```
$ ptt -board {board}
```

- Fetch number of pages on specific board

```
$ ptt -board {board} -page {number}
```
## Screenshot

![image](https://raw.githubusercontent.com/kevingo/go-ptt/master/screenshot/ptt-screenshot.png)
