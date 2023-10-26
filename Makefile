build:
	go build

test:
	go test


display: build test
	./ulam
	open output.png