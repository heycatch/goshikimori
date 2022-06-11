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
  p, err := c.SearchPublishers()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(p) == 0 {
    fmt.Println("not found publishers")
    return
  }
  for _, v := range p {
    fmt.Println(v.Id, v.Name)
  }
}