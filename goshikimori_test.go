package goshikimori

import (
  "testing"
  "math/rand"
  "time"

  "github.com/vexilology/goshikimori/api"
  "github.com/vexilology/goshikimori/req"
)

const (
  api_test = "Application"
  key_test = "SecretKey"
)

func returnConf() *req.Config {
  return &req.Config{
    api_test,
    key_test,
  }
}

func random_number(min, max int) int { return rand.Intn(max-min)+min }

func TestRequest(t *testing.T) {
  rand.Seed(time.Now().UnixNano())
  conf := returnConf()

  r, _ := NewRequest(
    conf.Application, conf.SecretKey, req.Get,
    Parameters(api.Characters, api.Id(random_number(1, 5))),
  )

  if conf.Application != "" && conf.SecretKey != "" {
    if r != nil {
      t.Logf("%s\n", r)
    } else {
      t.Errorf("%s\n", r)
    }
  } else {
    t.Error("Not found Application or SecretKey")
  }
}
