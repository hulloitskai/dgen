### Variables:
PKG_NAME = $(shell basename "$$(pwd)")

## Source config
SRC_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
SRC_PKGS = $(shell go list ./... | grep -v /vendor/)

## Testing config
TEST_TIMEOUT = 20s
COVER_OUT = coverage.txt


### Commands (targets):
## Prevent targeting filenames...
.PHONY: default run build all all-bench check fmt test test-v test-race bench

## Default target when no arguments are given to make (build and run program).
default: build run

## Builds and runs the program (package must be main).
run:
	@if [ -f ".env.sh" ]; then \
		printf 'Exporting environment variables by sourcing ".env.sh"... ' && \
		. .env.sh && \
		echo "done!"; \
	fi; \
	if [ -f "$(PKG_NAME)" ]; then \
		echo 'Running "$(PKG_NAME)"...' && \
		./$(PKG_NAME); \
	else echo '[ERROR] Could not find program "$(PKG_NAME)".'; \
	fi

## Builds the program specified by the main package.
build:
	@printf "Building... "
	@go build
	@printf "done!\n"

## Formats, checks, and tests the code.
review: fmt check test
## Like "review", but tests for race conditions.
review-race: fmt check test-race
## Like "review-race", but includes benchmarks.
review-bench: fmt check test-race bench

## Checks for formatting, linting, and suspicious code.
check:
## Check formatting...
	@printf "Check format:"
	@GOFMT_OUT="$(shell gofmt -l $(SRC_FILES))"; if [ -n "$$GOFMT_OUT" ]; then \
		printf "\n> [WARN] Fix formatting issues in the following files with "; \
		printf '"make fmt":\n$$GOFMT_OUT\n\n'; else printf " ...OK!\n"; fi
## Lint files...
	@printf "Check lint:"
	@GOLINT_OUT="$(shell for PKG in "$(SRC_PKGS)"; do golint $$PKG; done)"; \
		if [ -n "$$GOLINT_OUT" ]; then printf "\n" && \
		for PKG in "$$GOLINT_OUT"; do printf "> $$PKG\n"; done; printf "\n"; \
		else printf " ...OK!\n"; fi
## Check suspicious code...
	@printf "Check vet:"
	@GOVET_OUT="$(shell go vet 2>&1)"; if [ -n "$$GOVET_OUT" ]; \
		then printf '\n> [WARN] Fix suspicious code from "go vet":\n'; \
		printf "$$GOVET_OUT\n\n"; else printf " ...OK!\n"; fi

## Reformats code according to "gofmt".
fmt:
	@echo Formatting:
	@GOFMT_OUT=$(shell gofmt -l -s -w $(SRC_FILES)); if [ -n "$$GOFMT_OUT" ]; \
		then for FILE in "$$GOFMT_OUT"; do printf "> $$FILE\n"; done; \
		else printf "> ...all files formmatted correctly!\n"; fi

## Testing commands:
GOTEST = go test ./... --coverprofile=$(COVER_OUT) --timeout=$(TEST_TIMEOUT)
test:
	@echo "Testing:"
	@$(GOTEST)
test-v:
	@echo "Testing (verbose):"
	@$(GOTEST) -v
test-race:
	@echo "Testing (race):"
	@$(GOTEST) -race
bench:
	@echo "Benchmarking..."
	@go test -run=^$ -bench=. -benchmem
