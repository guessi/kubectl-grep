.PHONY: utilities lint dependency clean build release all

VERSION_MAJOR  := 1
VERSION_MINOR  := 7
VERSION_PATCH  := 2
VERSION_SUFFIX := # -dev

COMMIT  := $(shell git describe --always)
PKGS    := $(shell go list ./...)
REPO    := github.com/guessi/kubectl-grep
VERSION := v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)$(VERSION_SUFFIX)
LDFLAGS := -s -w -X $(REPO)/cmd.version=$(VERSION)

default: build

utilities:
	@echo "Download Utilities..."
	go get golang.org/x/lint/golint
	go get github.com/tcnksm/ghr

lint:
	@echo "Source Code Lint..."
	@for i in $(PKGS); do echo $${i}; golint $${i}; done

test:
	go version
	go fmt ./...
	go vet ./...
	# go test -v ./...

dependency:
	go mod download

build-linux:
	@echo "Creating Build for Linux (x86_64)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(VERSION)/Linux-x86_64/kubectl-grep
	@cp ./LICENSE ./releases/$(VERSION)/Linux-x86_64/LICENSE
	@tar zcf ./releases/$(VERSION)/kubectl-grep-Linux-x86_64.tar.gz -C releases/$(VERSION)/Linux-x86_64 kubectl-grep LICENSE

build-darwin-x86_64:
	@echo "Creating Build for macOS (x86_64)..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(VERSION)/Darwin-x86_64/kubectl-grep
	@cp ./LICENSE ./releases/$(VERSION)/Darwin-x86_64/LICENSE
	@tar zcf ./releases/$(VERSION)/kubectl-grep-Darwin-x86_64.tar.gz -C releases/$(VERSION)/Darwin-x86_64 kubectl-grep LICENSE

build-darwin-arm64:
	@echo "Creating Build for macOS (arm64)..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(VERSION)/Darwin-arm64/kubectl-grep
	@cp ./LICENSE ./releases/$(VERSION)/Darwin-arm64/LICENSE
	@tar zcf ./releases/$(VERSION)/kubectl-grep-Darwin-arm64.tar.gz -C releases/$(VERSION)/Darwin-arm64 kubectl-grep LICENSE

build-windows:
	@echo "Creating Build for Windows (x86_64)..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(VERSION)/Windows-x86_64/kubectl-grep.exe
	@cp ./LICENSE ./releases/$(VERSION)/Windows-x86_64/LICENSE.txt
	@tar zcf ./releases/$(VERSION)/kubectl-grep-Windows-x86_64.tar.gz -C releases/$(VERSION)/Windows-x86_64 kubectl-grep.exe LICENSE.txt

build: build-linux build-darwin-x86_64 build-darwin-arm64 build-windows

clean:
	@echo "Cleanup Releases..."
	rm -rvf ./releases/*

release:
	@echo "Creating Releases..."
	go get github.com/tcnksm/ghr
	ghr --replace --recreate -t ${GITHUB_TOKEN} $(VERSION) releases/$(VERSION)/
	sha1sum releases/$(VERSION)/*.tar.gz > releases/$(VERSION)/SHA1SUM

krew-release-bot:
	@echo "Preparing krew-release-bot"
	@curl -LO https://github.com/rajatjindal/krew-release-bot/releases/download/v0.0.43/krew-release-bot_v0.0.43_linux_amd64.tar.gz
	@tar -xvf krew-release-bot_v0.0.43_linux_amd64.tar.gz
	./krew-release-bot action

all: utilities lint dependency clean build
