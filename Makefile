
all: example
all: test

example: example.go
	go build

test:
	go test ./...

clean:
	go clean

.PHONY: test
