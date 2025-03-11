all:
	make build
	make test
	make gofmt
	make vet
	make js

build:
	go build ./...

test:
	go test -v ./...

gofmt:
	gofmt -l -s -w ./...

js:
	(cd play && make)

vet:
	go vet ./...

diff:
	git diff --exit-code
