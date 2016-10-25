all: test

test:
	chop < example.txt
	chomp < example.txt

govet:
	go vet -v

gofmt:
	gofmt -s -w .

lint: govet gofmt
