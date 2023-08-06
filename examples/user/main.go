package main

import (
  "fmt"
  "time"

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

  u, err := c.SearchUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
    return
  }
  // user info
  fmt.Println(u.Id, u.Sex, u.Last_online, u.Name, u.Image.X160)
  fmt.Println()
  // plan to watch anime
  for _, v := range u.Stats.Statuses.Anime {
    fmt.Println(v.Id, v.Grouped_id, v.Name, v.Size, v.Type)
  }
  fmt.Println()
  // plan to read manga
  for _, v := range u.Stats.Statuses.Manga {
    fmt.Println(v.Id, v.Grouped_id, v.Name, v.Size, v.Type)
  }
  fmt.Println()
  // user clubs
  uc, err := c.FastIdUser("incarnati0n").SearchUserClubs()
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
  // user friends
  uf, err := c.FastIdUser("incarnati0n").SearchUserFriends()
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
  // search anime and manga rates
  time.Sleep(5 * time.Second)
  fmt.Println("too many requests, wait 5 seconds")
  oar := &g.Options{Page: "1", Limit: "5", Status: "completed", Censored: ""}
  ar, err := c.FastIdUser("incarnati0n").SearchUserAnimeRates(oar)
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
  omr := &g.Options{Page: "1", Limit: "5", Censored: ""}
  mr, err := c.FastIdUser("incarnati0n").SearchUserMangaRates(omr)
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
  fmt.Println()
  // search favourites: anime, manga, characters, people, mangakas, seyu and producers
  suf, err := c.FastIdUser("incarnati0n").SearchUserFavourites()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(suf.Animes) == 0 {
    fmt.Println("favourite animes not found")
    return
  }
  for _, v := range suf.Animes {
    fmt.Println(v.Id, v.Name, v.Russian, v.Image)
  }
  if len(suf.Mangas) == 0 {
    fmt.Println("favourite mangas not found")
    return
  }
  for _, v := range suf.Animes {
    fmt.Println(v.Id, v.Name, v.Russian, v.Image)
  }
  // user history
  time.Sleep(5 * time.Second)
  fmt.Println("too many requests, wait 5 seconds")
  // NOTES: Target_id - Anime.id or Manga.id; convert to a string to search the history point-by-point.
  ouh := &g.Options{Page: "1", Limit: "10", Target_id: "", Target_type: "Anime"}
  uh, err := c.FastIdUser("incarnati0n").SearchUserHistory(ouh)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(uh) == 0 {
    fmt.Println("history not found")
    return
  }
  for _, v := range uh {
    fmt.Println(v.Id, v.Description, v.Target.Russian, v.Target.Episodes)
  }
  fmt.Println()
  // user bans
  ub, err := c.FastIdUser("incarnati0n").SearchUserBans()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(ub) == 0 {
    fmt.Println("bans not found")
    return
  }
  for _, v := range ub {
    fmt.Println(v.Comment, v.User.Id, v.User.Nickname,
      v.Moderator.Id, v.Moderator.Nickname,
    )
  }
  time.Sleep(5 * time.Second)
  fmt.Println("too many requests, wait 5 seconds")
  // Brief user information
  fi, err := c.FastIdUser("incarnati0n").UserBriefInfo()
  if err != nil {
    fmt.Println(err)
    return
  }
  if fi.Id == 0 {
    fmt.Println("user not found")
    return
  }
  fmt.Println(fi.Id, fi.Name, fi.Last_online_at, fi.Full_years, fi.Birth_on)
}
