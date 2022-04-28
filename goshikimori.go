package goshikimori

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "strings"
  "encoding/json"
  "net/url"

  "github.com/vexilology/goshikimori/api"
)

const (
  bearer   = "Bearer "
  urlOrig  = "%s://%s/%s"
  protocol = "https"
  urlShiki = "shikimori.one/api"
)

var client = &http.Client{}

type Configuration struct {
  Application string
  PrivateKey  string
}

func Add(app, key string) *Configuration {
  return &Configuration{Application: app, PrivateKey: key}
}

func convertAchievements(i int) string {
  return fmt.Sprintf("achievements?user_id=%d", i)
}

func convertAnimeScreenshots(i int) string {
  return fmt.Sprintf("animes/%d/screenshots", i)
}

func convertAnimeVideos(i int) string {
  return fmt.Sprintf("animes/%d/videos", i)
}

func convertSimilar(i int, s string) string {
  switch s {
  case "anime":
    return fmt.Sprintf("animes/%d/similar", i)
  case "manga":
    return fmt.Sprintf("mangas/%d/similar", i)
  case "ranobe":
    return fmt.Sprintf("ranobe/%d/similar", i)
  default:
    return ""
  }
}

func convertRelated(i int, s string) string {
  switch s {
  case "anime":
    return fmt.Sprintf("animes/%d/related", i)
  case "manga":
    return fmt.Sprintf("mangas/%d/related", i)
  case "ranobe":
    return fmt.Sprintf("ranobe/%d/related", i)
  default:
    return ""
  }
}

func NekoSearch(s string) string {
  r := strings.Replace(strings.ToLower(s), " ", "_", -1)
  return fmt.Sprintf("%s", r)
}

func (c *Configuration) NewGetRequest(f string) *http.Request {
  req, err := http.NewRequest(
    http.MethodGet,
    fmt.Sprintf(urlOrig, protocol, urlShiki, f),
    nil,
  )
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", bearer + c.PrivateKey)
  if err != nil {
    log.Fatal(err)
  }
  return req
}

func (c *Configuration) SearchUser(s string) api.Users {
  resp, err := client.Do(c.NewGetRequest(
    "users/" + url.QueryEscape(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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
  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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

func (c *Configuration) SearchAnimeScreenshots(i int) api.AnimeScreenshots {
  resp, err := client.Do(c.NewGetRequest(convertAnimeScreenshots(i)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var s []api.AnimeScreenshots
  var ss api.AnimeScreenshots

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &s); err != nil {
    log.Fatal(err)
  }

  for _, value := range s {
    ss = value
  }

  return ss
}

func (c *Configuration) SearchSimilarAnime(i int) api.Animes {
  resp, err := client.Do(c.NewGetRequest(
    convertSimilar(i, "anime")))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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

func (c *Configuration) SearchSimilarManga(i int) api.Mangas {
  resp, err := client.Do(c.NewGetRequest(
    convertSimilar(i, "manga")))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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

func (c *Configuration) SearchSimilarRanobe(i int) api.Mangas {
  resp, err := client.Do(c.NewGetRequest(
    convertSimilar(i, "ranobe")))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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

func (c *Configuration) SearchRelatedAnime(i int) api.RelatedAnimes {
  resp, err := client.Do(c.NewGetRequest(
    convertRelated(i, "anime")))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var a []api.RelatedAnimes
  var aa api.RelatedAnimes

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

func (c *Configuration) SearchRelatedManga(i int) api.RelatedMangas {
  resp, err := client.Do(c.NewGetRequest(
    convertRelated(i, "manga")))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var m []api.RelatedMangas
  var mm api.RelatedMangas

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

func (c *Configuration) SearchRelatedRanobe(i int) api.RelatedMangas {
  resp, err := client.Do(c.NewGetRequest(
    convertRelated(i, "ranobe")))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var m []api.RelatedMangas
  var mm api.RelatedMangas

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

func (c *Configuration) SearchManga(s string) api.Mangas {
  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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
  resp, err := client.Do(c.NewGetRequest(
    "ranobe?search=" + url.QueryEscape(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

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

func (c *Configuration) SearchClub(s string) api.Clubs {
  resp, err := client.Do(c.NewGetRequest(
    "clubs?search=" + url.QueryEscape(s)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var l []api.Clubs
  var ll api.Clubs

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &l); err != nil {
    log.Fatal(err)
  }

  for _, value := range l {
    ll = value
  }

  return ll
}

func (c *Configuration) SearchAchievement(i int) []api.Achievements {
  resp, err := client.Do(c.NewGetRequest(convertAchievements(i)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var a []api.Achievements

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &a); err != nil {
    log.Fatal(err)
  }

  return a
}

func (c *Configuration) SearchAnimeVideos(i int) api.AnimeVideos {
  resp, err := client.Do(c.NewGetRequest(convertAnimeVideos(i)))
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var v []api.AnimeVideos
  var vv api.AnimeVideos

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  if json.Unmarshal(data, &v); err != nil {
    log.Fatal(err)
  }

  for _, value := range v {
    vv = value
  }

  return vv
}
