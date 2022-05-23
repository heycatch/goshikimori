package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()
  b := c.SearchBans()
  for _, v := range b {
    fmt.Println(v.Id, v.User.Id, v.User.Nickname)
  }
}
