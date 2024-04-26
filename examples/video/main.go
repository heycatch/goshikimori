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
  fast, status, err := c.FastIdAnime("initial d first stage")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  w, err := fast.SearchAnimeVideos()
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
