package goshikimori

import (
  "time"
  "testing"
  "fmt"
  "os"
)

type StatusBar struct {
  Percent int
  Cur     int
  Total   int
  Rate    string
  Graph   string
}

const (
  app_test = ""
  tok_test = ""
)

func conf() *Configuration { return Add(app_test, tok_test) }

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
    s.Rate, s.Percent, s.Cur, s.Total,
  )
}

func (s *StatusBar) Finish() { fmt.Println() }

func TestStart(t *testing.T) {
  if app_test != "" && tok_test != "" {
    t.Logf("Found: %s and %s", app_test, tok_test)
  } else {
    t.Error("Not found application or key")
    os.Exit(1)
  }
}

func TestUser(t *testing.T) {
  name := "incarnati0n"
  c := conf()
  s, _ := c.SearchUser(name)

  if s.Id == 181833 && s.Sex == "male" {
    t.Logf("User: %s, Id: %d - found", s.Nickname, s.Id)
  } else {
    t.Errorf("User: %s - not found", name)
  }
}

func TestAnimes(t *testing.T) {
  c := conf()
  e := &Extra{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  s, _ := c.SearchAnime("Initial D", e)

  for _, v := range s {
    if v.Id == 12725 && v.Status == "released" {
      t.Logf("Anime: %s, Id: %d - found", v.Name, v.Id)
    } else {
      t.Errorf("Anime: %s, Id: %d - not found", v.Name, v.Id)
    }
  }
}

func TestMangas(t *testing.T) {
  c := conf()
  e := &Extra{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  r, _ := c.SearchManga("Initial D", e)

  for _, v := range r {
    if v.Volumes == 48 && v.Chapters == 724 {
      t.Logf("Manga: %s, Id: %d - found", v.Name, v.Id)
    } else {
      t.Errorf("Manga: %s, Id: %d - not found", v.Name, v.Id)
    }
  }
}

func TestClub(t *testing.T) {
  c := conf()
  e := &ExtraLimit{Page: "1", Limit: "1"}
  r, _ := c.SearchClub("milf thred", e)

  for _, v := range r {
    if v.Is_censored == true {
      t.Logf("Best club: %s - found", v.Name)
    } else {
      t.Errorf("Argument: %v or Id: %d - not found", v.Is_censored, v.Id)
    }
  }
}

func TestAchievements(t *testing.T) {
  fmt.Println("Too many requests at once, waiting 10 seconds...")

  var s StatusBar
  s.NewOption(0, 10)
  for i := 0; i <= 10; i++ {
    s.Play(int(i))
    time.Sleep(1 * time.Second)
  }
  s.Finish()

  c := conf()
  u, _ := c.SearchUser("incarnati0n")
  r, _ := c.SearchAchievement(u.Id)

  for _, v := range r {
    if v.Neko_id == NekoSearch("Initial D") {
      if v.Progress == 100 {
        t.Logf("Found: %d progress", v.Progress)
      } else {
        t.Error("Progress not found")
      }
    }
  }
}

func TestAnimeVideos(t *testing.T) {
  c := conf()
  f, _ := c.FastIdAnime("initial d first stage")
  a, _ := c.SearchAnimeVideos(f)

  for _, v := range a {
    if v.Id == 24085 {
      t.Logf("Video: %s", v.Name)
    } else {
      t.Log("Video not found, waiting...")
    }
  }
}

func TestUserUnreadMessages(t *testing.T) {
  c := conf()
  u, _ := c.SearchUser("incarnati0n")
  um, _ := c.UserUnreadMessages(u.Id)

  if um.News > 0 || um.News == 0 {
    t.Logf("Found: %d news", um.News)
  }
}
