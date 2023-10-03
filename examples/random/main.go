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

  a, status, err := c.RandomAnimes(5)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range a {
    fmt.Println(v.Id, v.Name, v.Score, v.Status)
  }

  m, status, err := c.RandomMangas(5)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range m {
    fmt.Println(v.Id, v.Name, v.Score, v.Chapters, v.Volumes)
  }
}
