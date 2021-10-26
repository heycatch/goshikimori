## About
A small library for interacting with shikimori, written in golang.
The library allows you to search the shikimori database.
Work with API occurs only through OAuth2.

## Documentation
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)
* [OAuth2](https://shikimori.one/oauth)

## Install
```
go get github.com/vexilology/goshikimori
```

## Version
```
go version go1.17 linux/amd64
```

## Example 1
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  r, err := c.SearchUser("incarnati0n")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%d - %s - %s\n", r.Id, r.Last_Online, r.Sex)
}
```
## Example 2
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  r, err := c.SearchAnime("Initial D")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    fmt.Println(values.Name, values.Status, values.Score)
  }
}
```

## Example 3
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  r, err := c.SearchManga("Initial D")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    fmt.Println(values.Name, values.Volumes, values.Chapters)
  }
}
```
