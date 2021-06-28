<div align="center">
  <h1>shikimori.one api for golang</h1>
</div>

## Foreword
Work with API occurs only through OAuth2.

## Documentation
* API v1 https://shikimori.one/api/doc/1.0
* API v2 https://shikimori.one/api/doc/2.0 
* OAuth2 https://shikimori.one/oauth

## Install
go get github.com/vexilology/goshikimori

## Example
``` golang
package main

import (
  "fmt"
  "log"

  "github.com/vexilology/goshikimori"
)

func main() {
  result, err := goshikimori.NewRequest(
    "APP_NAME",
    "ACCESS_TOKEN",
    goshikimori.Parameters(api.Users, api.FoundID("ID_HER"), api.Friends),
  )
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(result))
}
```
