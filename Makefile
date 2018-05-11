COVERAGE_FILE = coverage.out

cover:
	go test -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE)
testall:
	go test ./...
