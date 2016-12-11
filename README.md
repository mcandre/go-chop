# go-chop - a chop + chomp library and CLI written in Go

# EXAMPLES

```
$ cat example.txt
Go Go Go,
Go Gopher!

$ chop < example.txt
Go Go Go
Go Gopher

$ chomp < example.txt
Go Go Go,
Go Gopher!

$ chop -v
0.0.1

$ chomp -v
0.0.1
```

# DOWNLOADS

https://github.com/mcandre/go-chop/releases

# DEVELOPMENT REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [editorconfig-tools](https://www.npmjs.com/package/editorconfig-tools)
* [bashate](https://github.com/openstack-dev/bashate)
* [shlint](https://rubygems.org/gems/shlint)
* [shellcheck](http://hackage.haskell.org/package/ShellCheck)


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

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
