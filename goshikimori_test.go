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
    Censored: "", Mylist: "", Genre_v2: nil,
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
    Censored: "", Mylist: "", Genre_v2: nil,
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
    1, 1, "", "", "", "", "", "", "", false, nil,
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
    1, 1, "", "", "", "", "", false, nil,
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
    if v.Id == "17245" && v.MalId == "17245" && v.IsManga {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("CharacterGraphql not found")
    }
  }
}

func TestPeopleGraphQL(t *testing.T) {
  c := conf()
  s, _ := graphql.PeopleSchema(
    graphql.Values("id", "name", "birthOn{year}"),
    "satsuki",
    1, 1, true, false, false,
  )
  p, _, _ := c.SearchGraphql(s)

  for _, v := range p.Data.People {
    if v.Id == "3" && v.BirthOn.Year == 1970 {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("PeopleGraphql not found")
    }
  }
}

func TestAnimesUsingGenre(t *testing.T) {
  c := conf()
  o := &Options{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "", Duration: "",
    Censored: "", Mylist: "", Genre_v2: []int{3},
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

func TestMangasUsingGenre(t *testing.T) {
  c := conf()
  o := &Options{
    Page: "1", Limit: "1", Kind: "", Status: "",
    Season: "", Score: "", Rating: "",
    Censored: "", Mylist: "", Genre_v2: []int{84},
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
