package goshikimori

import (
  "testing"
  "math/rand"
  "time"
  "log"
  "encoding/json"

  "github.com/vexilology/goshikimori/api"
  "github.com/vexilology/goshikimori/req"
)

const (
  api_test = ""
  key_test = ""
)

func returnConf() *req.Config {
  return &req.Config{
    api_test,
    key_test,
  }
}

func random_number(min, max int) int { return rand.Intn(max-min)+min }

type Grep struct {
  Nickname string `json:"nickname"`
  Id       int    `json:"id"`
  Name     string `json:"name"`
}

func TestUser(t *testing.T) {
  conf := returnConf()

  u, _ := NewRequest(
    conf.Application, conf.SecretKey, req.Get,
    Parameters(api.Users, api.Id(181833)),
  )

  var G Grep
  err := json.Unmarshal(u, &G)
  if err != nil {
    log.Fatal(err)
  }

  if api_test != "" && key_test != "" {
    if G.Nickname == "incarnati0n" {
      t.Logf("User %s id %v found", G.Nickname, G.Id)
    } else {
      t.Errorf("User %s id %v not found", G.Nickname, G.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}

func TestAnimes(t *testing.T) {
  rand.Seed(time.Now().UnixNano())
  conf := returnConf()

  a, _ := NewRequest(
    conf.Application, conf.SecretKey, req.Get,
    Parameters(api.Characters, api.Id(random_number(1, 4))),
  )

  var G Grep
  err := json.Unmarshal(a, &G)
  if err != nil {
    log.Fatal(err)
  }

  if api_test != "" && key_test != "" {
    if G.Id == 1 || G.Id == 2 || G.Id == 3 || G.Id == 4 {
      t.Logf("Anime %s id %v found", G.Name, G.Id)
    } else {
      t.Errorf("Anime %s id %v not found", G.Name, G.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}

func TestMangas(t *testing.T) {
  rand.Seed(time.Now().UnixNano())
  conf := returnConf()

  m, _ := NewRequest(
    conf.Application, conf.SecretKey, req.Get,
    Parameters(api.Mangas, api.Id(random_number(1, 4))),
  )

  var G Grep
  err := json.Unmarshal(m, &G)
  if err != nil {
    log.Fatal(err)
  }

  if api_test != "" && key_test != "" {
    if G.Id == 1 || G.Id == 2 || G.Id == 3 || G.Id == 4 {
      t.Logf("Manga %s id %v found", G.Name, G.Id)
    } else {
      t.Errorf("Manga %s id %v not found", G.Name, G.Id)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}
