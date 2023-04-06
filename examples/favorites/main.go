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

  fia, err := c.FastIdAnime("Naruto")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fia == 0 {
    fmt.Println("anime not found.")
    return
  }

  ff, err := c.FavoritesCreate("Anime", fia, "")
  //ff, err := c.FavoritesDelete("Anime", fia)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(ff.Success, ff.Notice)
}