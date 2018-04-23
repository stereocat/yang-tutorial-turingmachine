SHELL := /bin/bash
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
TTFXMLFILE := $(FILEBODY)-config.xml
INITXMLFILE := $(FILEBODY)-rpc.xml
PBFILE := $(PBDIR)/$(FILEBODY).proto
PBFILEBASE := $(PBFILE).orig
PBTARGET := $(PBDIR)/$(FILEBODY).pb.go
EMACSBAK := $(shell find . -type f -name "*~")
CLIENTEXE := tm_client
CLIENTDIR := ./client
CLIENTSRC := $(wildcard $(CLIENTDIR)/*.go)
SERVEREXE := tm_server
SERVERDIR := ./server
SERVERSRC := $(wildcard $(SERVERDIR)/*.go)

all: fmt $(CLIENTEXE) $(SERVEREXE)

$(PBTARGET): $(PBFILE)
	$(PROTOC) --proto_path=$(PBDIR) --go_out=plugins=grpc:$(PBDIR) $(PBFILE)
	protoc-go-inject-tag -input $(PBTARGET)

$(CLIENTEXE): client.go $(PBTARGET) $(CLIENTSRC)
	$(GOCMD) build -o $@ client.go

$(SERVEREXE): server.go $(PBTARGET) $(SERVERSRC)
	$(GOCMD) build -o $@ server.go

protobuf: $(PBTARGET)

goyang: $(YANGFILE)
	$(GOYANGCMD) --format=proto $(YANGFILE) > $(PBFILEBASE)

fmt:
	$(GOFMT) -l -w -e $(SRC)

.PHONY : clean
clean:
	$(GOCLEAN)
	rm -f $(EMACSBAK) $(TARGET) $(PBTARGET) $(SERVEREXE) $(CLIENTEXE)
