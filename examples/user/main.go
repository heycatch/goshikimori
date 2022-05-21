package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "name",
    "key",
  )
}

func main() {
  c := conf()
  u := c.SearchUser("incarnati0n")
  fmt.Println(u.Id, u.Sex, u.Last_online, u.Name)
}
