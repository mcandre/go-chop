# Windows
ifeq ($(OS),Windows_NT)
	EXTENSION=.exe
	RM=cmd /c rmdir /s /q
else
	EXTENSION=
	RM=rm -rf
endif

CHOP=bin/chop$(EXTENSION)
CHOMP=bin/chomp$(EXTENSION)

all: test

test: $(CHOP) $(CHOMP)
	$(CHOP) < example.txt
	$(CHOMP) < example.txt

$(CHOP): cli-chop.go
	go build -o $(CHOP) cli-chop.go

$(CHOMP): cli-chomp.go
	go build -o $(CHOMP) cli-chomp.go

clean:
	-$(RM) bin
