EXECUTABLE :=golang-imgur
SOURCES ?= $(shell find . -name "*.go" -type f)
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

all: build

build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	go build -v -o bin/$@ ./

.PHONY: generate
generate:
	@which fileb0x > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/UnnoTed/fileb0x; \
	fi
	go generate $(PACKAGES)