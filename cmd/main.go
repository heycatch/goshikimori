package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "",
    "",
  )
}

func main() {
  c := conf()

  /*
  // ONE
  g, err := c.SearchGenres()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(g) == 0 {
    fmt.Println("not found genres")
    return
  }
  for _, v := range g {
    fmt.Println(v.Id, v.Name, v.Russian, v.Kind)
  }
  // TWO
  s, err := c.SearchStudios()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(s) == 0 {
    fmt.Println("not found studios")
    return
  }
  for _, v := range s {
    fmt.Println(v.Id, v.Name, v.Filtered_name, v.Real)
  }
  // THREE
  p, err := c.SearchPublishers()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(p) == 0 {
    fmt.Println("not found publishers")
    return
  }
  for _, v := range p {
    fmt.Println(v.Id, v.Name)
  }
  // FOUR
  f, err := c.SearchForums()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(f) == 0 {
    fmt.Println("not found forums")
    return
  }
  for _, v := range f {
    fmt.Println(v.Id, v.Position, v.Name, v.Permalink, v.Url)
  }
  */

  /*
  fa, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fa == 0 {
    fmt.Println("anime not found")
    return
  }
  fra, err := c.SearchAnimeExternalLinks(fa)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(fra) == 0 {
    fmt.Println("external links not found")
    return
  }
  for _, v := range fra {
    fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
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
  frm, err := c.SearchMangaExternalLinks(fm)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(frm) == 0 {
    fmt.Println("external links not found")
    return
  }
  for _, v := range frm {
    fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
  }
  */

  /*
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
  */

  /*
  e := &g.ExtraLimit{Page: "1", Limit: "10"}
  u, err := c.SearchUsers("angel", e)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range u {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }
  */

  /*
  u, err := c.SearchUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
    return
  }

  ub, err := c.SearchUserBans(u.Id)
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

  e := &g.Extra{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  a, err := c.SearchAnime("Watashi ga Motenai no wa Dou Kangaetemo Omaera ga Warui!", e)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("anime not found")
    return
  }
  var res string
  for _, v := range a {
    res = fmt.Sprintf("%d", v.Id)
  }

  ett := &g.ExtraTargetType{Page: "1", Limit: "15", Target_id: res, Target_type: "Anime"}
  h, err := c.SearchUserHistory(u.Id, ett)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(h) == 0 {
    fmt.Println("history not found")
    return
  }
  for _, v := range h {
    fmt.Println(v.Id, v.Description, v.Target.Russian, v.Target.Episodes)
  }

  uf, err := c.SearchUserFavourites(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(uf.Animes) == 0 {
    fmt.Println("favourite animes not found")
    return
  }
  for _, v := range uf.Animes {
    fmt.Println(v.Id, v.Name, v.Russian, v.Image)
  }
  if len(uf.Mangas) == 0 {
    fmt.Println("favourite mangas not found")
    return
  }
  for _, v := range uf.Animes {
    fmt.Println(v.Id, v.Name, v.Russian, v.Image)
  }

  gar := &g.ExtraAnimeRates{Page: "1", Limit: "5", Status: "completed", Censored: ""}
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

  gmr := &g.ExtraMangaRates{Page: "1", Limit: "5", Censored: ""}
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
  */

  /*
  f, err := c.SearchUserClubs(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(f) == 0 {
    fmt.Println("clubs not found")
    return
  }
  for _, v := range f {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }
  */
  /*
  f, err := c.SearchUserFriends(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(f) == 0 {
    fmt.Println("friends not found")
    return
  }
  for _, v := range f {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }
  */

  /*
  u, err := c.SearchUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
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
  */

  /*
  w, err := c.WhoAmi()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(w.Nickname, w.Avatar, w.Locale, w.Last_online_at)
  */

  /*
  w := &g.ExtraCensored{Censored: "false"}
  ca, err := c.SearchCalendar(w)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range ca {
    fmt.Println(
      v.Next_episode, v.Next_episode_at, v.Duration,
      v.Anime.Id, v.Anime.Name, v.Anime.Score,
    )
  }
  */

  /*
  e := &g.ExtraLimit{Page: "1", Limit: ""}
  a, err := c.SearchClub("milf", e)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }
  */

  /*
  b, err := c.SearchBans()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(b) == 0 {
    fmt.Println("bans not found")
    return
  }
  for _, v := range b {
    fmt.Println(v.Comment, v.User.Id, v.User.Nickname,
      v.Moderator.Id, v.Moderator.Nickname,
    )
  }
  */

  /*
  e := &g.Extra{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  a, err := c.SearchAnime("initial d first stage", e)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("anime not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Id, v.Name, v.Released_on, v.Score)
  }
  */

  /*
  e := &g.Extra{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "",
  }
  a, err := c.SearchManga("initial d", e)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(a) == 0 {
    fmt.Println("manga not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Name, v.Released_on, v.Score)
  }
  */

  /*
  u := c.SearchUser("incarnati0n")
  r := c.SearchAchievement(u.Id)
  for _, v := range r {
    if v.Neko_id == g.NekoSearch("initial d") {
      fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
      fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
    }
  }
  */

  /*
  // PART_1
  a := c.FastIdAnime("vampire knight")
  s := c.SearchSimilarAnime(a)
  for _, v := range s {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
  // PART_2
  aa := c.FastIdManga("initial d")
  ss := c.SearchSimilarManga(aa)
  for _, v := range ss {
    fmt.Println(v.Name, v.Id, v.Russian)
  }
  */

  /*
  // FOR BAD TEST
  r1 := c.SearchManga("Zhu Zhu Xia")
  s1 := c.SearchSimilarManga(r1.Id)
  fmt.Println(s1.Name, s1.Id, s1.Russian)
  */

  /*
  // PART_1
  a, _ := c.FastIdAnime("initial d second stage")
  r, _ := c.SearchRelatedAnime(a)
  for _, v := range r {
    fmt.Println(
      v.Relation, v.Relation_Russian, v.Anime.Name,
      v.Anime.Russian, v.Anime.Score,
    )
  }
  // PART_2
  mm := c.FastIdManga("vampire knight")
  rr := c.SearchRelatedManga(mm)
  for _, v := range rr {
    fmt.Println(
      v.Relation, v.Relation_Russian,
      v.Manga.Score, v.Manga.Status,
    )
  }
  */

  /*
  a, err := c.FastIdAnime("initial d")
  if err != nil {
    fmt.Println(err)
    return
  }
  if a == 0 {
    fmt.Println("anime not found")
    return
  }
  s, err := c.SearchAnimeScreenshots(a)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(s) == 0 {
    fmt.Println("screenshots not found")
    return
  }
  for _, v := range s {
    fmt.Println(v.Original, v.Preview)
  }
  */

  /*
  //f := c.FastIdAnime("initial d first stage")
  //f := c.FastIdAnime("Jidou Bungaku Library")
  if f == 0 {
    fmt.Println("Id not found")
    return
  }
  vv := c.SearchAnimeVideos(f)
  if len(vv) == 0 {
    fmt.Println("Video not found")
    return
  }
  for _, v := range vv {
    fmt.Println(v.Id, v.Url, v.Image_url, v.Player_url, v.Name, v.Kind, v.Hosting)
  }
  */

  /* FML
  for _, v := range a.([]interface{}) {
    //fmt.Println(v.(map[string]interface{})["name"])
    r := v.(map[string]interface{})
    fmt.Println(r["id"], r["name"])
  }
  */

  /*
  //f, err := c.FastIdAnime("naruto")
  //r, err := c.SearchAnimeRoles(f)
  f, err := c.FastIdManga("naruto")
  if err != nil {
    fmt.Println(err)
    return
  }
  r, err := c.SearchMangaRoles(f)
  if len(r) == 0 {
    fmt.Println("Id not found")
    return
  }
  for _, v := range r {
    fmt.Println(
      v.Roles, v.Roles_Russian,
      v.Character.Id, v.Character.Name,
    )
  }
  */

  /*
  u, err := c.SearchUser("morr")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
    return
  }

  fr, err := c.AddFriend(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(fr.Notice)

  fr, err := c.RemoveFriend(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(fr.Notice)
  */

  /*
  u, err := c.SearchUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
    return
  }

  um, err := c.UserUnreadMessages(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(um.Messages, um.News, um.Notifications)

  e := &g.ExtraMessages{Type: "news", Page: "1", Limit: "1"}
  m, err := c.UserMessages(u.Id, e)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range m {
    fmt.Println(v)
  }
  */

  ca, err := c.SearchConstantsAnime()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(ca.Kind, ca.Status)

  cm, err := c.SearchConstantsManga()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cm.Kind, cm.Status)

  ur, err := c.SearchConstantsUserRate()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(ur.Status)

  cc, err := c.SearchConstantsClub()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cc.Join_policy, cc.Comment_policy, cc.Image_upload_policy)

  cs, err := c.SearchConstantsSmileys()
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range cs {
    fmt.Println(v.Bbcode, v.Path)
  }
}
