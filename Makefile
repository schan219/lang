.PHONY: test

all:
	go build

test:
	go test ./test/parser
	go test ./pkg/stack