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
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    for _, v := range b {
      fmt.Println(v.Id, v.User.Id, v.User.Nickname)
    }
  } else {
    fmt.Println(status)
  }
}
