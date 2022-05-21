package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "APPLICATION_KEY",
  )
}

func main() {
  c := conf()
  f := c.FastIdAnime("initial d first stage")
  w := c.SearchAnimeVideos(f)
  for _, v := range w {
    fmt.Println(v.Id, v.Url, v.Image_url, v.Player_url, v.Name, v.Kind, v.Hosting)
  }
}
