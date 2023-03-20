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

  a, err := c.FastIdAnime("vampire knight")
  if err != nil {
    fmt.Println(err)
    return
  }
  if a == 0 {
    fmt.Println("Anime not found")
    return
  }
  fmt.Println(a)

  m, err := c.FastIdManga("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if m == 0 {
    fmt.Println("Manga not found")
    return
  }
  fmt.Println(m)

  cl, err := c.FastIdClub("shikimori api")
  if err != nil {
    fmt.Println(err)
    return
  }
  if cl == 0 {
    fmt.Println("Club not found")
    return
  }
  fmt.Println(cl)
}
