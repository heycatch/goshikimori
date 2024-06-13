## EN | [RU](https://github.com/heycatch/goshikimori/blob/master/README.md)

### About
A small library for interacting with shikimori, written in golang.
* Work with API occurs only through `OAuth2`.
* No dependencies on other libraries.
* The GNU [make](https://www.gnu.org/software/make/manual/make.html)
utility is used for tests and builds.

### Install
```bash
go get github.com/heycatch/goshikimori
```

### Examples
* [Click her](https://github.com/heycatch/goshikimori/tree/master/examples)

### GODOC / GOPKG documentation
**Godoc** support is also available.

Or you can use the page from the official Go pkg
[website](https://pkg.go.dev/github.com/heycatch/goshikimori).\
**P.S.** documentation is late in updating.
```bash
# Method #1: Use docker.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make docker-build
make docker-start

# Open in browser.
http://localhost:1337/pkg/github.com/heycatch/goshikimori
```
```bash
# Method #2(Linux): Install godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# Add 'export' to the file /home/$USER/.profile and reboot.
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Check that the application is working properly.
godoc --help

# After installation or if 'godoc' is already installed.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make doc

# Open in browser.
http://localhost:1337/pkg/github.com/heycatch/goshikimori
```

### Shikimori documentation
* [GraphQL](https://shikimori.one/api/doc/graphql)
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)

### Feedback
* Private message on the [website](https://shikimori.one/arctica).
* Open [issue](https://github.com/heycatch/goshikimori/issues).
```bash
# Current tasks and problems in the code can be viewed in the terminal using the command.
git grep TODO
git grep FIXME
```
