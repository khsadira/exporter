GO		:=	GO111MODULE=on go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)

#FILE	=

DOCKER_IMAGE_NAME       ?= exporter

all: format build

test:
	@echo ">> running tests"
	@go test -short $(pkgs)

format:
	@echo ">> formatting code"
#	@$(FMT) -w $(FILE)

module:
	@echo ">> creating module"
	@$(GO) mod init github.com/khsadira/exporter

build: 
	@echo ">> building binaries"
	@$(GO) build .

docker: all
	@echo ">> building docker image"
	@docker build -t $(DOCKER_IMAGE_NAME) .

fclean:
	@rm -rf exporter go.sum go.mod

re: fclean module all test

.PHONY: all format build test
