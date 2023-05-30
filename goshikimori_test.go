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
  s.Percent = s.getPercent()

  if s.Graph == "" { s.Graph = "#" }

  for i := 0; i < int(s.Percent); i += 1 {
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

  fmt.Printf(
    "\r[%-5s]%3d%% %8d/%d",
    s.Rate, s.Percent, s.Cur, s.Total,
  )
}

func (s *StatusBar) Finish() { fmt.Println() }

func TestConfiguration(t *testing.T) {
  if app_test != "" && tok_test != "" {
    t.Logf("Found: %s and %s", app_test, tok_test)
  } else {
    t.Error("Not found application or key")
    os.Exit(1)
  }
}

func TestOptionsMessages(t *testing.T) {
  empty := Options{Type: "", Page: "", Limit: ""}
  if empty.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Empty OptionsMessages passed")
  } else {
    t.Error("Empty OptionsMessages failed")
  }

  big := Options{Type: "11111111111111", Page: "100002", Limit: "102"}
  if big.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Big OptionsMessages passed")
  } else {
    t.Error("Big OptionsMessages failed")
  }

  zero := Options{Type: "0", Page: "0", Limit: "0"}
  if zero.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Zero OptionsMessages passed")
  } else {
    t.Error("Zero OptionsMessages failed")
  }

  normal := Options{Type: "private", Page: "2", Limit: "10"}
  if normal.OptionsMessages() == "limit=10&page=2&type=private" {
    t.Log("Normal OptionsMessages passed")
  } else {
    t.Error("Normal OptionsMessages failed")
  }
}

func TestOptionsUserHistory(t *testing.T) {
  empty := Options{Page: "", Limit: "", Target_id: "", Target_type: ""}
  if empty.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Empty OptionsUserHistory passed")
  } else {
    t.Error("Empty OptionsUserHistory failed")
  }

  big := Options{Page: "100002", Limit: "102", Target_id: "", Target_type: "11111111111111"}
  if big.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Big OptionsUserHistory passed")
  } else {
    t.Error("Big OptionsUserHistory failed")
  }

  zero := Options{Page: "0", Limit: "0", Target_id: "", Target_type: "0"}
  if zero.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Zero OptionsUserHistory passed")
  } else {
    t.Error("Zero OptionsUserHistory failed")
  }

  normal := Options{Page: "3", Limit: "20", Target_id: "1337", Target_type: "Manga"}
  if normal.OptionsUserHistory() == "limit=20&page=3&target_id=1337&target_type=Manga" {
    t.Log("Zero OptionsUserHistory passed")
  } else {
    t.Error("Zero OptionsUserHistory failed")
  }
}

func TestOptionsUsers(t *testing.T) {
  empty := Options{Page: "", Limit: ""}
  if empty.OptionsUsers() == "limit=1&page=1" {
    t.Log("Empty OptionsUsers passed")
  } else {
    t.Error("Empty OptionsUsers failed")
  }

  big := Options{Page: "100002", Limit: "102"}
  if big.OptionsUsers() == "limit=1&page=1" {
    t.Log("Big OptionsUsers passed")
  } else {
    t.Error("Big OptionsUsers failed")
  }

  zero := Options{Page: "0", Limit: "0"}
  if zero.OptionsUsers() == "limit=1&page=1" {
    t.Log("Zero OptionsUsers passed")
  } else {
    t.Error("Zero OptionsUsers failed")
  }

  normal := Options{Page: "7", Limit: "37"}
  if normal.OptionsUsers() == "limit=37&page=7" {
    t.Log("Normal OptionsUsers passed")
  } else {
    t.Error("Normal OptionsUsers failed")
  }
}

func TestOptionsAnime(t *testing.T) {
  empty := Options{
    Page: "", Limit: "", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  if empty.OptionsAnime() == "kind=&limit=1&page=1&rating=&score=&season=&status=" {
    t.Log("Empty OptionsAnime passed")
  } else {
    t.Error("Empty OptionsAnime failed")
  }

  big := Options{
    Page: "100002", Limit: "52", Kind: "11111111", Status: "111111111",
    Season: "1111111111", Score: "111111111111", Rating: "10",
  }
  if big.OptionsAnime() == "kind=&limit=1&page=1&rating=&score=&season=&status=" {
    t.Log("Big OptionsAnime passed")
  } else {
    t.Error("Big OptionsAnime failed")
  }

  zero := Options{
    Page: "0", Limit: "0", Kind: "0", Status: "0",
    Season: "0", Score: "0", Rating: "0",
  }
  if zero.OptionsAnime() == "kind=&limit=1&page=1&rating=&score=0&season=&status=" {
    t.Log("Zero OptionsAnime passed")
  } else {
    t.Error("Zero OptionsAnime failed")
  }

  normal := Options{
    Page: "2", Limit: "12", Kind: "tv", Status: "released",
    Season: "199x", Score: "8", Rating: "r",
  }
  if normal.OptionsAnime() == "kind=tv&limit=12&page=2&rating=r&score=8&season=199x&status=released" {
    t.Log("Normal OptionsAnime passed")
  } else {
    t.Error("Normal OptionsAnime failed")
  }
}

func TestOptionsManga(t *testing.T) {
  empty := Options{
    Page: "", Limit: "", Kind: "", Status: "",
    Season: "", Score: "",
  }
  if empty.OptionsManga() == "kind=&limit=1&page=1&score=&season=&status=" {
    t.Log("Empty OptionsManga passed")
  } else {
    t.Error("Empty OptionsManga failed")
  }

  big := Options{
    Page: "100002", Limit: "52", Kind: "11111111", Status: "111111111",
    Season: "1111111111", Score: "1111111111",
  }
  if big.OptionsManga() == "kind=&limit=1&page=1&score=&season=&status=" {
    t.Log("Big OptionsManga passed")
  } else {
    t.Error("Big OptionsManga failed")
  }

  zero := Options{
    Page: "0", Limit: "0", Kind: "0", Status: "0",
    Season: "0", Score: "0",
  }
  if zero.OptionsManga() == "kind=&limit=1&page=1&score=0&season=&status=" {
    t.Log("Zero OptionsManga passed")
  } else {
    t.Error("Zero OptionsManga failed")
  }

  normal := Options{
    Page: "4", Limit: "5", Kind: "manga", Status: "anons",
    Season: "summer_2017", Score: "7",
  }
  if normal.OptionsManga() == "kind=manga&limit=5&page=4&score=7&season=summer_2017&status=anons" {
    t.Log("Normal OptionsManga passed")
  } else {
    t.Error("Normal OptionsManga failed")
  }
}

func TestOptionsClub(t *testing.T) {
  empty := Options{Page: "", Limit: ""}
  if empty.OptionsClub() == "limit=1&page=1" {
    t.Log("Empty OptionsClub passed")
  } else {
    t.Error("Empty OptionsClub failed")
  }

  big := Options{Page: "100002", Limit: "32"}
  if big.OptionsClub() == "limit=1&page=1" {
    t.Log("Big OptionsClub passed")
  } else {
    t.Error("Big OptionsClub failed")
  }

  zero := Options{Page: "0", Limit: "0"}
  if zero.OptionsClub() == "limit=1&page=1" {
    t.Log("Zero OptionsClub passed")
  } else {
    t.Error("Zero OptionsClub failed")
  }

  normal := Options{Page: "2", Limit: "22"}
  if normal.OptionsClub() == "limit=22&page=2" {
    t.Log("Normal OptionsClub passed")
  } else {
    t.Error("Normal OptionsClub failed")
  }
}

func TestOptionsCalendar(t *testing.T) {
  empty := Options{Censored: ""}
  if empty.OptionsCalendar() == "censored=false" {
    t.Log("Empty OptionsCalendar passed")
  } else {
    t.Error("Empty OptionsCalendar failed")
  }

  big := Options{Censored: "11111111111"}
  if big.OptionsCalendar() == "censored=false" {
    t.Log("Big OptionsCalendar passed")
  } else {
    t.Error("Big OptionsCalendar failed")
  }

  zero := Options{Censored: "0"}
  if zero.OptionsCalendar() == "censored=false" {
    t.Log("Zero OptionsCalendar passed")
  } else {
    t.Error("Zero OptionsCalendar failed")
  }

  normal := Options{Censored: "true"}
  if normal.OptionsCalendar() == "censored=true" {
    t.Log("Normal OptionsCalendar passed")
  } else {
    t.Error("Normal OptionsCalendar failed")
  }
}

func TestOptionsAnimeRates(t *testing.T) {
  empty := Options{Page: "", Limit: "", Status: "", Censored: ""}
  if empty.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Empty OptionsAnimeRates passed")
  } else {
    t.Error("Empty OptionsAnimeRates failed")
  }

  big := Options{Page:"100002", Limit: "5002", Status: "11111111", Censored: "1111111111"}
  if big.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Big OptionsAnimeRates passed")
  } else {
    t.Error("Big OptionsAnimeRates failed")
  }

  zero := Options{Page:"0", Limit: "0", Status: "0", Censored: "0"}
  if zero.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Zero OptionsAnimeRates passed")
  } else {
    t.Error("Zero OptionsAnimeRates failed")
  }

  normal := Options{Page:"15", Limit: "405", Status: "dropped", Censored: "true"}
  if normal.OptionsAnimeRates() == "censored=true&limit=405&page=15&status=dropped" {
    t.Log("Normal OptionsAnimeRates passed")
  } else {
    t.Error("Normal OptionsAnimeRates failed")
  }
}

func TestOptionsMangaRates(t *testing.T) {
  empty := Options{Page: "", Limit: "", Censored: ""}
  if empty.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Empty OptionsMangaRates passed")
  } else {
    t.Error("Empty OptionsMangaRates failed")
  }

  big := Options{Page: "100002", Limit: "5002", Censored: "1111111"}
  if big.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Big OptionsMangaRates passed")
  } else {
    t.Error("Big OptionsMangaRates failed")
  }

  zero := Options{Page: "0", Limit: "0", Censored: "0"}
  if zero.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Zero OptionsMangaRates passed")
  } else {
    t.Error("Zero OptionsMangaRates failed")
  }

  normal := Options{Page: "33", Limit: "25", Censored: "true"}
  if normal.OptionsMangaRates() == "censored=true&limit=25&page=33" {
    t.Log("Normal OptionsMangaRates passed")
  } else {
    t.Error("Normal OptionsMangaRates failed")
  }
}

func TestOptionsPeople(t *testing.T) {
  empty := Options{Kind: ""}
  if empty.OptionsPeople() == "kind=seyu" {
    t.Log("Empty OptionsPeople passed")
  } else {
    t.Error("Empty OptionsPeople failed")
  }

  big := Options{Kind: "111111111"}
  if big.OptionsPeople() == "kind=seyu" {
    t.Log("Big OptionsPeople passed")
  } else {
    t.Error("Big OptionsPeople failed")
  }

  zero := Options{Kind: "0"}
  if zero.OptionsPeople() == "kind=seyu" {
    t.Log("Zero OptionsPeople passed")
  } else {
    t.Error("Zero OptionsPeople failed")
  }

  normal := Options{Kind: "mangaka"}
  if normal.OptionsPeople() == "kind=mangaka" {
    t.Log("Normal OptionsPeople passed")
  } else {
    t.Error("Normal OptionsPeople failed")
  }
}

func TestOptionsClubInformation(t *testing.T) {
  empty := Options{Page: ""}
  if empty.OptionsClubInformation() == "page=1" {
    t.Log("Empty OptionsClubInformation passed")
  } else {
    t.Error("Empty OptionsClubInformation failed")
  }

  big := Options{Page: "100002"}
  if big.OptionsClubInformation() == "page=1" {
    t.Log("Big OptionsClubInformation passed")
  } else {
    t.Error("Big OptionsClubInformation failed")
  }

  zero := Options{Page: "0"}
  if zero.OptionsClubInformation() == "page=1" {
    t.Log("Zero OptionsClubInformation passed")
  } else {
    t.Error("Zero OptionsClubInformation failed")
  }

  normal := Options{Page: "1337"}
  if normal.OptionsClubInformation() == "page=1337" {
    t.Log("Normal OptionsClubInformation passed")
  } else {
    t.Error("Normal OptionsClubInformation failed")
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
  o := &Options{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  s, _ := c.SearchAnime("Initial D", o)

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
  o := &Options{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
  }
  r, _ := c.SearchManga("Initial D", o)

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
  o := &Options{Page: "1", Limit: "1"}
  r, _ := c.SearchClub("milf thred", o)

  for _, v := range r {
    if v.Is_censored == true {
      t.Logf("Best club: %s - found", v.Name)
    } else {
      t.Errorf("Argument: %v or Id: %d - not found", v.Is_censored, v.Id)
    }
  }
}

func TestAchievements(t *testing.T) {
  fmt.Println("Too many requests at once, waiting 5 seconds...")

  var s StatusBar
  s.NewOption(0, 5)
  for i := 0; i <= 5; i++ {
    s.Play(int(i))
    time.Sleep(1 * time.Second)
  }
  s.Finish()

  c := conf()
  u, _ := c.FastIdUser("incarnati0n").SearchAchievement()

  for _, v := range u {
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
  a, _ := c.FastIdAnime("initial d first stage").SearchAnimeVideos()

  for _, v := range a {
    if v.Id == 24085 {
      t.Logf("Video: %s", v.Name)
    } else {
      t.Log("Video not found, waiting...")
    }
  }
}

func TestUserUnreadMessages(t *testing.T) {
  fmt.Println("Too many requests at once, waiting 5 seconds...")

  var s StatusBar
  s.NewOption(0, 5)
  for i := 0; i <= 5; i++ {
    s.Play(int(i))
    time.Sleep(1 * time.Second)
  }
  s.Finish()

  c := conf()
  um, _ := c.FastIdUser("incarnati0n").UserUnreadMessages()

  if um.News > 0 || um.News == 0 {
    t.Logf("Found: %d news", um.News)
  } else {
    t.Error("News not found")
  }
}

func TestConstantsAnime(t *testing.T) {
  c := conf()
  ca, _ := c.SearchConstantsAnime()

  if ca.Kind != nil {
    t.Logf("Found: %s", ca.Kind)
  } else {
    t.Error("Constants not found")
  }
  if ca.Status != nil {
    t.Logf("Found: %s", ca.Status)
  } else {
    t.Error("Constants not found")
  }
}

func TestConstantsManga(t *testing.T) {
  c := conf()
  cm, _ := c.SearchConstantsManga()

  if cm.Kind != nil {
    t.Logf("Found: %s", cm.Kind)
  } else {
    t.Error("Constants not found")
  }
  if cm.Status != nil {
    t.Logf("Found: %s", cm.Status)
  } else {
    t.Error("Constants not found")
  }
}

func TestPeople(t *testing.T) {
  fmt.Println("Too many requests at once, waiting 5 seconds...")

  var s StatusBar
  s.NewOption(0, 5)
  for i := 0; i <= 5; i++ {
    s.Play(int(i))
    time.Sleep(1 * time.Second)
  }
  s.Finish()

  c := conf()
  p, _ := c.FastIdPeople("Aya Hirano").People()

  if p.Id == 4 || p.Job_title == "Сэйю"  {
    t.Logf("%s - found", p.Name)
  } else {
    t.Error("People not found")
  }
}
