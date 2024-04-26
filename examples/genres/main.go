package main

import (
  "fmt"

  g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := config()
  g, status, err := c.SearchGenres("Anime")
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

