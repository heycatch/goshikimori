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
  fast, status, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
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
  } else {
    fmt.Println(status)
  }
}
