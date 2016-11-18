VERSION=0.0.1

.PHONY: build-ports clean clean-ports

all: gotest

gotest:
	go test

integration-test:
	chop < example.txt
	chomp < example.txt

govet:
	go list ./... | grep -v vendor | xargs go vet -v

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

lint: govet gofmt goimport

build-ports:
	sh -c "mkdir -p bin/chop-$(VERSION)/linux/amd64   && cd cmd/chop && env GOOS=linux   GOARCH=amd64 go build -o ../../bin/chop-$(VERSION)/linux/amd64/chop"
	sh -c "mkdir -p bin/chop-$(VERSION)/linux/386     && cd cmd/chop && env GOOS=linux   GOARCH=386   go build -o ../../bin/chop-$(VERSION)/linux/386/chop"
	sh -c "mkdir -p bin/chop-$(VERSION)/darwin/amd64  && cd cmd/chop && env GOOS=darwin  GOARCH=amd64 go build -o ../../bin/chop-$(VERSION)/darwin/amd64/chop"
	sh -c "mkdir -p bin/chop-$(VERSION)/darwin/386    && cd cmd/chop && env GOOS=darwin  GOARCH=386   go build -o ../../bin/chop-$(VERSION)/darwin/386/chop"
	sh -c "mkdir -p bin/chop-$(VERSION)/windows/amd64 && cd cmd/chop && env GOOS=windows GOARCH=amd64 go build -o ../../bin/chop-$(VERSION)/windows/amd64/chop.exe"
	sh -c "mkdir -p bin/chop-$(VERSION)/windows/386   && cd cmd/chop && env GOOS=windows GOARCH=386   go build -o ../../bin/chop-$(VERSION)/windows/386/chop.exe"

	sh -c "mkdir -p bin/chop-$(VERSION)/linux/amd64   && cd cmd/chomp && env GOOS=linux   GOARCH=amd64 go build -o ../../bin/chop-$(VERSION)/linux/amd64/chomp"
	sh -c "mkdir -p bin/chop-$(VERSION)/linux/386     && cd cmd/chomp && env GOOS=linux   GOARCH=386   go build -o ../../bin/chop-$(VERSION)/linux/386/chomp"
	sh -c "mkdir -p bin/chop-$(VERSION)/darwin/amd64  && cd cmd/chomp && env GOOS=darwin  GOARCH=amd64 go build -o ../../bin/chop-$(VERSION)/darwin/amd64/chomp"
	sh -c "mkdir -p bin/chop-$(VERSION)/darwin/386    && cd cmd/chomp && env GOOS=darwin  GOARCH=386   go build -o ../../bin/chop-$(VERSION)/darwin/386/chomp"
	sh -c "mkdir -p bin/chop-$(VERSION)/windows/amd64 && cd cmd/chomp && env GOOS=windows GOARCH=amd64 go build -o ../../bin/chop-$(VERSION)/windows/amd64/chomp.exe"
	sh -c "mkdir -p bin/chop-$(VERSION)/windows/386   && cd cmd/chomp && env GOOS=windows GOARCH=386   go build -o ../../bin/chop-$(VERSION)/windows/386/chomp.exe"

	sh -c "cd bin && zip -r chop-$(VERSION).zip chop-$(VERSION)/"

clean: clean-ports

clean-ports:
	rm -rf bin
