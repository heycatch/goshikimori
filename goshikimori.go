package main

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "strings"

  "goshikimori/api"
)

const (
  oauth_app    = ""
  access_token = ""
  bearer       = "Bearer "
  urlOrig      = "%s://%s"
  protocol     = "https"
  urlShiki     = "shikimori.one/api/"
)

func Parameters(s ...string) string {
  p := strings.Join(s, "/")
  return p
}

func NewRequest(method, input string) ([]byte, error) {
  req, err := http.NewRequest(
    method, fmt.Sprintf(urlOrig, protocol, urlShiki)+input, nil)
  req.Header.Add("User-Agent", oauth_app)
  req.Header.Add("Authorization", bearer+access_token)
  if err != nil {
    log.Fatal(err)
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  return ioutil.ReadAll(resp.Body)
}

func main() {
  result, err := NewRequest(
    "GET",
    Parameters(api.Animes, api.FoundID("24"), api.Similar),
  )
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(result))
}
