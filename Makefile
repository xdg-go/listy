GENERATED=$(shell grep -l "DO NOT EDIT" *.go | grep -v generator)

all: source test vet lint

source: *.go

$(GENERATED) : I8.go I8_test.go generator.go
	@go generate

.PHONY: test
test:
	@go test .

.PHONY: cover
cover:
	@go test -cover .

.PHONY: vet
vet:
	@go tool vet .

.PHONY: lint
lint:
	@golint -set_exit_status .

