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
    Page: 1, Limit: 5, Order: "", Kind: "", Status: "released",
    Season: "199x", Rating: "", Duration: "",
    Mylist: "", Genre_v2: []int{3}, // SKIP GENRE: Genre_v2: nil,
  }
  a, status, err := c.SearchAnimes("initial d", o)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  if len(a) == 0 {
    fmt.Println("Anime not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Name, v.Released_on, v.Score)
  }
}

func foundManga() {
  c := conf()
  fast, status, err := c.FastIdManga("initial d")
  if status != 200 || err != nil {
    fmt.Println(status, err)
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
    Page: 1, Limit: 1, Order: "", Kind: "", Status: "released",
    Season: "199x", Score: 8, Censored: false, Mylist: "", Genre_v2: []int{84}, // SKIP GENRE: Genre_v2: nil,
  }
  m, status, err := c.SearchMangas("initial d", o)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  if len(m) == 0 {
    fmt.Println("Manga not found")
    return
  }
  for _, v := range m {
    fmt.Println(v.Name, v.Released_on, v.Score, v.Chapters, v.Volumes)
  }
}

func foundRanobe() {
  c := conf()
  fast, status, err := c.FastIdRanobe("sword art")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  res, err := fast.SearchRanobe()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(res.Id, res.Name, res.Description, res.Released_on, res.Volumes, res.Chapters)
}

func foundRanobes() {
  c := conf()
  o := &g.Options{
    Page: 1, Limit: 10, Order: "", Status: "released",
    Season: "", Mylist: "", Genre_v2: []int{49}, // SKIP GENRE: Genre_v2: nil,
  }
  r, status, err := c.SearchRanobes("angel", o)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range r {
    fmt.Println(v.Name, v.Score, v.Released_on, v.Volumes, v.Chapters)
  }
}

func main() {
  foundAnime()
  foundAnimes()

  foundManga()
  foundMangas()

  foundRanobe()
  foundRanobes()
}
