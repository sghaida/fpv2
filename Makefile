VERSION=1.0.0

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

.PHONY: build clean-test test vet init coverage-badge

init:
	$(GOCMD) mod download

test: vet
	$(GOTEST)  -v ./src/...

clean-test:
	$(GOCLEAN) -testcache

test-race: clean-test
	$(GOTEST) --race -v ./src/...

vet:
	$(GOCMD) vet ./src/...

coverage-badge:
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic ./...
