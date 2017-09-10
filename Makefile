build:
	go install -v ./...

default: build unit integration

gen:
	goagen app -d github.com/JKhawaja/rest-example/ssot -o ./controllers
	goagen controller -d github.com/JKhawaja/rest-example/ssot -o ./controllers
	goagen swagger -d github.com/JKhawaja/rest-example/ssot -o ./docs
	goagen schema -d github.com/JKhawaja/rest-example/ssot -o ./docs
	goagen client -d github.com/JKhawaja/rest-example/ssot
	rm -r tool/
	go install ./controllers/app/ ./client/ ./controllers/
ifneq (ls -l | grep -v ^l | grep -l "goagen" | wc -l, 0)
	rm -r goagen*
endif

get: 
	go get ./...

#fmt:
#	diff -u <(echo -n) <(gofmt -s -d ./...)

integration:
	# Run:
	go test -v -tags integration ./...

	# Build:
	# go test -o ./test/behavior/integration/integration.test.exe -c --tags integration ./test/behavior/integration

repo: get build

# don't bother using this test command
test:
	go test -v ./...

# NOTE: unit testing requires a running instance of the server on localhost:8080
unit:
	go test -v -tags unit ./...

vet:
	go vet -v -x ./...
