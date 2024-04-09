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
  o := &g.Options{Censored: false}
  ca, status, err := c.SearchCalendar(o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    for _, v := range ca {
      fmt.Println(
        v.Next_episode, v.Next_episode_at, v.Duration,
        v.Anime.Id, v.Anime.Name, v.Anime.Score,
      )
    }
  } else {
    fmt.Println(status)
  }
}
