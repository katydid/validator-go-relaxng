.PHONY: all
all:
	make gentests
	make build

.PHONY: build
build:
	go tool gopherjs build -m main.go

.PHONY: gentests
gentests:
	(cd ../relaxng && go test -v -run=TestGenJSSuite 2> ../play/tests.js)
