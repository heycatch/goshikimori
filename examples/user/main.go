package main

import (
  "fmt"
  "time"

  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  user := "incarnati0n"
  c := conf()
  u, err := c.SearchUser(user)
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Printf("Not found %s\n", user)
    return
  }
  fmt.Println(u.Id, u.Sex, u.Last_online, u.Name)
  fmt.Println()
  for _, v := range u.Stats.Statuses.Anime {
    fmt.Println(v.Id, v.Grouped_id, v.Name, v.Size, v.Type)
  }
  fmt.Println()
  for _, v := range u.Stats.Statuses.Manga {
    fmt.Println(v.Id, v.Grouped_id, v.Name, v.Size, v.Type)
  }
  fmt.Println()
  // Search clubs
  uc, err := c.SearchUserClubs(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(uc) == 0 {
    fmt.Println("clubs not found")
    return
  }
  for _, v := range uc {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }
  fmt.Println()
  // Search friends
  uf, err := c.SearchUserFriends(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(uf) == 0 {
    fmt.Println("friends not found")
    return
  }
  for _, v := range uf {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }
  fmt.Println()
  // Search anime and manga rates
  time.Sleep(5 * time.Second)
  fmt.Println("too many requests, wait 5 seconds")
  gar := &g.ExtraAnimeRates{Limit: "5", Status: "completed", Censored: ""}
  ar, err := c.SearchUserAnimeRates(u.Id, gar)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(ar) == 0 {
    fmt.Println("not found anime_rates")
    return
  }
  for _, v := range ar {
    fmt.Println(v.Status, v.Anime.Name, v.Episodes, v.Score)
  }
  fmt.Println()
  gmr := &g.ExtraMangaRates{Limit: "5", Censored: ""}
  mr, err := c.SearchUserMangaRates(u.Id, gmr)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(mr) == 0 {
    fmt.Println("not found manga_rates")
    return
  }
  for _, v := range mr {
    fmt.Println(v.Status, v.Manga.Name, v.Chapters, v.Volumes, v.Score)
  }
}
