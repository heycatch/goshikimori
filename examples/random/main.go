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

  r, status, err := c.RandomRanobes(5)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range r {
    fmt.Println(v.Id, v.Name, v.Score, v.Chapters, v.Volumes)
  }
}
