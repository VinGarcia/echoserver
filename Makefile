
GOBIN=$(shell go env GOPATH)/bin

-include config.env
export

run:
	go run main.go

VERSION=0.0.7
release:
	- docker buildx create --name echoserver-builder
	docker buildx use echoserver-builder
	- docker buildx build --push --platform=linux/arm64,linux/amd64 --tag vingarcia/echoserver:$(VERSION) .
	- docker buildx rm echoserver-builder

test: setup
	$(GOBIN)/richgo test $(path) $(args)

lint: setup
	@$(GOBIN)/staticcheck $(path) $(args)
	@go vet $(path) $(args)
	@echo "StaticCheck & Go Vet found no problems on your code!"

setup: $(GOBIN)/richgo $(GOBIN)/staticcheck

$(GOBIN)/richgo:
	go get github.com/kyoh86/richgo@latest

$(GOBIN)/staticcheck:
	go install 'honnef.co/go/tools/cmd/staticcheck@latest'
