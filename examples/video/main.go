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

func main() {
  c := conf()
  w, err := c.FastIdAnime("initial d first stage").SearchAnimeVideos()
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
