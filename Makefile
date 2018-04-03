# note: call scripts from /scripts

OUTDIR ?= build/out
CMDDIR ?= cmd

PACKAGES = $(shell go list ./...)
CMDS = $(wildcard $(CMDDIR)/*/main.go)



PLATFORMS := linux-amd64 windows-amd64 darwin-amd64 linux-arm

temp = $(subst -, ,$(word 3, $(subst /, ,$@)))
os = $(word 1, $(temp))
arch = $(word 2, $(temp))


outdirs := $(addprefix $(OUTDIR)/,$(PLATFORMS))

BINS = $(foreach platform,$(PLATFORMS),$(patsubst $(CMDDIR)/%/main.go,$(OUTDIR)/$(platform)-%,$(CMDS)))





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

# tmp:
# 	echo $(BINS)


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
	@echo "Building $@ for $(os):$(arch)"
	@mkdir -p $(OUTDIR)
	@GOOS=$(os) GOARCH=$(arch) go build -o $@ $(patsubst $(OUTDIR)/$(os)-$(arch)-%,$(CMDDIR)/%,$@)/*.go
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

run:
	go run cmd/spotifyd/*.go
.PHONY:run