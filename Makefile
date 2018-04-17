SHELL := /bin/bash
TARGET := yttm
.DEFAULT_GOAL := $(TARGET)
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

$(TARGET): $(SRC)
	@go build -o $(TARGET)

clean:
	@rm $(TARGET)

fmt:
	@gofmt -l -w -e $(SRC)
