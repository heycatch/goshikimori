package goshikimori

import "testing"

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

func TestUser(t *testing.T) {
  c := conf()
  s := c.SearchUser("incarnati0n")
  if api_test != "" && key_test != "" {
    if s.Id == 181833 && s.Sex == "male" {
      t.Logf("User %s id %d - found", s.Nickname, s.Id)
    } else {
      t.Errorf("User %s id %d - not found", s.Nickname, s.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}

func TestAnimes(t *testing.T) {
  c := conf()
  s := c.SearchAnime("Initial D")
  if api_test != "" && key_test != "" {
    if s.Id == 12725 && s.Status == "released" {
      t.Logf("Anime %s id %d - found", s.Name, s.Id)
    } else {
      t.Errorf("Anime %s id %d - not found", s.Name, s.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}

func TestMangas(t *testing.T) {
  c := conf()
  r := c.SearchManga("Initial D")
  if api_test != "" && key_test != "" {
    if r.Volumes == 48 && r.Chapters == 724 {
      t.Logf("Manga %s id %d - found", r.Name, r.Id)
    } else {
      t.Errorf("Manga %s id %d - not found", r.Name, r.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}

func TestRanobe(t *testing.T) {
  c := conf()
  r := c.SearchRanobe("Vampire Knight")
  if api_test != "" && key_test != "" {
    if r.Volumes == 1 && r.Chapters == 6 {
      t.Logf("Ranobe %s id %d - found", r.Name, r.Id)
    } else {
      t.Errorf("Ranobe %s id %d - not found", r.Name, r.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}
