GO ?= go
GOPATH := $(CURDIR)/:$(GOPATH)

all: build

build:
	$(GO) build -o gobitt cmd/gobitt/main.go
install:
	cp gobitt /usr/bin/
