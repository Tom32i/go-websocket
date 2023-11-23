.SILENT:
.PHONY: build

## Colors
COLOR_RESET   = \033[0m
COLOR_INFO    = \033[32m
COLOR_COMMENT = \033[33m

## Help
help:
	printf "${COLOR_COMMENT}Usage:${COLOR_RESET}\n"
	printf " make [target]\n\n"
	printf "${COLOR_COMMENT}Available targets:${COLOR_RESET}\n"
	awk '/^[a-zA-Z\-\_0-9\.@]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf " ${COLOR_INFO}%-16s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## Install dependencies
install:
	go get
	npm install

## Start watcher
watch:
	npx webpack --watch --mode=development

build-client:
	npx webpack --mode=production

## Lint
#lint:
#	npx eslint src/* --ext .js,.json --fix

build-server@dev:
	go build -o ./bin/server main.go

build-server@staging: export GOOS=linux
build-server@staging: export GOARCH=arm64
build-server@staging:
	go build -o ./bin/server main.go

build-server@production: export GOOS=linux
build-server@production: export GOARCH=amd64
build-server@production:
	go build -o ./bin/server main.go

## Build
build: build-server@dev build-client

## Build and run server
run: build
	./bin/server

## Start server
start: export GODEBUG=gctrace=1
start:
	go run main.go

##########
# Deploy #
##########

deploy@staging: build-client build-server@staging
	rsync -arzv --delete public/* tom32i@deployer.vm:/home/tom32i/go/public
	rsync -arzv bin/server tom32i@deployer.vm:/home/tom32i/go/

## Build and deploy to production
deploy@production: build-client build-server@production
	rsync -arzv --delete public/* tom32i@tom32i.fr:/home/tom32i/go/public
	rsync -arzv bin/server tom32i@tom32i.fr:/home/tom32i/go
