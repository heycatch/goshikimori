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

  fra, err := c.FastIdAnime("initial d").SearchAnimeExternalLinks()
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

  frm, err := c.FastIdManga("initial d").SearchMangaExternalLinks()
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
