package main

import (
  "fmt"
  g "github.com/heycatch/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()
  b, status, err := c.SearchBans()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range b {
    fmt.Println(v.Id, v.User.Id, v.User.Nickname)
  }
}
