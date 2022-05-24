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
  a, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if a == 0 {
    fmt.Println("Anime not found")
    return
  }
  s, err := c.SearchAnimeScreenshots(a)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(s) == 0 {
    fmt.Println("Screenshots not found")
    return
  }
  for _, v := range s {
    fmt.Println(v.Original, v.Preview)
  }
}
