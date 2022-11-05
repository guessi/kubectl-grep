.PHONY: staticcheck dependency clean build release all

PKGS       := $(shell go list ./...)
REPO       := github.com/guessi/kubectl-grep
BUILDTIME  := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GITVERSION := $(shell git describe --tags --abbrev=8)
GOVERSION  := $(shell go version | cut -d' ' -f3)
LDFLAGS    := -s -w -X "$(REPO)/cmd.gitVersion=$(GITVERSION)" -X "$(REPO)/cmd.goVersion=$(GOVERSION)" -X "$(REPO)/cmd.buildTime=$(BUILDTIME)"

default: build

staticcheck:
	@echo "Golang Staticcheck..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@for i in $(PKGS); do echo $${i}; staticcheck $${i}; done

test:
	go version
	go fmt ./...
	go vet ./...
	# go test -v ./...

dependency:
	go mod download

build-linux-x86_64:
	@echo "Creating Build for Linux (x86_64)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(GITVERSION)/Linux-x86_64/kubectl-grep
	@cp ./LICENSE ./releases/$(GITVERSION)/Linux-x86_64/LICENSE
	@tar zcf ./releases/$(GITVERSION)/kubectl-grep-Linux-x86_64.tar.gz -C releases/$(GITVERSION)/Linux-x86_64 kubectl-grep LICENSE

build-linux-arm64:
	@echo "Creating Build for Linux (arm64)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(GITVERSION)/Linux-arm64/kubectl-grep
	@cp ./LICENSE ./releases/$(GITVERSION)/Linux-arm64/LICENSE
	@tar zcf ./releases/$(GITVERSION)/kubectl-grep-Linux-arm64.tar.gz -C releases/$(GITVERSION)/Linux-arm64 kubectl-grep LICENSE

build-darwin-x86_64:
	@echo "Creating Build for macOS (x86_64)..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(GITVERSION)/Darwin-x86_64/kubectl-grep
	@cp ./LICENSE ./releases/$(GITVERSION)/Darwin-x86_64/LICENSE
	@tar zcf ./releases/$(GITVERSION)/kubectl-grep-Darwin-x86_64.tar.gz -C releases/$(GITVERSION)/Darwin-x86_64 kubectl-grep LICENSE

build-darwin-arm64:
	@echo "Creating Build for macOS (arm64)..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(GITVERSION)/Darwin-arm64/kubectl-grep
	@cp ./LICENSE ./releases/$(GITVERSION)/Darwin-arm64/LICENSE
	@tar zcf ./releases/$(GITVERSION)/kubectl-grep-Darwin-arm64.tar.gz -C releases/$(GITVERSION)/Darwin-arm64 kubectl-grep LICENSE

build-windows-x86_64:
	@echo "Creating Build for Windows (x86_64)..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./releases/$(GITVERSION)/Windows-x86_64/kubectl-grep.exe
	@cp ./LICENSE ./releases/$(GITVERSION)/Windows-x86_64/LICENSE.txt
	@tar zcf ./releases/$(GITVERSION)/kubectl-grep-Windows-x86_64.tar.gz -C releases/$(GITVERSION)/Windows-x86_64 kubectl-grep.exe LICENSE.txt

build-linux: build-linux-x86_64 build-linux-arm64
build-darwin: build-darwin-x86_64 build-darwin-arm64
build-windows: build-windows-x86_64

build: build-linux build-darwin build-windows

clean:
	@echo "Cleanup Releases..."
	rm -rvf ./releases/*

release:
	@echo "Creating Releases..."
	@curl -LO https://github.com/tcnksm/ghr/releases/download/v0.16.0/ghr_v0.16.0_linux_amd64.tar.gz
	@tar --strip-components=1 -xvf ghr_v0.16.0_linux_amd64.tar.gz ghr_v0.16.0_linux_amd64/ghr
	./ghr -version
	./ghr -replace -recreate -token ${GITHUB_TOKEN} $(GITVERSION) releases/$(GITVERSION)/
	sha1sum releases/$(GITVERSION)/*.tar.gz > releases/$(VERSION)/SHA1SUM

krew-release-bot:
	@echo "Preparing krew-release-bot"
	@curl -LO https://github.com/rajatjindal/krew-release-bot/releases/download/v0.0.43/krew-release-bot_v0.0.43_linux_amd64.tar.gz
	@tar -xvf krew-release-bot_v0.0.43_linux_amd64.tar.gz
	./krew-release-bot action

all: staticcheck dependency clean build
