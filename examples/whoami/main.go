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
  w, status, err := c.WhoAmi()
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 { fmt.Println(w.Nickname, w.Avatar, w.Locale, w.Last_online_at) }
}
