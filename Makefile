GO_FILES = $(wildcard *.go)

.PHONY: update
## update: update all BUILD files
update:
	bazel run //:gazelle -- update

.PHONY: deps
## deps: fetch dependencies and update the DEPENDENCIES.bzl file
deps:
	bazel run //:gazelle -- update-repos -from_file=go.mod -prune=true -to_macro=DEPENDENCIES.bzl%go_repositories

.PHONY: format
## format: format all files
format:
	go fmt **/*.go
	goimports -w $(GO_FILES)

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
