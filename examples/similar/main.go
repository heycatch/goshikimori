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
  a, err := c.FastIdAnime("vampire knight")
  if err != nil {
    fmt.Println(err)
    return
  }
  if a == 0 {
    fmt.Println("Anime not found")
    return
  }
  s, err := c.SearchSimilarAnime(a)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range s {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
}

func similarManga() {
  c := conf()
  a, err := c.FastIdManga("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if a == 0 {
    fmt.Println("Manga not found")
    return
  }
  s, err := c.SearchSimilarManga(a)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range s {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
}

func main() {
  similarAnime()
  similarManga()
}
