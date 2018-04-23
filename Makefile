SHELL := /bin/bash
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
EMACSBAK := $(shell find . -type f -name "*~")

GOCMD := go
GOYANGCMD := goyang
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
PROTOC := protoc
GOFMT := gofmt

FILEBODY := turing-machine

DATADIR := ./data
YANGFILE := $(DATADIR)/$(FILEBODY).yang
TTFXMLFILE := $(DATADIR)/$(FILEBODY)-config.xml
INITXMLFILE := $(DATADIR)/$(FILEBODY)-rpc.xml

PBDIR := ./proto
PBFILE := $(PBDIR)/$(FILEBODY).proto
PBFILEBASE := $(PBFILE).orig
PBTARGET := $(PBDIR)/$(FILEBODY).pb.go

CLIENTTGT := tmclient
CLIENTDIR := ./client
CLIENTSRC := $(wildcard $(CLIENTDIR)/*.go)
SERVERTGT := tmserver
SERVERDIR := ./server
SERVERSRC := $(wildcard $(SERVERDIR)/*.go)
TARGETS := $(CLIENTTGT) $(SERVERTGT)

# Rules

all: fmt $(TARGETS)

proto: $(PBTARGET)
$(PBTARGET): $(PBFILE)
	$(PROTOC) --proto_path=$(PBDIR) --go_out=plugins=grpc:$(PBDIR) $(PBFILE)
	protoc-go-inject-tag -input $(PBTARGET)

.PHONY: client
client: $(CLIENTTGT)
$(CLIENTTGT): client.go $(PBTARGET) $(CLIENTSRC)
	$(GOBUILD) -o $@ client.go

.PHONY: server
server: $(SERVERTGT)
$(SERVERTGT): server.go $(PBTARGET) $(SERVERSRC)
	$(GOBUILD) -o $@ server.go

goyang: $(YANGFILE)
	$(GOYANGCMD) --format=proto $(YANGFILE) > $(PBFILEBASE)

fmt: $(SRC)
	$(GOFMT) -l -w -e $(SRC)

.PHONY : clean
clean:
	$(GOCLEAN)
	rm -f $(EMACSBAK) $(PBTARGET) $(TARGETS)
