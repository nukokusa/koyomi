GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: lint test install

lint: $(GOBIN)/staticcheck
	staticcheck ./...

test: lint
	go test ./...

install:
	go install github.com/nukokusa/koyomi/cmd/koyomi

$(GOBIN)/staticcheck:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
