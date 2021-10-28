package goshikimori

import (
  "fmt"
  "log"
  "net/http"
  "io"
  "io/ioutil"
  "strings"
  "encoding/json"

  "github.com/vexilology/goshikimori/api"
)

const (
  bearer   = "Bearer "
  urlOrig  = "%s://%s/%s"
  protocol = "https"
  urlShiki = "shikimori.one/api"
)

type Configuration struct {
  Application string
  PrivateKey  string
}

func Add(app, key string) *Configuration {
  return &Configuration{Application: app, PrivateKey: key}
}

func convertUser(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("users/%s", c)
}

func convertAnime(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("animes?search=%s", c)
}

func convertManga(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("mangas?search=%s", c)
}

func convertRanobe(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("ranobe?search=%s", c)
}

func decodeJSON(r io.Reader, i interface{}) error {
  data, err := ioutil.ReadAll(r)
  if err != nil {
    log.Fatal(err)
  }
  return json.Unmarshal(data, i)
}

func (c *Configuration) NewGetRequest(f string) *http.Request {
  req, err := http.NewRequest(
    http.MethodGet,
    fmt.Sprintf(urlOrig, protocol, urlShiki, f), nil,
  )
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", bearer + c.PrivateKey)
  if err != nil {
    log.Fatal(err)
  }
  return req
}

func (c *Configuration) SearchUser(s string) (api.Users, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertUser(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var u api.Users
  return u, decodeJSON(resp.Body, &u)
}

func (c *Configuration) SearchAnime(s string) ([]api.Animes, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertAnime(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var a []api.Animes
  return a, decodeJSON(resp.Body, &a)
}

func (c *Configuration) SearchManga(s string) ([]api.Mangas, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertManga(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var m []api.Mangas
  return m, decodeJSON(resp.Body, &m)
}

func (c *Configuration) SearchRanobe(s string) ([]api.Mangas, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertRanobe(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var r []api.Mangas
  return r, decodeJSON(resp.Body, &r)
}
