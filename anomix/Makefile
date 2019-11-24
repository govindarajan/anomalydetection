override BINARY=anomix
VERSION=V1.0.0
API_ENV?=local
host?="0.0.0.0"
port?="80"

generate:
	(cd pkg/contracts && go generate -v && goimports -w *.go > /dev/null)
	(cd cmd/http/handlers && go generate -v && goimports -w *.go > /dev/null)
	(cd cmd/http/ && go generate -v  && goimports -w *.go > /dev/null )

clean: ## Clean up all the resources
	(cd ./../../contracts && rm *_extractor.go ; rm *_validator.go ;echo "true"> /dev/null)
	(cd handlers ; rm *_handler.go ;echo "true" > /dev/null)
	rm -rf target
	go clean -r

compile: ## Compile the project to generate the binary in the target folder
ifndef VERSION
	$(error VERSION environment variable has to be set)
endif
	(cd cmd/http/ && go build -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=`git rev-parse HEAD` -X main.BuildTime=`date +%FT%T%z`" -o ../../build/${BINARY})

test: ## Run the test cases in random order via ginkgo
	go get
	ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race

package: clean generate compile ## Compile and package the service
	mkdir -p build
	cp cmd/http/config.json.* build
	tar -cvf anomix.tar.gz target/

install: logdir config ## Installs anomix

config:
ifndef API_ENV
	$(error "API_ENV variable not defined, valid value : local,prod,staging")
endif
	ln -sf  ${CURDIR}/cmd/http/config.json.${API_ENV} build/config.json

logdir: ## creates log directory
	mkdir -p /var/log/anomix

run: ## executes the anomix service
	./${BINARY} --host ${host} --port ${port}

debug: ## create debug file and debugs it with gdb
	go build -gcflags "-N -l" -o gdb_sandbox
	gdb gdb_sandbox

help: ## You can always run this command to see whau options are available to you while running the make command
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help compile clean package all clean doc logdir
.DEFAULT_GOAL := help
