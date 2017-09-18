REST = github.com/JKhawaja/rest-example/ssot

default: bild unit bnch prof

bnch: benchmark-util benchmark-controllers

benchmark-util:
	go test -tags=benchmark -benchtime=2s -bench=. ./test/performance/benchmark/util/... > ./test/performance/benchmark/util/current.benchmark

benchmark-controllers:
	go test -tags=benchmark -benchtime=2s -bench=. ./test/performance/benchmark/controllers/... > ./test/performance/benchmark/controllers/current.benchmark

benchcmp: benchcmp-util benchcmp-controllers

benchcmp-util:
	go test -tags=benchmark -benchtime=2s -bench=. ./test/performance/benchmark/util/... > ./test/performance/benchmark/util/new.benchmark
	benchcmp -changed -mag ./test/performance/benchmark/util/current.benchmark ./test/performance/benchmark/util/new.benchmark > ./test/performance/benchmark/util/compare.benchmark

benchcmp-controllers:
	go test -tags=benchmark -benchtime=2s -bench=. ./test/performance/benchmark/controllers/... > ./test/performance/benchmark/controllers/new.benchmark
	benchcmp -changed -mag ./test/performance/benchmark/controllers/current.benchmark ./test/performance/benchmark/controllers/new.benchmark > ./test/performance/benchmark/controllers/compare.benchmark

prof: profile-util profile-controllers

profile-util:
	go test -tags=benchmark -benchtime=2s -bench=. -cpuprofile=./test/performance/profile/util/cpu.profile ./test/performance/benchmark/util
	go-torch -f=./test/performance/profile/profile.svg ./test/performance/profile/util/cpu.profile

profile-controllers:
	go test -tags=benchmark -benchtime=2s -bench=. -cpuprofile=./test/performance/profile/controllers/cpu.profile ./test/performance/benchmark/controllers
	go-torch -f=./test/performance/profile/controllers/profile.svg ./test/performance/profile/controllers/cpu.profile

bild:
	go install -v ./...

gen:
	goagen app -d $(REST) -o ./controllers
	goagen controller -d $(REST)  -o ./controllers
	goagen swagger -d $(REST)  -o ./docs
	goagen schema -d $(REST)  -o ./docs
	goagen client -d $(REST) 
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

repo: get bild

# don't bother using this test command
test:
	go test -v ./...

	# Cross compile test binaries
	# GOOS=linux GOARCH=amd64 go test -c

unit:
	go test -v -tags=unit -race ./...

vet:
	go vet -v -x ./...
