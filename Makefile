# note: call scripts from /scripts

OUTDIR ?= dist
CMDDIR ?= cmd

PACKAGES = $(shell go list ./...)
CMDS = $(wildcard $(CMDDIR)/*)
BINS = $(patsubst $(CMDDIR)/%,$(OUTDIR)/%,$(CMDS))

GOPATH ?= $(HOME)/go
GOSRC ?= $(GOPATH)/src

PWD = $(shell pwd)

SRC = $(addsuffix /*.go,$(subst $(PWD)/,,$(addprefix $(GOSRC)/,$(PACKAGES))))

all: test $(BINS);

generate:
	@echo "Running go generate"
	@go generate $(PACKAGES)

$(BINS): generate $(SRC)
	@echo "Building $@ for $(GOOS):$(GOARCH)"
	@go build -race -v -o $@ $(patsubst $(OUTDIR)/%,$(CMDDIR)/%,$@)/*.go

test:
  @echo "Running go test -race"
	@go test -race $(PACKAGES)
