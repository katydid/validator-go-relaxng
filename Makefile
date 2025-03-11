all:
	make build
	make test
	make gofmt
	make vet
	make js

build:
	go build ./...

vet:
	go vet ./...

test:
	go test -v ./...

gofmt:
	gofmt -l -s -w ./...

js:
	(cd play && make)

testsuite:
	(cd testsuite && ./simplify.sh)
	(cd testsuite && ./prettyprint.sh)

diff:
	git diff --exit-code
