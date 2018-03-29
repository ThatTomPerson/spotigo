# note: call scripts from /scripts

OUTDIR ?= build/out
CMDDIR ?= cmd

PACKAGES = $(shell go list ./...)
CMDS = $(wildcard $(CMDDIR)/*/main.go)
BINS = $(patsubst $(CMDDIR)/%/main.go,$(OUTDIR)/%,$(CMDS))

GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
GOSRC ?= $(GOPATH)/src

GODERIVE = $(GOBIN)/goderive
GEN = $(GOBIN)/gen
DEP = $(GOBIN)/dep

# PROTOS = $(wildcard api/*.proto)
PROTOS = api/authentication.proto api/keyexchange.proto api/mercury.proto api/metadata.proto api/pubsub.proto api/spirc.proto
COMPILEDPROTOS = $(patsubst api/%.proto,internal/pkg/api/%,$(PROTOS))

PWD = $(shell pwd)

ROOT = $(subst $(GOSRC)/,,$(PWD))

SRC = $(addsuffix /*.go,$(subst $(PWD)/,,$(addprefix $(GOSRC)/,$(PACKAGES))))

DERIVED = $(addsuffix /derived.gen.go,$(subst $(PWD)/,,$(addprefix $(GOSRC)/,$(PACKAGES))))


all: generate vendor test $(BINS);

generate: $(COMPILEDPROTOS) $(GODERIVE)
	@echo Running go generate
	@go generate $(PACKAGES)

$(GODERIVE):
	go get github.com/awalterschulze/goderive

$(GEN):
	go get github.com/clipperhouse/gen

$(DEP):
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

$(BINS): $(SRC)
	@echo "Building $@ for $(GOOS):$(GOARCH)"
	@mkdir -p $(OUTDIR)
	@go build -race -v -o $@ $(patsubst $(OUTDIR)/%,$(CMDDIR)/%,$@)/*.go
	@file $@

test: $(SRC)
	@echo "Running go test -race"
	@go test -race $(PACKAGES)
.PHONY: test

vendor: $(DEP) Gopkg.toml Gopkg.lock
	$(DEP) ensure

$(PACKAGES):
	@echo "Running go test -race $@"
	@go test -race -c $@

internal/pkg/api/%: api/%.proto
	@mkdir -p $@
	protoc --go_out $@ -I api $<

