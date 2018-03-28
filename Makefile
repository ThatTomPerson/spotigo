# note: call scripts from /scripts

OUTDIR ?= build/out
CMDDIR ?= cmd

PACKAGES = $(shell go list ./...)
CMDS = $(wildcard $(CMDDIR)/*/main.go)
BINS = $(patsubst $(CMDDIR)/%/main.go,$(OUTDIR)/%,$(CMDS))

GOPATH ?= $(HOME)/go
GOSRC ?= $(GOPATH)/src

# PROTOS = $(wildcard api/*.proto)
PROTOS = api/authentication.proto api/keyexchange.proto api/mercury.proto api/metadata.proto api/pubsub.proto api/spirc.proto
COMPILEDPROTOS = $(patsubst api/%.proto,internal/pkg/api/%,$(PROTOS))

PWD = $(shell pwd)

SRC = $(addsuffix /*.go,$(subst $(PWD)/,,$(addprefix $(GOSRC)/,$(PACKAGES))))

all: generate test $(BINS);

generate: $(COMPILEDPROTOS)
	@echo "Running go generate"
	@go generate $(PACKAGES)

$(BINS): $(SRC)
	@echo "Building $@ for $(GOOS):$(GOARCH)"
	@mkdir -p $(OUTDIR)
	@go build -race -v -o $@ $(patsubst $(OUTDIR)/%,$(CMDDIR)/%,$@)/*.go
	@file $@

test:
	@echo "Running go test -race"
	@go test -race $(PACKAGES)

$(PACKAGES):
	@echo "Running go test -race $@"
	@go test -race -c $@

internal/pkg/api/%: api/%.proto
	@mkdir -p $@
	protoc --go_out $@ -I api $<

