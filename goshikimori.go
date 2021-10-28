package goshikimori

import (
  "fmt"
  "log"
  "net/http"
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

func ConvertUser(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("users/%s", c)
}

func ConvertAnime(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("animes?search=%s", c)
}

func ConvertManga(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("mangas?search=%s", c)
}

func ConvertRanobe(s string) string {
  c := strings.Replace(s, " ", "+", -1)
  return fmt.Sprintf("ranobe?search=%s", c)
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
  resp, err := client.Do(c.NewGetRequest(ConvertUser(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  var u api.Users
  return u, json.Unmarshal(data, &u)
}

func (c *Configuration) SearchAnime(s string) ([]api.Animes, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(ConvertAnime(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  var a []api.Animes
  return a, json.Unmarshal(data, &a)
}

func (c *Configuration) SearchManga(s string) ([]api.Mangas, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(ConvertManga(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  var m []api.Mangas
  return m, json.Unmarshal(data, &m)
}

func (c *Configuration) SearchRanobe(s string) ([]api.Mangas, error) {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(ConvertRanobe(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  var m []api.Mangas
  return m, json.Unmarshal(data, &m)
}
