
BINARY_NAME=crservice
COVERAGE_FILE = coverage.out
COVERAGE_HTML = coverage.html

all: clean lint test build

build:
	go build -o $(BINARY_NAME) app/main.go

run: build
	./..$(BINARY_NAME)

test:
	go test ./... -v

lint:
	golangci-lint run

clean:
	rm -f $(BINARY_NAME) $(COVERAGE_FILE) $(COVERAGE_HTML)

coverage-html: test
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

coverage:
	go tool cover -func=$(COVERAGE_FILE)

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
	go install github.com/vektra/mockery/v2@latest


.PHONY: all build run test lint clean
