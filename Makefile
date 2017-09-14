default: build unit benchmark

benchmark:

ifeq ($(shell ls -R test/performance/benchmark/ | grep \.benchmark$ | wc -l), 1)
	go test -tags=benchmark -benchtime=2s -bench=. ./... > ./test/performance/benchmark/current.benchmark
else ifeq ($(shell ls -R test/performance/benchmark/ | grep \.benchmark$ | wc -l), 2)
	go test -tags=benchmark -benchtime=2s -bench=. ./... > ./test/performance/benchmark/new.benchmark
else
	go test -tags=benchmark -benchtime=2s -bench=. ./... > ./test/performance/benchmark/new-$(shell date "+%F-%H-%M-%S").benchmark
endif

benchcmp:
	benchcmp -changed -mag ./test/performance/benchmark/current.benchmark ./test/performance/benchmark/new.benchmark

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
	go test -v -tags=integration -race ./...

	# Build:
	# go test -o ./test/behavior/integration/integration.test.exe -c --tags integration ./test/behavior/integration

load:
	# NOTE: must have instance of server running before performing load tests
	go test -v -tags=load -race ./...

repo: get build

# don't bother using this test command
test:
	go test -v ./...

	# Cross compile test binaries
	# GOOS=linux GOARCH=amd64 go test -c

unit:
	go test -v -tags=unit -race ./...

vet:
	go vet -v -x ./...
