default: vet build unit integration

build:
	go install -v ./...

#fmt:
#	diff -u <(echo -n) <(gofmt -s -d ./...)

# don't bother using this test command

test:
	go test -v ./...

integration:
	go test -v --tags integration ./...

unit:
	go test -v --tags unit ./...

vet:
	go vet -v -x ./...
