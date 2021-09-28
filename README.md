## About
A small library for interacting with shikimori, written in golang.
Work with API occurs only through OAuth2.

## Documentation
* API v1 https://shikimori.one/api/doc/1.0
* API v2 https://shikimori.one/api/doc/2.0 
* OAuth2 https://shikimori.one/oauth

## Install
```
go get github.com/vexilology/goshikimori
```

## Version
```
go version go1.17 linux/amd64
GNU Make 4.2.1
```

## Example 1
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori"
  "github.com/vexilology/goshikimori/api"
  "github.com/vexilology/goshikimori/req"
)

func returnConf() *req.Config {
  return &req.Config{
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  }
}

func main() {
  conf := returnConf()

  r, err := g.NewRequest(
    conf.Application, conf.SecretKey, req.Get,
    g.Parameters(api.Users, api.Whoami),
  )
  if err != nil {
    log.Fatal(err)
  }

  // GET /api/users/whoami
  fmt.Printf("%s\n", r)
}

```
## Example 2
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori"
  "github.com/vexilology/goshikimori/api"
  "github.com/vexilology/goshikimori/req"
)

func returnConf() *req.Config {
  return &req.Config{
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  }
}

func main() {
  conf := returnConf()

  r, err := g.NewRequest(
    conf.Application, conf.SecretKey, req.Get,
    g.Parameters(api.Animes, api.Id(1)),
  )
  if err != nil {
    log.Fatal(err)
  }

  // GET /api/animes/1
  fmt.Printf("%s\n", r)
}
```
