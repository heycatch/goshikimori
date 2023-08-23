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
  o := &g.Options{Page: "1", Limit: "2"}
  // Second *Options->SearchClub...().
  oo := &g.Options{Page: "1"}

  // PART 1
  a, status, err := c.SearchClub("milf", o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    if len(a) == 0 {
      fmt.Println("Club not found")
      return
    }
    for _, v := range a {
      fmt.Println(v.Id, v.Name, v.Is_censored)
    }
  } else {
    fmt.Println(status)
  }

  // PART 2
  fast_spok, status, err := c.FastIdClub("Спокойные деньки")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
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
  } else {
    fmt.Println(status)
  }

  // PART 3
  fast_pers, status, err := c.FastIdClub("Самые прекрасные персонажи")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
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
  } else {
    fmt.Println(status)
  }

  // PART 4
  fast_yuri, status, err := c.FastIdClub("Yuritopia")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
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
  } else {
    fmt.Println(status)
  }

  // PART 5
  // Discussion at the club.
  fast_inter, status, err := c.FastIdClub("Интерактивы от DaHanka")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
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
  } else {
    fmt.Println(status)
  }

  // PART 6
  // Join/leave from club.
  fast_faq, status, err := c.FastIdClub("FAQ - Часто задаваемые вопросы")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    cc, err := fast_faq.ClubJoin()
    //cc, err := fast_fqa.ClubLeave()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(cc)
  } else {
    fmt.Println(status)
  }
}
