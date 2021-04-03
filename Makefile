deps:
	go get -d -t ./...

test: deps
	go test -v

build: deps
	go build

lint:
	golint ./...
