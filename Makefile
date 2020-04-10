.PHONY: all
all: build

.PHONY: ci
ci: build gofmt govet golangci test coverage

.PHONY: build
build:
	go build

.PHONY: gofmt
gofmt:
	gofmt -l .

.PHONY: govet
govet:
	go vet -all ./...

.PHONY: golangci
golangci:
	golangci-lint run --timeout 30m -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: coverage
coverage:
	bash scripts/coverage.bash
