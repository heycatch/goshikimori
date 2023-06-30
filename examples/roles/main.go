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

  a, err := c.FastIdAnime("naruto").SearchAnimeRoles()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("Anime not found")
    return
  }
  for _, v := range a {
    fmt.Println(
      v.Roles, v.Roles_Russian,
      v.Character.Id, v.Character.Name,
    )
  }

  m, err := c.FastIdManga("naruto").SearchMangaRoles()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(m) == 0 {
    fmt.Println("Manga not found")
    return
  }
  for _, v := range m {
    fmt.Println(
      v.Roles, v.Roles_Russian,
      v.Character.Id, v.Character.Name,
    )
  }
}
