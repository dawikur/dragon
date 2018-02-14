
all: example
all: test
all: vet

example: example.go
	go build

test:
	go test ./...

vet:
	go vet ./...

clean:
	go clean

.PHONY: test
