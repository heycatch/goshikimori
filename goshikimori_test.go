package goshikimori

import (
  "fmt"
  "os"
  "testing"

  graph "github.com/heycatch/goshikimori/graphql"
)

const (
  app_test = ""
  tok_test = ""
)

func conf() *Configuration { return SetConfiguration(app_test, tok_test) }

func TestConfiguration(t *testing.T) {
  if app_test != "" && tok_test != "" {
    t.Logf("Found: %s and %s", app_test, tok_test)
  } else {
    t.Error("Not found application or token")
    os.Exit(1)
  }
}

func TestUser(t *testing.T) {
  c := conf()
  s, _, _:= c.SearchUser("arctica")

  if s.Id == 181833 && s.Sex == "male" {
    t.Logf("User: %s, Id: %d - found", s.Nickname, s.Id)
  } else {
    t.Error("User: arctica - not found")
  }
}

func TestAnimes(t *testing.T) {
  c := conf()
  o := &Options{
    Page: 1, Limit: 1, Score: 1,
    Censored: false,
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
  o := &Options{Page: 1, Limit: 1}
  r, _, _ := c.SearchMangas("Initial D", o)

  for _, v := range r {
    if v.Volumes == 48 && v.Chapters == 724 {
      t.Logf("Manga: %s, Id: %d - found", v.Name, v.Id)
    } else {
      t.Errorf("Manga: %s, Id: %d - not found", v.Name, v.Id)
    }
  }
}

func TestClubs(t *testing.T) {
  c := conf()
  o := &Options{Page: 1, Limit: 1}
  clubs, _, _ := c.SearchClubs("milf", o)

  for _, v := range clubs {
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
  s.settings(5, "#", 1)
  s.run()

  c := conf()
  fast, _, _ := c.FastIdUser("arctica")
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
  s.settings(5, "#", 1)
  s.run()

  c := conf()
  fast, _, _ := c.FastIdUser("arctica")
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

func TestAnimeGraphql(t *testing.T) {
  fmt.Println("Too many requests at once, waiting 5 seconds...")

  var s StatusBar
  s.settings(5, "#", 1)
  s.run()

  c := conf()
  sch, _ := graph.AnimeSchema(
    graph.Values("id", "malId", "name", "rating", "kind", "episodes"),
    "initial d first stage", 1, 1, "", "", "", "", "", "", "", false, nil,
  )
  a, _, _ := c.SearchGraphql(sch)

  for _, v := range a.Data.Animes {
    if v.Id == "185" && v.MalId == "185" && v.Rating == ANIME_RATING_PG_13 &&
      v.Kind == ANIME_KIND_TV && v.Episodes == 26 {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("AnimeGraphql not found")
    }
  }
}

func TestMangaGraphQL(t *testing.T) {
  c := conf()
  s, _ := graph.MangaSchema(
    graph.Values("id", "malId", "name", "kind", "status", "volumes"),
    "initial d", 1, 1, "", "", "", "", "", false, nil,
  )
  m, _, _ := c.SearchGraphql(s)

  for _, v := range m.Data.Mangas {
    if v.Id == "375" && v.MalId == "375" && v.Kind == MANGA_KIND_MANGA &&
      v.Status == MANGA_STATUS_RELEASED && v.Volumes == 48 {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("MangaGraphql not found")
    }
  }
}

func TestCharacterGraphQL(t *testing.T) {
  c := conf()
  s, _ := graph.CharacterSchema(
    graph.Values("id", "malId", "name", "isManga"),
    "Natsuno Yuuki", 1, 1,
  )
  ch, _, _ := c.SearchGraphql(s)

  for _, v := range ch.Data.Characters {
    if v.Id == "7582" && v.MalId == "7582" && v.IsManga {
      t.Logf("%s - found", v.Name)
    } else {
      t.Error("CharacterGraphql not found")
    }
  }
}

func TestPeopleGraphQL(t *testing.T) {
  c := conf()
  s, _ := graph.PeopleSchema(
    graph.Values("id", "name", "birthOn{year}"),
    "satsuki", 1, 1, true, false, false,
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
  o := &Options{Page: 1, Limit: 1, Genre_v2: []int{3}}
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
  o := &Options{Page: 1, Limit: 1, Genre_v2: []int{84}}
  r, _, _ := c.SearchMangas("Initial D", o)

  for _, v := range r {
    if v.Volumes == 48 && v.Chapters == 724 {
      t.Logf("Manga: %s, Id: %d - found", v.Name, v.Id)
    } else {
      t.Errorf("Manga: %s, Id: %d - not found", v.Name, v.Id)
    }
  }
}
