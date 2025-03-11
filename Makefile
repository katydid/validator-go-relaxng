.PHONY: build
build:
	go build ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	gofmt -l -s -w ./relaxng

.PHONY: js
genjs:
	(cd play && make)

.PHONY: testsuite
testsuite:
	(cd testsuite && ./simplify.sh)
	(cd testsuite && ./prettyprint.sh)

.PHONY: diff
diff:
	git diff --exit-code

.PHONY: checklicense
checklicense:
	checklicense ./relaxng
