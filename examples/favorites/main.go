package main

import (
  "fmt"
  g "github.com/heycatch/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()
  fast, status, err := c.FastIdAnime("Naruto")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    // add/remove favorites anime
    f, err := fast.FavoritesCreate("Anime", "")
    //f, err := fast.FavoritesDelete("Anime")

    // add/remove favorites manga
    //f, err := fast.FavoritesCreate("Manga", "")
    //f, err := fast.FavoritesDelete("Manga")
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(f.Success, f.Notice)
  } else {
    fmt.Println(status)
  }
}
