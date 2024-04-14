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
  o := &g.Options{}
  ca, status, err := c.SearchCalendar(o)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range ca {
    fmt.Println(
      v.Next_episode, v.Next_episode_at, v.Duration,
      v.Anime.Id, v.Anime.Name, v.Anime.Score,
    )
  }
}
