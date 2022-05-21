package goshikimori

import (
  "time"
  "testing"
  "fmt"
)

type StatusBar struct {
  Percent int
  Cur     int
  Total   int
  Rate    string
  Graph   string
}

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

func (s *StatusBar) NewOption(start, end int) {
  s.Cur = start
  s.Total = end

  if s.Graph == "" {
    s.Graph = "#"
  }

  s.Percent = s.getPercent()

  for i := 0; i < int(s.Percent); i += 2 {
    s.Rate += s.Graph
  }
}

func (s *StatusBar) getPercent() int {
  return int((float32(s.Cur) / float32(s.Total)) * 100)
}

func (s *StatusBar) Play(cur int) {
  s.Cur = cur
  last := s.Percent
  s.Percent = s.getPercent()

  if s.Percent != last && s.Percent%2 == 0 {
    s.Rate += s.Graph
  }

  fmt.Printf("\r[%-10s]%3d%% %8d/%d",
    s.Rate, s.Percent, s.Cur, s.Total)
}

func (s *StatusBar) Finish() { fmt.Println() }

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
  e := &Extra{Limit: "1"}
  s := c.SearchAnime("Initial D", e)

  for _, v := range s {
    if v.Id == 12725 && v.Status == "released" {
      t.Logf("Anime %s id %d - found", v.Name, v.Id)
    } else {
      t.Errorf("Anime %s id %d - not found", v.Name, v.Id)
    }
  }
}

func TestMangas(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  e := &Extra{Limit: "1"}
  r := c.SearchManga("Initial D", e)

  for _, v := range r {
    if v.Volumes == 48 && v.Chapters == 724 {
      t.Logf("Manga %s id %d - found", v.Name, v.Id)
    } else {
      t.Errorf("Manga %s id %d - not found", v.Name, v.Id)
    }
  }
}

func TestClub(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  e := &Extra{Limit: "1"}
  r := c.SearchClub("milf thred", e)

  for _, v := range r {
    if v.Is_censored == true {
      t.Logf("Best club %s - found", v.Name)
    } else {
      t.Errorf("Argument %v or id %d - not found", v.Is_censored, v.Id)
    }
  }
}

func TestAchievements(t *testing.T) {
  var s StatusBar

  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  fmt.Println("Too many requests at once, waiting 10 seconds...")
  s.NewOption(0, 10)
  for i := 0; i <= 10; i++ {
    s.Play(int(i))
    time.Sleep(1 * time.Second)
  }
  s.Finish()

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

  c := conf()
  f := c.FastIdAnime("initial d first stage")
  w := c.SearchAnimeVideos(f)

  for _, v := range w {
    if v.Id == 24085 {
      t.Logf("Found %s", v.Name)
    } else {
      t.Log("Videos not found")
    }
  }
}
