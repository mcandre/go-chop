all: test

test:
	chop < example.txt
	chomp < example.txt

gofmt:
	gofmt -s -w .

lint: gofmt
