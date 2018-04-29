SHELL := /bin/bash
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
EMACSBAK := $(shell find . -type f -name "*~")

GOCMD := go
GOYANGCMD := goyang
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
PROTOC := protoc
GOFMT := gofmt
GOLINT := golint

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
CLIENTSRC := $(CLIENTTGT).go $(wildcard $(CLIENTDIR)/*.go) $(wildcard $(PBDIR)/*.go)
SERVERTGT := tmserver
SERVERDIR := ./server
SERVERSRC := $(SERVERTGT).go $(wildcard $(SERVERDIR)/*.go) $(wildcard $(PBDIR)/*.go)
TARGETS := $(CLIENTTGT) $(SERVERTGT)

# Rules

all: fmt $(TARGETS)

proto: $(PBTARGET)
$(PBTARGET): $(PBFILE)
	$(PROTOC) --proto_path=$(PBDIR) --go_out=plugins=grpc:$(PBDIR) $(PBFILE)
	protoc-go-inject-tag -input $(PBTARGET)

.PHONY: client
client: $(CLIENTTGT)
$(CLIENTTGT): $(CLIENTTGT).go $(PBTARGET) $(CLIENTSRC)
	$(GOBUILD) -o $@ $(CLIENTTGT).go

.PHONY: server
server: $(SERVERTGT)
$(SERVERTGT): $(SERVERTGT).go $(PBTARGET) $(SERVERSRC)
	$(GOBUILD) -o $@ $(SERVERTGT).go

goyang: $(YANGFILE)
	$(GOYANGCMD) --format=proto $(YANGFILE) > $(PBFILEBASE)

fmt: $(SRC)
	$(GOFMT) -l -w -e $(SRC)

lint: $(SRC)
	$(GOLINT) $(SERVERDIR)
	$(GOLINT) $(CLIENTDIR)
	$(GOLINT) $(PBDIR)/$(FILEBODY).go

.PHONY : clean
clean:
	$(GOCLEAN)
	rm -f $(EMACSBAK) $(PBTARGET) $(TARGETS)
