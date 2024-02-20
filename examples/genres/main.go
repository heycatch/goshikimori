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
  g, status, err := c.SearchGenres()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  if len(g) == 0 {
    fmt.Println("not found genres")
    return
  }
  for _, v := range g {
    fmt.Println(v.Id, v.Name, v.Russian, v.Kind, v.Entry_type)
  }
}

