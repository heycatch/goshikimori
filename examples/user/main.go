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

  u, status, err := c.SearchUser("arctica")
  if status != 200 || err != nil || u.Id == 0 {
    fmt.Println(err)
    return
  }
  // user info
  fmt.Println(u.Id, u.Sex, u.Last_online, u.Name, u.Image.X160)
  // plan to watch anime
  for _, v := range u.Stats.Statuses.Anime {
    fmt.Println(v.Id, v.Grouped_id, v.Name, v.Size, v.Type)
  }
  // plan to read manga
  for _, v := range u.Stats.Statuses.Manga {
    fmt.Println(v.Id, v.Grouped_id, v.Name, v.Size, v.Type)
  }

  // user clubs
  fast, status, err := c.FastIdUser("arctica")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  uc, err := fast.SearchUserClubs()
  if err != nil || len(uc) == 0 {
    fmt.Println(err)
    return
  }
  for _, v := range uc {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }
  // user friends
  ufo := &g.Options{Page: 1, Limit: 5}
  uf, err := fast.SearchUserFriends(ufo)
  if err != nil || len(uf) == 0{
    fmt.Println(err)
    return
  }
  for _, v := range uf {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }
  // search anime and manga rates
  oar := &g.Options{Page: 1, Limit: 5, Status: "completed", Censored: true}
  ar, err := fast.SearchUserAnimeRates(oar)
  if err != nil || len(ar) == 0 {
    fmt.Println(err)
    return
  }
  for _, v := range ar {
    fmt.Println(v.Status, v.Anime.Name, v.Episodes, v.Score)
  }
  omr := &g.Options{Page: 1, Limit: 5, Censored: true}
  mr, err := fast.SearchUserMangaRates(omr)
  if err != nil || len(mr) == 0 {
    fmt.Println(err)
    return
  }
  for _, v := range mr {
    fmt.Println(v.Status, v.Manga.Name, v.Chapters, v.Volumes, v.Score)
  }
  // search favourites: anime, manga, characters, people, mangakas, seyu and producers
  suf, err := fast.SearchUserFavourites()
  if err != nil || len(suf.Animes) == 0 || len(suf.Mangas) == 0 {
    fmt.Println(err)
    return
  }
  for _, v := range suf.Animes {
    fmt.Println(v.Id, v.Name, v.Russian, v.Image)
  }
  for _, v := range suf.Mangas {
    fmt.Println(v.Id, v.Name, v.Russian, v.Image)
  }
  // user history
  // NOTES: Target_id - Anime.id or Manga.id; convert to a string to search the history point-by-point.
  ouh := &g.Options{Page: 1, Limit: 10, Target_id: "", Target_type: "Anime"}
  uh, err := fast.SearchUserHistory(ouh)
  if err != nil || len(uh) == 0 {
    fmt.Println(err)
    return
  }
  for _, v := range uh {
    fmt.Println(v.Id, v.Description, v.Target.Russian, v.Target.Episodes)
  }
  // user bans
  ub, err := fast.SearchUserBans()
  if err != nil || len(ub) == 0 {
    fmt.Println(err)
    return
  }
  for _, v := range ub {
    fmt.Println(v.Comment, v.User.Id, v.User.Nickname,
      v.Moderator.Id, v.Moderator.Nickname,
    )
  }
  // Brief user information
  fi, err := fast.UserBriefInfo()
  if err != nil || fi.Id == 0 {
    fmt.Println(err)
    return
  }
  fmt.Println(fi.Id, fi.Name, fi.Last_online_at, fi.Full_years, fi.Birth_on)
}
