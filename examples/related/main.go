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
  a := c.FastIdAnime("initial d second stage")
  r := c.SearchRelatedAnime(a)
  for _, v := range r {
    fmt.Println(v.Relation, v.Relation_Russian, v.Anime.Score)
  }
}

func relatedManga() {
  c := conf()
  m := c.FastIdManga("vampire knight")
  r := c.SearchRelatedManga(m)
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
