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
  f := c.FastIdAnime("naruto")
  r := c.SearchMangaRoles(f)
  for _, v := range r {
    fmt.Println(
      v.Roles, v.Roles_Russian,
      v.Character.Id, v.Character.Name,
    )
  }
}

func rolesManga() {
  c := conf()
  f := c.FastIdManga("naruto")
  r := c.SearchMangaRoles(f)
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
