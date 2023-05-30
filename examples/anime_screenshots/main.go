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
  s, err := c.FastIdAnime("initial d").SearchAnimeScreenshots()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(s) == 0 {
    fmt.Println("Screenshots not found")
    return
  }
  for _, v := range s {
    fmt.Println(v.Original, v.Preview)
  }
}
