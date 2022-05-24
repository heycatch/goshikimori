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
  f, err := c.FastIdAnime("initial d first stage")
  if err != nil {
    fmt.Println(err)
    return
  }
  if f == 0 {
    fmt.Println("Id not found")
    return
  }
  w, err := c.SearchAnimeVideos(f)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(w) == 0 {
    fmt.Println("Video not found")
    return
  }
  for _, v := range w {
    fmt.Println(v.Id, v.Url, v.Image_url, v.Player_url, v.Name, v.Kind, v.Hosting)
  }
}
