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
  // franchise anime
  fast_anime, status, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    fa, err := fast_anime.SearchAnimeFranchise()
    if err != nil {
      fmt.Println(err)
      return
    }
    if len(fa.Nodes) == 0 {
      fmt.Println("anime franchise not found")
      return
    }
    for _, v := range fa.Nodes {
      fmt.Println(v.Id, v.Name, v.Kind)
    }
  } else {
    fmt.Println(status)
  }
  // franchise manga
  fast_manga, status, err := c.FastIdManga("naruto")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    fm, err := fast_manga.SearchMangaFranchise()
    if err != nil {
      fmt.Println(err)
      return
    }
    if len(fm.Nodes) == 0 {
      fmt.Println("manga franchise not found")
      return
    }
    for _, v := range fm.Nodes {
      fmt.Println(v.Id, v.Name, v.Kind)
    }
  } else {
    fmt.Println(status)
  }
}
