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

// Too many examples, take turns using.
func main() {
  c := conf()
  // First *Options->SearchClub().
  o := &g.Options{Page: "1", Limit: "2"}
  // Second *Options->SearchClub...().
  oo := &g.Options{Page: "1"}

  // PART 1
  a, err := c.SearchClub("milf", o)
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

  // PART 2
  // Finding anime titles in the club.
  sca, err := c.FastIdClub("Спокойные деньки").SearchClubAnimes(oo)
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
  scm, err := c.FastIdClub("Спокойные деньки").SearchClubMangas(oo)
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
  scmem, err := c.FastIdClub("Спокойные деньки").SearchClubMembers()
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
  // Added characters in the club.
  scc, err := c.FastIdClub("Самые прекрасные персонажи").SearchClubCharacters(oo)
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
  // Third-party added clubs.
  scl, err := c.FastIdClub("Yuritopia").SearchClubClubs(oo)
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
  sci, err := c.FastIdClub("Yuritopia").SearchClubImages()
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
  sccl, err := c.FastIdClub("Интерактивы от DaHanka").SearchClubCollections(oo)
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
  cc, err := c.FastIdClub("FAQ - Часто задаваемые вопросы").ClubJoin()
  //cc, err := c.FastIdClub("FAQ - Часто задаваемые вопросы").ClubLeave()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cc)
}
