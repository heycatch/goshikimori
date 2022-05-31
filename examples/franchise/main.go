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

  fa, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fa == 0 {
    fmt.Println("anime not found")
    return
  }
  fr, err := c.SearchAnimeFranchise(fa)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(fr.Nodes) == 0 {
    fmt.Println("franchise not found")
    return
  }
  for _, v := range fr.Nodes {
    fmt.Println(v.Id, v.Name, v.Kind)
  }
  fmt.Println()
  fm, err := c.FastIdManga("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fm == 0 {
    fmt.Println("manga not found")
    return
  }
  frr, err := c.SearchMangaFranchise(fm)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(frr.Nodes) == 0 {
    fmt.Println("franchise not found")
    return
  }
  for _, v := range frr.Nodes {
    fmt.Println(v.Id, v.Name, v.Kind)
  }
}
