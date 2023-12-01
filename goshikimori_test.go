package goshikimori

import (
  "os"
  "fmt"
  "time"
  "testing"

  "github.com/heycatch/goshikimori/graphql"
)

type StatusBar struct {
  Percent, Cur, Total int
  Rate, Graph string
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

  for i := 0; i < s.Percent; i += 1 { s.Rate += s.Graph }
}

func (s *StatusBar) getPercent() int {
  return int((float32(s.Cur) / float32(s.Total)) * 100)
}

func (s *StatusBar) Play(cur int) {
  s.Cur = cur
  last := s.Percent
  s.Percent = s.getPercent()

  if s.Percent != last && s.Percent%2 == 0 { s.Rate += s.Graph }

  fmt.Printf("\r[%-5s]%3d%% %8d/%d", s.Rate, s.Percent, s.Cur, s.Total)
}

func TestConfiguration(t *testing.T) {
  if app_test != "" && tok_test != "" {
    t.Logf("Found: %s and %s", app_test, tok_test)
  } else {
    t.Error("Not found application or key")
    os.Exit(1)
  }
}

func TestOptionsTopics(t *testing.T) {
  empty := Options{
    Page: "", Limit: "", Forum: "",
    Linked_id: "", Linked_type: "",
  }
  if empty.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Empty OptionsTopics passed")
  } else {
    t.Error("Empty OptionsTopics failed")
  }

  big := Options{
    Page: "100002", Limit: "32", Forum: "1111111111",
    Linked_id: "222222222", Linked_type: "1111111111111111",
  }
  if big.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  zero := Options{
    Page: "0", Limit: "0", Forum: "0",
    Linked_id: "0", Linked_type: "0",
  }
  if zero.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  negative := Options{
    Page: "-1", Limit: "-1", Forum: "-1",
    Linked_id: "-1", Linked_type: "-1",
  }
  if negative.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  normal_one := Options{
    Page: "5", Limit: "10", Forum: "animanga",
    Linked_id: "342908", Linked_type: "Anime",
  }
  if normal_one.OptionsTopics() == "forum=animanga&limit=10&linked_id=342908&linked_type=Anime&page=5" {
    t.Log("Normal-one OptionsTopics passed")
  } else {
    t.Error("Normal-one OptionsTopics failed")
  }

  normal_two := Options{
    Page: "3", Limit: "8", Forum: "animanga",
    Linked_id: "2323", Linked_type: "Manga",
  }
  if normal_two.OptionsTopics() == "forum=animanga&limit=8&linked_id=2323&linked_type=Manga&page=3" {
    t.Log("Normal-two OptionsTopics passed")
  } else {
    t.Error("Normal-two OptionsTopics failed")
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

  negative := Options{Type: "-1", Page: "-1", Limit: "-1"}
  if negative.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Negative OptionsMessages passed")
  } else {
    t.Error("Negative OptionsMessages failed")
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

  negative := Options{Page: "-1", Limit: "-1", Target_id: "", Target_type: "-1"}
  if negative.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Negative OptionsUserHistory passed")
  } else {
    t.Error("Negative OptionsUserHistory failed")
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

  negative := Options{Page: "-1", Limit: "-1"}
  if negative.OptionsUsers() == "limit=1&page=1" {
    t.Log("Negative OptionsUsers passed")
  } else {
    t.Error("Negative OptionsUsers failed")
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
    Page: "", Limit: "", Order: "", Kind: "", Status: "",
    Season: "", Score: "", Rating: "", Duration: "",
    Censored: "", Mylist: "",
  }
  if empty.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&score=&season=&status=" {
    t.Log("Empty OptionsAnime passed")
  } else {
    t.Error("Empty OptionsAnime failed")
  }

  big := Options{
    Page: "100002", Limit: "52", Order: "111111111", Kind: "11111111", Status: "111111111",
    Season: "1111111111", Score: "111111111111", Rating: "10",
    Duration: "11111111111111", Censored: "111111111111111", Mylist: "1111111111111111",
  }
  if big.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&score=&season=&status=" {
    t.Log("Big OptionsAnime passed")
  } else {
    t.Error("Big OptionsAnime failed")
  }

  zero := Options{
    Page: "0", Limit: "0", Order: "0", Kind: "0", Status: "0",
    Season: "0", Score: "0", Rating: "0", Duration: "0",
    Censored: "0", Mylist: "0",
  }
  if zero.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&score=&season=&status=" {
    t.Log("Zero OptionsAnime passed")
  } else {
    t.Error("Zero OptionsAnime failed")
  }

  negative := Options{
    Page: "-1", Limit: "-1", Order: "-1", Kind: "-1", Status: "-1",
    Season: "-1", Score: "-1", Rating: "-1", Duration: "-1",
    Censored: "-1", Mylist: "-1",
  }
  if negative.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&score=&season=&status=" {
    t.Log("Negative OptionsAnime passed")
  } else {
    t.Error("Negative OptionsAnime failed")
  }

  normal := Options{
    Page: "2", Limit: "12", Order: "id", Kind: "tv", Status: "released",
    Season: "199x", Score: "8", Rating: "r", Duration: "D",
    Censored: "true", Mylist: "watching",
  }
  if normal.OptionsAnime() == "censored=true&duration=D&kind=tv&limit=12&mylist=watching&order=id&page=2&rating=r&score=8&season=199x&status=released" {
    t.Log("Normal OptionsAnime passed")
  } else {
    t.Error("Normal OptionsAnime failed")
  }
}

func TestOptionsManga(t *testing.T) {
  empty := Options{
    Page: "", Limit: "", Order: "", Kind: "", Status: "",
    Season: "", Score: "", Censored: "", Mylist: "",
  }
  if empty.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&score=&season=&status=" {
    t.Log("Empty OptionsManga passed")
  } else {
    t.Error("Empty OptionsManga failed")
  }

  big := Options{
    Page: "100002", Limit: "52", Order: "111111111", Kind: "11111111", Status: "111111111",
    Season: "1111111111", Score: "1111111111",
    Censored: "11111111111", Mylist: "11111111111111",
  }
  if big.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&score=&season=&status=" {
    t.Log("Big OptionsManga passed")
  } else {
    t.Error("Big OptionsManga failed")
  }

  zero := Options{
    Page: "0", Limit: "0", Order: "0", Kind: "0", Status: "0",
    Season: "0", Score: "0", Censored: "0", Mylist: "0",
  }
  if zero.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&score=&season=&status=" {
    t.Log("Zero OptionsManga passed")
  } else {
    t.Error("Zero OptionsManga failed")
  }

  negative := Options{
    Page: "-1", Limit: "-1", Order: "-1", Kind: "-1", Status: "-1",
    Season: "-1", Score: "-1", Censored: "-1", Mylist: "-1",
  }
  if negative.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&score=&season=&status=" {
    t.Log("Negative OptionsManga passed")
  } else {
    t.Error("Negative OptionsManga failed")
  }

  normal := Options{
    Page: "4", Limit: "5", Order: "id", Kind: "manga", Status: "anons",
    Season: "198x", Score: "7", Censored: "false", Mylist: "planned",
  }
  if normal.OptionsManga() == "censored=false&kind=manga&limit=5&mylist=planned&order=id&page=4&score=7&season=198x&status=anons" {
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

  negative := Options{Page: "-1", Limit: "-1"}
  if negative.OptionsClub() == "limit=1&page=1" {
    t.Log("Negative OptionsClub passed")
  } else {
    t.Error("Negative OptionsClub failed")
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

  negative := Options{Censored: "-1"}
  if negative.OptionsCalendar() == "censored=false" {
    t.Log("Negative OptionsCalendar passed")
  } else {
    t.Error("Negative OptionsCalendar failed")
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

  negative := Options{Page:"-1", Limit: "-1", Status: "-1", Censored: "-1"}
  if negative.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Negative OptionsAnimeRates passed")
  } else {
    t.Error("Negative OptionsAnimeRates failed")
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

  negative := Options{Page: "-1", Limit: "-1", Censored: "-1"}
  if negative.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Negative OptionsMangaRates passed")
  } else {
    t.Error("Negative OptionsMangaRates failed")
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

  negative := Options{Kind: "-1"}
  if negative.OptionsPeople() == "kind=seyu" {
    t.Log("Negative OptionsPeople passed")
  } else {
    t.Error("Negative OptionsPeople failed")
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

  negative := Options{Page: "-1"}
  if negative.OptionsClubInformation() == "page=1" {
    t.Log("Negative OptionsClubInformation passed")
  } else {
    t.Error("Negative OptionsClubInformation failed")
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
  s, _, _:= c.SearchUser(name)

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
    Season: "", Score: "", Rating: "", Duration: "",
    Censored: "", Mylist: "",
  }
  s, _, _ := c.SearchAnimes("Initial D", o)

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
    Censored: "", Mylist: "",
  }
  r, _, _ := c.SearchMangas("Initial D", o)

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
  r, _, _ := c.SearchClubs("milf thred", o)

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
    s.Play(i)
    time.Sleep(1 * time.Second)
  }

  c := conf()
  fast, _, _ := c.FastIdUser("incarnati0n")
  u, _ := fast.SearchAchievement()
  neko, _ := NekoSearch("Hellsing")

  for _, v := range u {
    if v.Neko_id == neko {
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
  fast, _, _ := c.FastIdAnime("initial d first stage")
  a, _ := fast.SearchAnimeVideos()

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
    s.Play(i)
    time.Sleep(1 * time.Second)
  }

  c := conf()
  fast, _, _ := c.FastIdUser("incarnati0n")
  um, _ := fast.UserUnreadMessages()

  if um.News > 0 || um.News == 0 {
    t.Logf("Found: %d news", um.News)
  } else {
    t.Error("News not found")
  }
}

func TestConstantsAnime(t *testing.T) {
  c := conf()
  ca, _, _ := c.SearchConstantsAnime()

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
  cm, _, _ := c.SearchConstantsManga()

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
    s.Play(i)
    time.Sleep(1 * time.Second)
  }

  c := conf()
  fast, _, _ := c.FastIdPeople("Aya Hirano")
  p, _ := fast.SearchPeople()

  if p.Id == 4 || p.Job_title == "Сэйю"  {
    t.Logf("%s - found", p.Name)
  } else {
    t.Error("People not found")
  }
}

func TestAnimeGraphql(t *testing.T) {
  c := conf()
  s, _ := graphql.AnimeSchema(
    graphql.Values("id", "malId", "name", "rating", "kind", "episodes"),
    "initial d first stage",
    1, 1, "", "", "", "", "", "", "", false,
  )
  a, _, _ := c.SearchGraphql(s)

  for _, v := range a.Data.Animes {
    if v.Id == "185" && v.MalId == "185" && v.Rating == "pg_13" && v.Kind == "tv" && v.Episodes == 26 {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("AnimeGraphql not found")
    }
  }
}

func TestMangaGraphQL(t *testing.T) {
  c := conf()
  s, _ := graphql.MangaSchema(
    graphql.Values("id", "malId", "name", "kind", "status", "volumes"),
    "initial d",
    1, 1, "", "", "", "", "", false,
  )
  m, _, _ := c.SearchGraphql(s)

  for _, v := range m.Data.Mangas {
    if v.Id == "375" && v.MalId == "375" && v.Kind == "manga" && v.Status == "released" && v.Volumes == 48 {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("MangaGraphql not found")
    }
  }
}

func TestCharacterGraphQL(t *testing.T) {
  c := conf()
  s, _ := graphql.CharacterSchema(
    graphql.Values("id", "malId", "name", "isManga"),
    "onizuka",
    1, 1,
  )
  ch, _, _ := c.SearchGraphql(s)

  for _, v := range ch.Data.Characters {
    if v.Id == "20847" && v.MalId == "20847" && v.IsManga {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("CharacterGraphql not found")
    }
  }
}
