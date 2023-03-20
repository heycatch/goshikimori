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

  a, err := c.RandomAnime()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(a)

  m, err := c.RandomManga()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(m)
}
