package goshikimori

import (
  "time"
  "testing"
)

const (
  api_test = ""
  key_test = ""
)

func conf() *Configuration {
  return Add(
    api_test,
    key_test,
  )
}

func start() bool {
  if api_test != "" && key_test != "" {
    return true
  }
  return false
}

func TestUser(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  s := c.SearchUser("incarnati0n")

  if s.Id == 181833 && s.Sex == "male" {
    t.Logf("User %s id %d - found", s.Nickname, s.Id)
  } else {
    t.Errorf("User %s id %d - not found", s.Nickname, s.Id)
  }
}

func TestAnimes(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  s := c.SearchAnime("Initial D")

  if s.Id == 12725 && s.Status == "released" {
    t.Logf("Anime %s id %d - found", s.Name, s.Id)
  } else {
    t.Errorf("Anime %s id %d - not found", s.Name, s.Id)
  }
}

func TestMangas(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  r := c.SearchManga("Initial D")

  if r.Volumes == 48 && r.Chapters == 724 {
    t.Logf("Manga %s id %d - found", r.Name, r.Id)
  } else {
    t.Errorf("Manga %s id %d - not found", r.Name, r.Id)
  }
}

func TestRanobe(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  r := c.SearchRanobe("Vampire Knight")

  if r.Volumes == 1 && r.Chapters == 6 {
    t.Logf("Ranobe %s id %d - found", r.Name, r.Id)
  } else {
    t.Errorf("Ranobe %s id %d - not found", r.Name, r.Id)
  }
}

func TestClub(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  r := c.SearchClub("milf thred")

  if r.Is_censored == true {
    t.Logf("Best club %s - found", r.Name)
  } else {
    t.Errorf("Argument %v or id %d - not found", r.Is_censored, r.Id)
  }
}

func TestAchievements(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  u := c.SearchUser("incarnati0n")
  r := c.SearchAchievement(u.Id)

  for _, v := range r {
    if v.Neko_id == NekoSearch("Initial D") {
      if v.Progress == 100 {
        t.Logf("Found %d progress", v.Progress)
      } else {
        t.Errorf("Not found %d progress", v.Progress)
      }
    }
  }
}

func TestAnimeVideos(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  // too many requests at once
  t.Log("Waiting 5 seconds...")
  time.Sleep(5 * time.Second)

  c := conf()
  a := c.SearchAnime("initial d first stage")
  v := c.SearchAnimeVideos(a.Id)

  if v.Id == 24085 {
    t.Logf("Found %s", v.Name)
  } else {
    t.Log("Videos not found")
  }
}
