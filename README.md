A small library for interacting with shikimori, written in golang. \
The library allows you to search the shikimori database. \
Work with API occurs only through OAuth2. \
No dependencies on other libraries.

### Install
```bash
go get github.com/vexilology/goshikimori
```

### Examples
* [Click her](https://github.com/vexilology/goshikimori/tree/main/examples)

### Documentation
A large number of functions are now available. \
For convenience, it is better to use **godoc**.
```bash
# Debian/Ubuntu etc.
sudo apt update && sudo apt install golang-doc golang-go.tools -y
# Or use go install.
go install -v golang.org/x/tools/cmd/godoc@latest
```
```bash
# After install.
git clone git@github.com:vexilology/goshikimori.git && cd goshikimori
godoc -http=:1337
```
```bash
# Open in browser.
http://localhost:1337/pkg/github.com/vexilology/goshikimori/
```

### Shikimori documentation
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)