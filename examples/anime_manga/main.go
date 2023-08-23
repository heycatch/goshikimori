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

func foundAnime() {
  c := conf()
  o := &g.Options{
    Page: "1", Limit: "2", Kind: "", Status: "released",
    Season: "199x", Score: "", Rating: "", Duration: "",
    Censored: "", Mylist: "",
  }
  a, status, err := c.SearchAnime("initial d", o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    if len(a) == 0 {
      fmt.Println("Anime not found")
      return
    }
    for _, v := range a {
      fmt.Println(v.Name, v.Released_on, v.Score)
    }
  } else {
    fmt.Println(status)
  }
}

func foundManga() {
  c := conf()
  o := &g.Options{
    Page: "1", Limit: "1", Kind: "", Status: "released",
    Season: "199x", Score: "8", Censored: "", Mylist: "",
  }
  m, status, err := c.SearchManga("initial d", o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    if len(m) == 0 {
      fmt.Println("Manga not found")
      return
    }
    for _, v := range m {
      fmt.Println(v.Name, v.Released_on, v.Score)
    }
  } else {
    fmt.Println(status)
  }
}

func main() {
  foundAnime()
  foundManga()
}
