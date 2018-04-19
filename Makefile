SHELL := /bin/bash
TARGET := yttm
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOFMT := gofmt

all: fmt build

build: $(SRC)
	$(GOBUILD) -o $(TARGET)

fmt:
	$(GOFMT) -l -w -e $(SRC)

clean:
	$(GOCLEAN)
	@rm $(find . -type f -name "*~")
	@rm $(TARGET) 

