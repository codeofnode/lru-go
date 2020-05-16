#!make

# Dev
.PHONY: dev

dev:
	@bash scripts/watch.sh test
##

# format
.PHONY: fmt

fmt:
	@go fmt ./...
#

# Test
.PHONY: test

test:
	@go test ./...
##
