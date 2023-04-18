VERSION=1.0.0

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

.PHONY: build test clean init coverage

init:
	$(GOCMD) mod download

test:
	$(GOTEST)  -v ./src/...

test-race:
	$(GOTEST) --race -v ./src/...

coverage:
	$(GOTEST)  ./src/... -coverprofile cover.out
	go tool cover -html=./src/cover.out

coverage-badge:
	gopherbadger -md="README.md"