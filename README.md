### EN | [RU](https://github.com/heycatch/goshikimori/blob/master/README_ru.md)

### About
A small library for interacting with shikimori, written in golang.
* Work with API occurs only through `OAuth2`.
* No dependencies on other libraries.
* The GNU [make](https://www.gnu.org/software/make/manual/make.html) utility is used for tests and builds.
* Minimum language version >= `1.18`.

### Install
```bash
go get github.com/heycatch/goshikimori
```

### Examples
* [Click her](https://github.com/heycatch/goshikimori/tree/master/examples)

### Documentation
**godoc** support is also available.
```bash
# Method #1: Use docker.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make docker-build
make docker-start

# Open in browser.
http://localhost:1337/pkg/github.com/heycatch/goshikimori/
```
```bash
# Method #2(Linux): Install godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# Add 'export' to the file /home/$USER/.profile
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Reboot and check that the application is working properly.
godoc --help

# After installation or if 'godoc' is already installed.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make doc

# Open in browser.
http://localhost:1337/pkg/github.com/heycatch/goshikimori/
```

### Shikimori documentation
* [GraphQL](https://shikimori.one/api/doc/graphql)
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)
