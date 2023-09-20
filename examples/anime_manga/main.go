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
  fast, status, err := c.FastIdAnime("initial d first stage")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  res, err := fast.SearchAnime()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(res.Id, res.Name, res.Description, res.Released_on, res.Score)
}

func foundAnimes() {
  c := conf()
  o := &g.Options{
    Page: "1", Limit: "2", Order: "", Kind: "", Status: "released",
    Season: "199x", Score: "", Rating: "", Duration: "",
    Censored: "", Mylist: "",
  }
  a, status, err := c.SearchAnimes("initial d", o)
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
  fast, status, err := c.FastIdManga("initial d")
  if status != 200 || err != nil {
    fmt.Println(err)
    return
  }
  res, err := fast.SearchManga()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(res.Id, res.Name, res.Description, res.Released_on, res.Volumes, res.Chapters)
}

func foundMangas() {
  c := conf()
  o := &g.Options{
    Page: "1", Limit: "1", Order: "", Kind: "", Status: "released",
    Season: "199x", Score: "8", Censored: "", Mylist: "",
  }
  m, status, err := c.SearchMangas("initial d", o)
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
  foundAnimes()

  foundManga()
  foundMangas()
}
