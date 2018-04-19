SHELL := /bin/bash
TARGET := yttm
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
PROTOC := protoc
GOFMT := gofmt
PBDIR := ./proto
PBTARGET := $(PBDIR)/turing-machine.proto

all: fmt build

protobuf: $(PBTARGET)
	$(PROTOC) --proto_path=$(PBDIR) --go_out=plugins=grpc:$(PBDIR) $(PBTARGET)

build: $(SRC) protobuf

fmt:
	$(GOFMT) -l -w -e $(SRC)

clean:
	$(GOCLEAN)
	@rm $(find . -type f -name "*~")
	@rm $(TARGET) 
