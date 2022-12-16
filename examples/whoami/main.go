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

func foundAnime() {
  c := conf()
  w, err := c.WhoAmi()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(w.Nickname, w.Avatar, w.Locale, w.Last_online_at)
}
