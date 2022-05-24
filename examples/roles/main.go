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

func rolesAnime() {
  c := conf()
  f, err := c.FastIdAnime("naruto")
  if err != nil {
    fmt.Println(err)
    return
  }
  if f == 0 {
    fmt.Println("Anime not found")
    return
  }
  r, err := c.SearchMangaRoles(f)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(r) == 0 {
    fmt.Println("Roles not found")
    return
  }
  for _, v := range r {
    fmt.Println(
      v.Roles, v.Roles_Russian,
      v.Character.Id, v.Character.Name,
    )
  }
}

func rolesManga() {
  c := conf()
  f, err := c.FastIdManga("naruto")
  if err != nil {
    fmt.Println(err)
    return
  }
  if f == 0 {
    fmt.Println("Manga not found")
    return
  }
  r, err := c.SearchMangaRoles(f)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(r) == 0 {
    fmt.Println("Roles not found")
    return
  }
  for _, v := range r {
    fmt.Println(
      v.Roles, v.Roles_Russian,
      v.Character.Id, v.Character.Name,
    )
  }
}

func main() {
  rolesAnime()
  rolesManga()
}
