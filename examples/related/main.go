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

  a, err := c.FastIdAnime("initial d second stage").SearchRelatedAnime()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("anime not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Relation, v.Relation_Russian, v.Anime.Score)
  }

  m, err := c.FastIdManga("vampire knight").SearchRelatedManga()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(m) == 0 {
    fmt.Println("Manga not found")
    return
  }
  for _, v := range m {
    fmt.Println(
      v.Relation, v.Relation_Russian,
      v.Manga.Score, v.Manga.Status,
    )
  }
}
