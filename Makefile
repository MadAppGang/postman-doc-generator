COVERAGE_FILE = coverage.out

cover:
	go test -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE)
testall:
	go test ./...
testrun:
	go build
	cp postman-doc-generator ~/go/bin/
	cd test; go generate
