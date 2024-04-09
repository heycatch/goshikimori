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

// Too many examples, take turns using.
func main() {
  c := conf()
  // First *Options->SearchClub().
  o := &g.Options{Page: 1, Limit: 2}
  // Second *Options->SearchClub...().
  oo := &g.Options{Page: 1}

  // PART 1
  a, status, err := c.SearchClubs("milf", o)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  if len(a) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range a {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }

  // PART 2
  fast_spok, status, err := c.FastIdClub("Спокойные деньки")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  // Finding anime titles in the club.
  sca, err := fast_spok.SearchClubAnimes(oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(sca) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range sca {
    fmt.Println(v.Id, v.Name, v.Score)
  }

  // Finding manga titles in the club.
  scm, err := fast_spok.SearchClubMangas(oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(scm) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range scm {
    fmt.Println(v.Id, v.Name, v.Score)
  }

  // Club members.
  scmem, err := fast_spok.SearchClubMembers()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(scmem) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range scmem {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }

  // PART 3
  fast_pers, status, err := c.FastIdClub("Самые прекрасные персонажи")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  // Added characters in the club.
  scc, err := fast_pers.SearchClubCharacters(oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(scc) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range scc {
    fmt.Println(v.Id, v.Name, v.Russian)
  }

  // PART 4
  fast_yuri, status, err := c.FastIdClub("Yuritopia")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  // Third-party added clubs.
  scl, err := fast_yuri.SearchClubClubs(oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(scl) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range scl {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }
  // Added pictures in the club.
  sci, err := fast_yuri.SearchClubImages()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(sci) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range sci {
    fmt.Println(v.Id, v.Original_url, v.Can_destroy, v.User_id)
  }

  // PART 5
  // Discussion at the club.
  fast_inter, status, err := c.FastIdClub("Интерактивы от DaHanka")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  sccl, err := fast_inter.SearchClubCollections(oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(sccl) == 0 {
    fmt.Println("Club not found")
    return
  }
  for _, v := range sccl {
    fmt.Println(v.Body, v.Created_at, v.Comments_count)
  }

  // PART 6
  // Join/leave from club.
  fast_faq, status, err := c.FastIdClub("FAQ - Часто задаваемые вопросы")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  cc, err := fast_faq.ClubJoin()
  //cc, err := fast_fqa.ClubLeave()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cc)
}
