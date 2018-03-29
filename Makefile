PACKAGES=$(shell go list ./... | grep -v /vendor/)
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))

# target
VERSION := `git rev-parse --short HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=github.com/axelspringer/templeton/version.Version=$(VERSION)"

help:
	@echo 'Available commands:'
	@echo
	@echo 'Usage:'
	@echo '    make deps     		Install go deps.'
	@echo '    make build    		Compile the project.'
	@echo '    make build/docker	Restore all dependencies.'
	@echo '    make restore  		Restore all dependencies.'
	@echo '    make clean    		Clean the directory tree.'
	@echo

test: ## run tests, except integration tests
	@go test ${RACE} ${PACKAGES}

deps:
	go get -u github.com/tcnksm/ghr
	go get -u github.com/mitchellh/gox
	go get -u github.com/golang/dep/cmd/dep

build:
	@echo "Compiling..."
	@mkdir -p ./bin
	@gox $(LDFLAGS) -output "bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -os="linux" -os="darwin" -arch="386" -arch="amd64" ./
	@go build -i -o ./bin/templeton
	@echo "All done! The binaries is in ./bin let's have fun!"

build/docker: build
	@docker build -t templeton:latest .

vet: ## run go vet
	@test -z "$$(go vet ${PACKAGES} 2>&1 | grep -v '*composite literal uses unkeyed fields|exit status 0)' | tee /dev/stderr)"

ci: vet test

restore:
	@dep ensure
