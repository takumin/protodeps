APPNAME := $(shell basename $(CURDIR))

ifeq (,$(shell git describe --abbrev=0 --tags 2>/dev/null))
VERSION := v0.0.0
else
VERSION := $(shell git describe --abbrev=0 --tags)
endif

ifeq (,$(shell git rev-parse --short HEAD 2>/dev/null))
REVISION := unknown
else
REVISION := $(shell git rev-parse --short HEAD)
endif

LDFLAGS_APPNAME  := -X "main.AppName=$(APPNAME)"
LDFLAGS_VERSION  := -X "main.Version=$(VERSION)"
LDFLAGS_REVISION := -X "main.Revision=$(REVISION)"
LDFLAGS          := -ldflags '-s -w $(LDFLAGS_APPNAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION) -extldflags -static'

SRCS := $(shell find . -type f -name '*.go')

.PHONY: all
all: build

.PHONY: build
build: $(APPNAME)
$(APPNAME): bin/$(APPNAME)
bin/$(APPNAME): $(SRCS)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $@

.PHONY: install
install: $(SRCS)
	CGO_ENABLED=0 go install $(LDFLAGS)

.PHONY: run
run: bin/$(APPNAME)
	bin/$(APPNAME) ensure

.PHONY: archive
archive: bin/$(APPNAME).zip
bin/$(APPNAME).zip: bin/$(APPNAME)
	cd bin && zip $@ $(APPNAME)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	staticcheck

.PHONY: release
release:
ifneq ($(GITHUB_TOKEN),)
	goreleaser release --rm-dist
endif

.PHONY: snapshot
snapshot:
	goreleaser release --rm-dist --snapshot

.PHONY: download
download:
	$(MAKE) -C tools download
	go mod download

.PHONY: tidy
tidy:
	$(MAKE) -C tools tidy
	go mod tidy

.PHONY: tools
tools:
	$(MAKE) -C tools install

.PHONY: clean
clean:
	rm -rf bin
	rm -rf dist
