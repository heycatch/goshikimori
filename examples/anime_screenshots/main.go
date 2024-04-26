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
  fast, status, err := c.FastIdAnime("initial d")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  s, err := fast.SearchAnimeScreenshots()
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
