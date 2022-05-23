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
  a := c.FastIdAnime("initial d")
  s := c.SearchAnimeScreenshots(a)
  for _, v := range s {
    fmt.Println(v.Original, v.Preview)
  }
}
