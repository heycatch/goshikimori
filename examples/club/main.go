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

// ПРИМЕРЫ ЛУЧШЕ БРАТЬ ПО ОЧЕРЕДИ,
// СЛИШКОМ МНОГО ЗАПРОСОВ В СЕКУНДУ.
func main() {
  c := conf()
  // ExtraLimit относится только к SearchClub().
  e := &g.ExtraLimit{Page: "1", Limit: "2"}
  // ExtraClub относится ко всем другим функциям.
  ec := &g.ExtraClub{Page: "1"}
  // Можно вести более широкий поиск клубов через Page/Limit.
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

  // Находим по названию айди клуба.
  fid, err := c.FastIdClub("Спокойные деньки")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fid == 0 {
    fmt.Println("Club not found")
    return
  }

  // Аниме добавленная в клубе
  sca, err := c.SearchClubAnimes(fid, ec)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range sca {
    fmt.Println(v.Id, v.Name, v.Score)
  }

  // Манга добавленная в клубе.
  scm, err := c.SearchClubMangas(fid, ec)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scm {
    fmt.Println(v.Id, v.Name, v.Score)
  }

  // Участники клуба.
  scmem, err := c.SearchClubMembers(fid)
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, v := range scmem {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }

  // Находим по названию айди клуба.
  fid, err := c.FastIdClub("Самые прекрасные персонажи")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fid == 0 {
    fmt.Println("Club not found")
    return
  }

  // Персонажи добавленные в клубе.
  scc, err := c.SearchClubCharacters(fid, ec)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scc {
    fmt.Println(v.Id, v.Name, v.Russian)
  }

  // Находим по названию айди клуба.
  fid, err := c.FastIdClub("Yuritopia")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fid == 0 {
    fmt.Println("Club not found")
    return
  }

  // Другие клубы добавленные в клубе.
  scl, err := c.SearchClubClubs(fid, ec)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scl {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }

  // Картинки добавленные в клубе.
  sci, err := c.SearchClubImages(fid)
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, v := range sci {
    fmt.Println(v.Id, v.Original_url, v.Can_destroy, v.User_id)
  }

  // Находим по названию айди клуба.
  fid, err := c.FastIdClub("Интерактивы от DaHanka")
  if err != nil {
    fmt.Println(err)
    return
  }
  if fid == 0 {
    fmt.Println("Club not found")
    return
  }

  // Обсуждение в клубе.
  sccl, err := c.SearchClubCollections(fid, ec)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range sccl {
    fmt.Println(v)
  }

  // Присоединиться или ливнуть с клуба.
  cc, err := c.ClubJoin(fid)
  //cc, err := c.ClubLeave(fid)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cc)
}
