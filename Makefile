

.PHONY: all
all: help

.PHONY: help
help:
	@echo "make fmt - use gofmt & goimports"
	@echo "make lint - run golangci-lint"
	@echo "make test - run go test including race detection"
	@echo "make bench - run go test including benchmarking"
	@echo "make setup - install development libs"


.PHONY: fmt
fmt:
	$(info: Make: Format)
	gofmt -w ./**/*.go
	gofmt -w ./*.go
	goimports -w ./**/*.go
	goimports -w ./*.go
	golines -w ./**/*.go
	golines -w ./*.go

.PHONY: lint
lint:
	$(info: Make: Lint)
	@golangci-lint run --tests=false


.PHONY: test
test:
	CGO_ENABLED=1 go test -race ./...

.PHONY: bench
bench:
	CGO_ENABLED=1 go test -race ./...

.PHONY: setup
setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/segmentio/golines@latest