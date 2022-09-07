.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test
## test: runs tests
test:
	@go test -mod=vendor ./app/... -coverprofile cover.out

.PHONY: build
## build: builds application
build:
	@cd app && go build -v -mod=vendor

.PHONY: run
## run: build and runs app locally
run:
	@cd app && go build -v -mod=vendor
	@cd app && ./app --dbconn=mongodb://localhost:27017 --dbname=quake3 --path=../stats