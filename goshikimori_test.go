package goshikimori

import (
  "testing"
  "log"
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

func TestUser(t *testing.T) {
  c := conf()
  r, err := c.SearchUser("incarnati0n")
  if err != nil {
    log.Fatal(err)
  }
  if api_test != "" && key_test != "" {
    if r.Id == 181833 && r.Sex == "male" {
      t.Logf("User %s id %d - found", r.Nickname, r.Id)
    } else {
      t.Errorf("User %s id %d - not found", r.Nickname, r.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}

func TestAnimes(t *testing.T) {
  c := conf()
  r, err := c.SearchAnime("Initial D")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    if api_test != "" && key_test != "" {
      if values.Id == 12725 && values.Status == "released" {
        t.Logf("Anime %s id %d - found", values.Name, values.Id)
      } else {
        t.Errorf("Anime %s id %d - not found", values.Name, values.Id)
      }
    } else {
      t.Error("Not found Application or SecretKey")
    }
  }
}

func TestMangas(t *testing.T) {
  c := conf()
  r, err := c.SearchManga("Initial D")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    if api_test != "" && key_test != "" {
      if values.Volumes == 48 && values.Chapters == 724 {
        t.Logf("Manga %s id %d - found", values.Name, values.Id)
      } else {
        t.Errorf("Manga %s id %d - not found", values.Name, values.Id)
      }
    } else {
      t.Error("Not found Application or SecretKey")
    }
  }
}

func TestRanobe(t *testing.T) {
  c := conf()
  r, err := c.SearchRanobe("Vampire Knight")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    if api_test != "" && key_test != "" {
      if values.Volumes == 1 && values.Chapters == 6 {
        t.Logf("Ranobe %s id %d - found", values.Name, values.Id)
      } else {
        t.Errorf("Ranobe %s id %d - not found", values.Name, values.Id)
      }
    } else {
      t.Error("Not found Application or SecretKey")
    }
  }
}
