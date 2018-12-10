.DEFAULT_GOAL := build

BUILD_VERSION?=snapshot
SOURCE_FILES?=$$(go list ./... | grep -v /vendor/)
TEST_PATTERN?=.
TEST_OPTIONS?=-race

BINARY=bin
BUILD_TIME=`date +%FT%T%z`
COMMIT=`git log --pretty=format:'%h' -n 1`

# Choose to install ngind with or without SputnikVM.
WITH_SVM?=1

# Provide default value of GOPATH, if it's not set in environment
export GOPATH?=${HOME}/go

LDFLAGS=-ldflags "-s -w -X main.Version="`git describe --tags`

build: build_ngind ## Build a local snapshot binary versions of ngind
	@ls -ld $(BINARY)/*

cmd/ngind: chainconfig ## Build a local snapshot binary version of ngind. Use WITH_SVM=1 to enable building with SputnikVM (default: WITH_SVM=1)
ifeq (${WITH_SVM}, 1)
	./scripts/build_sputnikvm.sh
else
	mkdir -p ./${BINARY}
	CGO_CFLAGS_ALLOW='.*' go build ${LDFLAGS} -o ${BINARY}/ngind -tags="netgo" ./cmd/ngind
endif
	@echo "Done building ngin."
	@echo "Run \"$(BINARY)/ngind\" to launch ngind."

cmd/abigen: ## Build a local snapshot binary version of abigen.
	mkdir -p ./${BINARY} && go build ${LDFLAGS} -o ${BINARY}/abigen ./cmd/abigen
	@echo "Done building abigen."
	@echo "Run \"$(BINARY)/abigen\" to launch abigen."

cmd/bootnode: ## Build a local snapshot of bootnode.
	mkdir -p ./${BINARY} && go build ${LDFLAGS} -o ${BINARY}/bootnode ./cmd/bootnode
	@echo "Done building bootnode."
	@echo "Run \"$(BINARY)/bootnode\" to launch bootnode."

cmd/disasm: ## Build a local snapshot of disasm.
	mkdir -p ./${BINARY} && go build ${LDFLAGS} -o ${BINARY}/disasm ./cmd/disasm
	@echo "Done building disasm."
	@echo "Run \"$(BINARY)/disasm\" to launch disasm."

cmd/evm: ## Build a local snapshot of evm.
	mkdir -p ./${BINARY} && CGO_CFLAGS_ALLOW='.*' go build ${LDFLAGS} -o ${BINARY}/evm ./cmd/evm
	@echo "Done building evm."
	@echo "Run \"$(BINARY)/evm\" to launch evm."

cmd/rlpdump: ## Build a local snapshot of rlpdump.
	mkdir -p ./${BINARY} && go build ${LDFLAGS} -o ${BINARY}/rlpdump ./cmd/rlpdump
	@echo "Done building rlpdump."
	@echo "Run \"$(BINARY)/rlpdump\" to launch rlpdump."

build_ngind: ## Build ngind to ./bin. Use WITH_SVM=0 to disable building with SputnikVM (default: WITH_SVM=0)
	$(info Building bin/ngind)
ifeq (${WITH_SVM}, 1)
	chmod +x scripts/build_sputnikvm.sh && ./scripts/build_sputnikvm.sh
else
	CGO_CFLAGS_ALLOW='.*' go build ${LDFLAGS} -tags="netgo" ./cmd/ngind ; fi
endif

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' -not -wholename './_vendor*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

chainconfig: core/assets/assets.go ## Rebuild assets if source config files changed.

core/assets/assets.go: ${GOPATH}/bin/resources core/config/*.json # core/config/*.csv
	${GOPATH}/bin/resources -fmt -declare -var=DEFAULTS -package=assets -output=core/assets/assets.go core/config/*.json core/config/*.csv

${GOPATH}/bin/resources:
	go get -u github.com/omeid/go-resources/cmd/resources

clean: ## Remove local snapshot binary directory
	if [ -d ${BINARY} ] ; then rm -rf ${BINARY} ; fi
	if [ -d "sputnikvm-ffi" ] ; then rm -rf "sputnikvm-ffi" ; fi
	go clean -i ./...

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt build cmd/ngind cmd/abigen cmd/bootnode cmd/disasm cmd/evm cmd/rlpdump build_ngind clean help static
