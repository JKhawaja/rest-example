default: build unit integration

repo: get build

get: 
	go get ./...	

build:
	go install -v ./...

#fmt:
#	diff -u <(echo -n) <(gofmt -s -d ./...)

# don't bother using this test command
test:
	go test -v ./...

integration:
	# Run:
	go test -v -tags integration ./...

	# Build:
	# go test -o ./test/behavior/integration/integration.test.exe -c --tags integration ./test/behavior/integration

# NOTE: unit testing requires a running instance of the server on localhost:8080
unit:
	go test -v -tags unit ./...

vet:
	go vet -v -x ./...
