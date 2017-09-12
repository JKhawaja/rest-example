SYSTEM_DATE = $(SYSTEM_DATE)/$(shell date)

default: build unit integration benchmark

benchmark:
	go test -tags=benchmark -bench=. ./... > ./test/performance/benchmark/$(shell date "+%F-%H-%M-%S").benchmark

build:
	go install -v ./...

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

fuzzbuild:
	go-fuzz-build -func=FuzzRemoveDuplicates -o ./test/behavior/fuzz/removeDuplicates.zip github.com/JKhawaja/rest-example/test/behavior/fuzz
	go-fuzz-build -func=FuzzNameVerification -o ./test/behavior/fuzz/nameVerification.zip github.com/JKhawaja/rest-example/test/behavior/fuzz

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

	# Cross compile test binaries
	# GOOS=foo GOARCH=bar go test -c

unit:
	go test -v -tags unit ./...

vet:
	go vet -v -x ./...
