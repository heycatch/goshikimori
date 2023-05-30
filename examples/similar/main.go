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

  a, err := c.FastIdAnime("vampire knight").SearchSimilarAnime()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("Anime not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Name, v.Id, v.Russian)
  }

  m, err := c.FastIdManga("initial d").SearchSimilarManga()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(m) == 0 {
    fmt.Println("Manga not found")
    return
  }
  for _, v := range m {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
}
