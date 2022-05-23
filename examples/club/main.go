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
  e := &g.ExtraLimit{Limit: "2"}
  a := c.SearchClub("milf", e)
  for _, v := range a {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }
}
