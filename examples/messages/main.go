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
  u, err := c.SearchUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
    return
  }

  o := &g.Options{Type: "news", Page: "1", Limit: "1"}
  m, err := c.UserMessages(u.Id, o)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range m {
    fmt.Println(v.Id, v.Kind, v.HTMLBody, v.Created_at)
  }
}
