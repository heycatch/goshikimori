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

  fa, err := c.FastIdAnime("initial d").SearchAnimeFranchise()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(fa.Nodes) == 0 {
    fmt.Println("anime franchise not found")
    return
  }
  for _, v := range fa.Nodes {
    fmt.Println(v.Id, v.Name, v.Kind)
  }

  fmt.Println()

  fm, err := c.FastIdManga("naruto").SearchMangaFranchise()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(fm.Nodes) == 0 {
    fmt.Println("manga franchise not found")
    return
  }
  for _, v := range fm.Nodes {
    fmt.Println(v.Id, v.Name, v.Kind)
  }
}
