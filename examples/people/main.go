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

  e := &g.ExtraPeople{Kind: "seyu"}
  sp, err := c.SearchPeople("Aya Hirano", e)
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, v := range sp {
    fmt.Println(v)
  }

  fp, err := c.FastIdPeople("Aya Hirano")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fp == 0 {
    fmt.Println("people not found")
    return
  }

  p, err := c.People(fp)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(
    p.Id, p.Name, p.Japanese, p.Job_title, p.Website,
    p.Birth_on.Day, p.Birth_on.Month, p.Birth_on.Year,
  )

  for _, v := range p.Groupped_roles {
    fmt.Println(v[0], v[1])
  }

  for _, v := range p.Roles {
    for _, vv := range v.Characters {
      fmt.Println(vv.Id, vv.Name)
    }
  }

  for _, v := range p.Roles {
    for _, vv := range v.Animes {
      fmt.Println(vv.Id, vv.Name, vv.Score)
    }
  }

  for _, v := range p.Works {
    fmt.Println(v.Anime.Id, v.Anime.Name, v.Anime.Score)
  }
}