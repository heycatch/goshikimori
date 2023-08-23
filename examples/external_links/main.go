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
  fast, status, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    fra, err := fast.SearchAnimeExternalLinks()
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

    frm, err := fast.SearchMangaExternalLinks()
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
  } else {
    fmt.Println(status)
  }
}
