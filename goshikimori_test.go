package goshikimori

import (
  "log"
  "testing"

  "github.com/vexilology/goshikimori/api"
)

const (
  test_method = "GET"
  test_app = ""
  test_access_token = ""
)

func TestRequest(t *testing.T) {
  req, _ := NewRequest(
    test_app,
    test_access_token,
    test_method,
    Parameters(api.Users, api.Whoami),
  )

  if test_method != "" && test_app != "" && test_access_token != "" {
    t.Log("Correct reuqest")
    log.Println(string(req))
  } else {
    t.Error("Bad method or app or token")
  }
}
