GO_FILES = $(wildcard *.go)

.PHONY: run_products
## run_products: run the products micro-services
run_products:
	npx bazelisk run //services/products:products

.PHONY: run_reviews
## run_reviews: run the reviews micro-services
run_reviews:
	npx bazelisk run //services/reviews:reviews

.PHONY: run_gateway
## run_gateway: run the GraphQL gateway
run_gateway:
	npx graphql-mesh serve

.PHONY: update
## update: update all BUILD files
update:
	npx bazelisk run //:gazelle -- update

.PHONY: deps
## deps: fetch dependencies and update the DEPENDENCIES.bzl file
deps:
	npx bazelisk run //:gazelle -- update-repos -from_file=go.mod -prune=true -to_macro=DEPENDENCIES.bzl%go_repositories

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
