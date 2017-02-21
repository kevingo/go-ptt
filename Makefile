BINARY=ptt
GOOS=darwin
GOARCH=amd64

install:build
	mv ${BINARY} ${GOBIN}

build:
	go build -o ${BINARY} *.go
