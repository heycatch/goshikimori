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
  fast, status, err := c.FastIdUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    a, err := fast.SearchAchievement()
    if err != nil {
      fmt.Println(err)
      return
    }
    if len(a) == 0 {
      fmt.Println("achievements not found")
      return
    }
    for _, v := range a {
      if v.Neko_id == g.NekoSearch("Hellsing") {
        fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
        fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
      }
    }
  } else {
    fmt.Println(status)
  }
}
