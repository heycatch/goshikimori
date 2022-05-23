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
  e := &g.ExtraCensored{Censored: "false"}
  ca := c.SearchCalendar(e)
  for _, v := range ca {
    fmt.Println(
      v.Next_episode, v.Next_episode_at, v.Duration,
      v.Anime.Id, v.Anime.Name, v.Anime.Score,
    )
  }
}
