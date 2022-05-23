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

func similarAnime() {
  c := conf()
  a := c.FastIdAnime("vampire knight")
  s := c.SearchSimilarAnime(a)
  for _, v := range s {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
}

func similarManga() {
  c := conf()
  a := c.FastIdManga("initial d")
  s := c.SearchSimilarManga(a)
  for _, v := range s {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
}

func main() {
  similarAnime()
  similarManga()
}
