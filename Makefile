all: clean test build

build:
	go build ./client -o ./bin/client
	go build ./file -o ./bin/file
	go build ./resolver -o ./bin/resolver

test:
	go -v ./...

clean:
	go clean
	rm -rf ./bin
