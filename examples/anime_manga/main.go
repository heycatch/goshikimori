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

func foundAnime() {
  c := conf()
  e := &g.Extra{
    Limit: "2", Kind: "", Status: "released",
    Season: "199x", Score: "", Rating: "",
  }
  a := c.SearchAnime("initial d", e)
  for _, v := range a {
    fmt.Println(v.Name, v.Released_on, v.Score)
  }
}

func foundManga() {
  c := conf()
  e := &g.Extra{
    Limit: "1", Kind: "", Status: "released",
    Season: "199x", Score: "8",
  }
  m := c.SearchManga("initial d", e)
  for _, v := range m {
    fmt.Println(v.Name, v.Released_on, v.Score)
  }
}

func main() {
  foundAnime()
  foundManga()
}
