build:
	go build -o ./bin/ulam ./cmd/ulam/ulam.go

test:
	go test ./...


all: build test
	./bin/ulam
	open output.png


ulam: build test
	./bin/ulam ulam
	open output.png

langton: build test
	./bin/ulam langton
	open output.png