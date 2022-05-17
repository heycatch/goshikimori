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
  a := c.SearchAnime("initial d first stage")
  v := c.SearchAnimeVideos(a.Id)

  if v.Id == 24085 {
    t.Logf("Found %s", v.Name)
  } else {
    t.Log("Videos not found")
  }
}

func TestExtraAnimeSearch(t *testing.T) {
  if ok := start(); ok == false {
    t.Log("not found application or key")
  }

  c := conf()
  e := &Extra{
    Limit: "1", Kind: "", Status: "released",
    Season: "199x", Score: "", Rating: "",
  }
  a := c.ExtraSearchAnime("initial d", e)
  for _, v := range a {
    if v.Released_on == "1998-12-06" {
      t.Logf("%s found", v.Name)
    } else {
      t.Log("Anime not found")
    }
  }
}
