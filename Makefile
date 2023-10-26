build:
	go build

test:
	go test ./...


all: build test
	./ulam
	open output.png


ulam: build test
	./ulam ulam
	open output.png

langton: build test
	./ulam langton
	open output.png