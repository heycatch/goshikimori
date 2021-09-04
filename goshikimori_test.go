package goshikimori

import (
  "log"
  "testing"
  "math/rand"
  "time"
  "fmt"

  "github.com/vexilology/goshikimori/api"
)

const (
  test_method = "GET"
  test_app = ""
  test_access_token = ""
)

func random_number(min, max int) string {
  r := rand.Intn(max-min) + min
  return fmt.Sprintf("%v", r)
}

func TestRequest(t *testing.T) {
  rand.Seed(time.Now().UnixNano())

  req, _ := NewRequest(
    test_app,
    test_access_token,
    test_method,
    Parameters(api.Characters, api.FoundID(random_number(1, 5))),
  )

  if test_app != "" && test_access_token != "" {
    t.Log("Correct reuqest")
    log.Println(string(req))
  } else {
    t.Error("Bad request or problems with your app")
  }
}
