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

  // related anime
  fast_anime, status, err := c.FastIdAnime("initial d second stage")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    a, err := fast_anime.SearchRelatedAnime()
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
  } else {
    fmt.Println(status)
  }

  // related manga
  fast_manga, status, err := c.FastIdManga("vampire knight")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    m, err := fast_manga.SearchRelatedManga()
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
  } else {
    fmt.Println(status)
  }

  // related ranobe
  fast_ranobe, status, err := c.FastIdRanobe("sword art")
  if status != 200 || err != nil {
    fmt.Println(err)
    return
  }
  r, err := fast_ranobe.SearchRelatedRanobe()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(r) == 0 {
    fmt.Println("Ranobe not found")
    return
  }
  for _,v := range r {
    fmt.Println(
      v.Relation, v.Relation_Russian,
      v.Manga.Score, v.Manga.Status,
    )
  }
}
