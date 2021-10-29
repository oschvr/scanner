
GOCMD=go
BINARY_NAME=scanner
VERSION?=0.0.1

.PHONY: build vendor clean

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin/$(VERSION)
	for OS in darwin freebsd linux windows; do for ARCH in 386 amd64 arm64 arm; do GOOS=$$OS GOARCH=386 GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(VERSION)/$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH .; tar -czvf out/bin/$(VERSION)/$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH.tar.gz out/bin/$(VERSION)/$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH; rm out/bin/$(VERSION)/$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH; done; done

clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor
