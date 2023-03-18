
GOBIN=$(shell go env GOPATH)/bin

run:
	go run main.go

release:
	docker build -t vingarcia/echoserver:latest .
	docker push vingarcia/echoserver:latest

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
