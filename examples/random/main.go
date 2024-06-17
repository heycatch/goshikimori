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

func main() {
  c := config()

  a, status, err := c.RandomAnimes(&g.Options{
    Limit: 5, Score: 5, Kind: g.ANIME_KIND_TV,
    Status: g.ANIME_STATUS_RELEASED,
    Season: "", Rating: "", Duration: "", Mylist: "",
    Censored: false, Genre_v2: []int{10, 539},
  })
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range a {
    fmt.Println(
      v.Id, v.Name, v.Score,
      v.Status, v.Released_on, v.Url,
    )
  }

  m, status, err := c.RandomMangas(&g.Options{
    Limit: 5, Score: 8, Kind: g.MANGA_KIND_MANGA,
    Status: g.MANGA_STATUS_RELEASED,
    Mylist: "", Season: "", Censored: false, Genre_v2: nil,
  })
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range m {
    fmt.Println(
      v.Id, v.Name, v.Score, v.Chapters,
      v.Volumes, v.Released_on, v.Url,
    )
  }

  r, status, err := c.RandomRanobes(&g.Options{
    Limit: 5, Score: 8, Status: g.MANGA_STATUS_RELEASED, Mylist: "",
    Season: "", Censored: false, Genre_v2: nil,
  })
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range r {
    fmt.Println(
      v.Id, v.Name, v.Score, v.Chapters,
      v.Volumes, v.Released_on, v.Url,
    )
  }
}
