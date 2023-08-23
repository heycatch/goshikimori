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
  fast_anime, status, err := c.FastIdAnime("vampire knight")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    a, err := fast_anime.SearchSimilarAnime()
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
  } else {
    fmt.Println(status)
  }
  fast_manga, status, err := c.FastIdManga("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    m, err := fast_manga.SearchSimilarManga()
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
  } else {
    fmt.Println(status)
  }
}
