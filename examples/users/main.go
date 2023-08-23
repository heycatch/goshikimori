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
  o := &g.Options{Page: "1", Limit: "10"}
  u, status, err := c.SearchUsers("angel", o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    if len(u) == 0 {
      fmt.Println("users not found")
      return
    }
    for _, v := range u {
      fmt.Println(v.Id, v.Nickname, v.Last_online_at)
    }
  } else {
    fmt.Println(status)
  }
}
