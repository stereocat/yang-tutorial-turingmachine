SHELL := /bin/bash
TARGET := yttm
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOCMD := go
GOYANGCMD := goyang
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
PROTOC := protoc
GOFMT := gofmt
PBDIR := ./proto
FILEBODY := turing-machine
YANGFILE := $(FILEBODY).yang
PBFILE := $(PBDIR)/$(FILEBODY).proto
PBFILEBASE := $(PBFILE).orig
PBTARGET := $(PBDIR)/$(FILEBODY).pb.go
EMACSBAK := $(shell find . -type f -name "*~")

all: fmt protobuf

client: client.go protobuf
	$(GOCMD) run client.go -t $(FILEBODY)-config.xml

server: server.go protobuf
	$(GOCMD) run server.go

goyang: $(YANGFILE)
	$(GOYANGCMD) --format=proto $(YANGFILE) > $(PBFILEBASE)

protobuf: $(PBFILE)
	$(PROTOC) --proto_path=$(PBDIR) --go_out=plugins=grpc:$(PBDIR) $(PBFILE)
	protoc-go-inject-tag -input $(PBTARGET)

fmt:
	$(GOFMT) -l -w -e $(SRC)

.PHONY : clean
clean:
	$(GOCLEAN)
	rm -f $(EMACSBAK) $(TARGET) $(PBTARGET)
