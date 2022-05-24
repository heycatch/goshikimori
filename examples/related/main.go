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

func relatedAnime() {
  c := conf()
  a, err := c.FastIdAnime("initial d second stage")
  if err != nil {
    fmt.Println(err)
    return
  }
  if a == 0 {
    fmt.Println("Anime not found")
    return
  }
  r, err := c.SearchRelatedAnime(a)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range r {
    fmt.Println(v.Relation, v.Relation_Russian, v.Anime.Score)
  }
}

func relatedManga() {
  c := conf()
  m, err := c.FastIdManga("vampire knight")
  if err != nil {
    fmt.Println(err)
    return
  }
  if m == 0 {
    fmt.Println("Manga not found")
    return
  }
  r, err := c.SearchRelatedManga(m)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range r {
    fmt.Println(
      v.Relation, v.Relation_Russian,
      v.Manga.Score, v.Manga.Status,
    )
  }
}

func main() {
  relatedAnime()
  relatedManga()
}
