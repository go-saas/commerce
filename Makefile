GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BUF_VERSION=v1.14.0

.PHONY: init
# init env
init:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-saas/kit/cmd/protoc-gen-go-grpc-proxy@ab0c474680a48b0f3673be8d75b1be7260b6bd68
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/go-saas/kit/cmd/protoc-gen-go-errors-i18n/v2@ab0c474680a48b0f3673be8d75b1be7260b6bd68
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.7
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@$(BUF_VERSION)
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@$(BUF_VERSION)


.PHONY: api
# generate api proto
api:
	buf generate --path ./cart --template ./cart/buf.gen.yaml
	buf generate --path ./order --template ./order/buf.gen.yaml
	buf generate --path ./payment --template ./payment/buf.gen.yaml
	buf generate --path ./product --template ./product/buf.gen.yaml
	buf generate --path ./ticketing --template ./ticketing/buf.gen.yaml
	buf generate

.PHONY: build
# build
build:
	cd cart && make build
	cd order && make build
	cd payment && make build
	cd product && make build
	cd ticketing make build
	cd sys && make build
	cd user && make build

.PHONY: all
# generate all
all:
	make init;
	make api;
	make config;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
