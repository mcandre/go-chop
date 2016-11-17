# go-chop - a chop + chomp library and CLI written in Go

# EXAMPLES

```
$ chop < example.txt
Go Go Go
Go Gopher

$ chomp < example.txt
Go Go Go,
Go Gopher!
```

# REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [Git](https://git-scm.com)
* [Make](https://www.gnu.org/software/make/)
* [Bash](https://www.gnu.org/software/bash/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/go-chop/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/go-chop.git $GOPATH/src/github.com/mcandre/go-chop
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-chop/cmd/chop && go install"
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-chop/cmd/chomp && go install"
```

# TEST REMOTELY

```
$ go test github.com/mcandre/go-chop
```

# TEST LOCALLY

```
$ go test
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
