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
  e := &g.ExtraLimit{Page: "1", Limit: "10"}
  u, err := c.SearchUsers("angel", e)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(u) == 0 {
    fmt.Println("users not found")
    return
  }
  for _, v := range u {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }
}
