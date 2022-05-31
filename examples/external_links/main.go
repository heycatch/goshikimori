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

  fa, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fa == 0 {
    fmt.Println("anime not found")
    return
  }
  fra, err := c.SearchAnimeExternalLinks(fa)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(fra) == 0 {
    fmt.Println("external links not found")
    return
  }
  for _, v := range fra {
    fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
  }
  fmt.Println()
  fm, err := c.FastIdManga("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fm == 0 {
    fmt.Println("manga not found")
    return
  }
  frm, err := c.SearchMangaExternalLinks(fm)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(frm) == 0 {
    fmt.Println("external links not found")
    return
  }
  for _, v := range frm {
    fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
  }
}
