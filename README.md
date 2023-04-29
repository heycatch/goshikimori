### About
A small library for interacting with shikimori, written in golang.
* Work with API occurs only through `OAuth2`.
* No dependencies on other libraries.

### Install
```bash
go get github.com/vexilology/goshikimori
```

### Examples
* [Click her](https://github.com/vexilology/goshikimori/tree/main/examples)

### Documentation
A large number of functions are now available. For convenience, it is better to use **godoc**.
```bash
# Install godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# GOPATH export.
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Check that the application is working properly.
godoc --help
```
```bash
# After install.
git clone git@github.com:vexilology/goshikimori.git && cd goshikimori

# Now use godoc.
make doc
# Or use docker.
make docker-build
make docker-start

# Open in browser.
http://localhost:1337/pkg/github.com/vexilology/goshikimori/
http://localhost:1337/pkg/github.com/vexilology/goshikimori/api/
```

### Shikimori documentation
* [API v1](https://shikimori.me/api/doc/1.0)
* [API v2](https://shikimori.me/api/doc/2.0)
* [OAuth2](https://shikimori.me/oauth)
