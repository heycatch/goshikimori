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
  o := &g.Options{Kind: "seyu"}
  sp, status, err := c.SearchPeople("Aya Hirano", o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    for _, v := range sp {
      fmt.Println(v.Id, v.Name, v.Russian, v.Image.Original)
    }
  } else {
    fmt.Println(status)
  }
  fast, status, err := c.FastIdPeople("Aya Hirano")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    p, err := fast.People()
    if err != nil {
      fmt.Println(err)
      return
    }
    if p.Id == 0 {
      fmt.Println("people not found")
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
  } else {
    fmt.Println(status)
  }
}
