###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

.PHONY: test

VERSION := 0.1.0

ver:
	@sed -i '' 's/^const Version = "[0-9]\{1,3\}.[0-9]\{1,3\}.[0-9]\{1,3\}"/const Version = "${VERSION}"/' audit.go

lint:
	~/go/bin/golint .

build:
	go build ./...

demo: build
	go build -o demo cmd/demo.go
	./demo

clean:
	rm -f demo

test: build
	go test -v ./...

github:
	open "https://github.com/mlavergn/goaudit"

release:
	zip -r goaudit.zip LICENSE README.md Makefile cmd *.go go.mod
	hub release create -m "${VERSION} - Go Audit" -a goaudit.zip -t master "v${VERSION}"
	open "https://github.com/mlavergn/goaudit/releases"
