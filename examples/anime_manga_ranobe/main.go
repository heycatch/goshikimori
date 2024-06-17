package main

import (
  "fmt"

  g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func foundAnime() {
  c := config()
  fast, status, err := c.FastIdAnime("initial d first stage")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  res, err := fast.SearchAnime()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(
    res.Id, res.Name, res.Description, res.Released_on, res.Score,
    res.User_rate.Score, res.User_rate.Rewatches, res.User_rate.Text,
  )
}

func foundAnimes() {
  c := config()
  o := &g.Options{
    Page: 1, Limit: 5, Order: "", Kind: "", Status: g.ANIME_STATUS_RELEASED,
    Season: g.SEASON_199x, Rating: "", Duration: "",
    Mylist: g.MY_LIST_COMPLETED, Genre_v2: []int{3}, // SKIP GENRE: Genre_v2: nil,
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
  c := config()
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
  fmt.Println(
    res.Id, res.Name, res.Description, res.Released_on, res.Volumes, res.Chapters,
    res.User_rate.Score, res.User_rate.Rewatches, res.User_rate.Text,
  )
}

func foundMangas() {
  c := config()
  o := &g.Options{
    Page: 1, Limit: 1, Order: "", Kind: "", Status: g.MANGA_STATUS_RELEASED,
    Season: g.SEASON_199x, Score: 8, Censored: false, Mylist: "", Genre_v2: []int{84}, // SKIP GENRE: Genre_v2: nil,
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
  c := config()
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
  fmt.Println(
    res.Id, res.Name, res.Description, res.Released_on, res.Volumes, res.Chapters,
    res.User_rate.Score, res.User_rate.Rewatches, res.User_rate.Text,
  )
}

func foundRanobes() {
  c := config()
  o := &g.Options{
    Page: 1, Limit: 10, Order: g.MANGA_ORDER_POPULARITY, Status: g.MANGA_STATUS_RELEASED,
    Season: "", Mylist: g.MY_LIST_PLANNED, Genre_v2: []int{49}, // SKIP GENRE: Genre_v2: nil,
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
