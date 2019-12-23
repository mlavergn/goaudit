###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

.PHONY: test

GOPATH = "${PWD}"

lint:
	GOPATH=${GOPATH} ~/go/bin/golint .

build:
	GOPATH=${GOPATH} go build ./...

demo: build
	GOPATH=${GOPATH} go build -o demo cmd/demo.go
	./demo

clean:
	rm -f demo

test: build
	GOPATH=${GOPATH} go test -v ./src/...

github:
	open "https://github.com/mlavergn/goaudit"