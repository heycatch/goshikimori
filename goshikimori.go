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

func checkStatus(s string) bool {
  if s == "200 OK" {
    return true
  } else {
    return false
  }
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

func (c *Configuration) SearchUser(s string) api.Users {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertUser(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  ok := checkStatus(resp.Status); if ok != true {
    log.Fatal("request failed, check your app or private key")
  }

  var u api.Users

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &u); err != nil {
    log.Fatal(err)
  }

  return u
}

func (c *Configuration) SearchAnime(s string) api.Animes {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertAnime(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  ok := checkStatus(resp.Status); if ok != true {
    log.Fatal("request failed, check your app or private key")
  }

  var a []api.Animes
  var aa api.Animes

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &a); err != nil {
    log.Fatal(err)
  }

  for _, value := range a {
    aa = value
  }

  return aa
}

func (c *Configuration) SearchManga(s string) api.Mangas {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertManga(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  ok := checkStatus(resp.Status); if ok != true {
    log.Fatal("request failed, check your app or private key")
  }

  var m []api.Mangas
  var mm api.Mangas

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &m); err != nil {
    log.Fatal(err)
  }

  for _, value := range m {
    mm = value
  }

  return mm
}

func (c *Configuration) SearchRanobe(s string) api.Mangas {
  client := &http.Client{}
  resp, err := client.Do(c.NewGetRequest(convertRanobe(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  ok := checkStatus(resp.Status); if ok != true {
    log.Fatal("request failed, check your app or private key")
  }

  var r []api.Mangas
  var rr api.Mangas

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &r); err != nil {
    log.Fatal(err)
  }

  for _, value := range r {
    rr = value
  }

  return rr
}
