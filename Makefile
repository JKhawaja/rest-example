.PHONY: build fmt test vet

default: fmt vet test

build:
	go build .

fmt:
	diff -u <(echo -n) <(gofmt -s -d ./...)

test:
	go test -v ./...

integration:
	go test -v --tags integration ./...

vet:
	go vet -x ./...
