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

  // add/remove favorites anime
  f, err := c.FastIdAnime("Naruto").FavoritesCreate("Anime", "")
  //f, err := c.FastIdAnime("Naruto").FavoritesDelete("Anime")

  // add/remove favorites manga
  //f, err := c.FastIdManga("Naruto").FavoritesCreate("Manga", "")
  //f, err := c.FastIdManga("Naruto").FavoritesDelete("Manga")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(f.Success, f.Notice)
}