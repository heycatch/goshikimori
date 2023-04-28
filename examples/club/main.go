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
  // Find the Id by club-name.
  id_den, err := c.FastIdClub("Спокойные деньки")
  if err != nil {
    fmt.Println(err)
    return
  }
  if id_den == 0 {
    fmt.Println("Club not found")
    return
  }

  // Finding anime titles in the club.
  sca, err := c.SearchClubAnimes(id_den, oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range sca {
    fmt.Println(v.Id, v.Name, v.Score)
  }

  // Finding manga titles in the club.
  scm, err := c.SearchClubMangas(id_den, oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scm {
    fmt.Println(v.Id, v.Name, v.Score)
  }

  // Club members.
  scmem, err := c.SearchClubMembers(id_den)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scmem {
    fmt.Println(v.Id, v.Nickname, v.Last_online_at)
  }

  // PART 3
  // Find the Id by club-name.
  id_per, err := c.FastIdClub("Самые прекрасные персонажи")
  if err != nil {
    fmt.Println(err)
    return
  }
  if id_per == 0 {
    fmt.Println("Club not found")
    return
  }

  // Added characters in the club.
  scc, err := c.SearchClubCharacters(id_per, oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scc {
    fmt.Println(v.Id, v.Name, v.Russian)
  }

  // PART 4
  // Find the Id by club-name.
  id_yur, err := c.FastIdClub("Yuritopia")
  if err != nil {
    fmt.Println(err)
    return
  }
  if id_yur == 0 {
    fmt.Println("Club not found")
    return
  }

  // Third-party added clubs.
  scl, err := c.SearchClubClubs(id_yur, oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range scl {
    fmt.Println(v.Id, v.Name, v.Is_censored)
  }

  // Added pictures in the club.
  sci, err := c.SearchClubImages(id_yur)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range sci {
    fmt.Println(v.Id, v.Original_url, v.Can_destroy, v.User_id)
  }

  // PART 5
  // Find the Id by club-name.
  id_dah, err := c.FastIdClub("Интерактивы от DaHanka")
  if err != nil {
    fmt.Println(err)
    return
  }
  if id_dah == 0 {
    fmt.Println("Club not found")
    return
  }

  // Discussion at the club.
  sccl, err := c.SearchClubCollections(id_dah, oo)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range sccl {
    fmt.Println(v)
  }

  // PART 6
  // Join/leave from club.
  cc, err := c.ClubJoin(id_dah)
  //cc, err := c.ClubLeave(id_dah)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cc)
}
