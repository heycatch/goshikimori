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
  user := "incarnati0n"
  c := conf()
  u, err := c.SearchUser(user)
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Printf("Not found %s\n", user)
    return
  }
  fmt.Println(u.Id, u.Sex, u.Last_online, u.Name)
}
