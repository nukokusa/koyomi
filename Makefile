.PHONY: test lint deps install clean

test: deps lint
	go test -v ./...

lint: deps
	go vet -all -printfuncs=Criticalf,Infof,Warningf,Debugf,Tracef ./...

deps:
	go get golang.org/x/lint/golint

clean:
	go clean
