package goshikimori

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strings"
  "encoding/json"
  "net/url"
  "strconv"
  "time"

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
  Application, PrivateKey string
}

type Extra struct {
  Limit, Kind, Status, Season, Score, Rating string
}

type ExtraLimit struct {
  Limit string
}

type ExtraCensored struct {
  Censored string
}

type Result interface {
  OptionsAnime() string
  OptionsManga() string
}

type ResultLimit interface {
  OptionsClub() string
}

type ResultCensored interface {
  OptionsCalendar() string
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

func convertRoles(i int, s string) string {
  return fmt.Sprintf("%s/%d/roles", s, i)
}

func convertSimilar(i int, s string) string {
  return fmt.Sprintf("%s/%d/similar", s, i)
}

func convertRelated(i int, s string) string {
  return fmt.Sprintf("%s/%d/related", s, i)
}

func convertCalendar(s string) string {
  return fmt.Sprintf("calendar?%s", s)
}

// String formatting for achievements search
func NekoSearch(s string) string {
  r := strings.Replace(strings.ToLower(s), " ", "_", -1)
  return fmt.Sprintf("%s", r)
}

// Limit  -> 50 maximum
// Kind   -> tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48
// Status -> anons, ongoing, released
// Season -> summer_2017, 2016, 2014_2016, 199x
// Score  -> 9 maximum
// Rating -> none, g, pg, pg_13, r, r_plus, rx
func (e *Extra) OptionsAnime() string {
  l, _ := strconv.Atoi(e.Limit)
  for i := 51; i <= l; i++ {
    e.Limit = "1"
  }

  var ok bool

  kind_map := map[string]int{
    "tv": 1, "movie": 2, "ova": 3, "ona": 4,
    "special": 5, "music": 6,
    "tv_13": 7, "tv_24": 8, "tv_48": 9,
  }
  _, ok = kind_map[e.Kind]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Kind = ""
  }

  status_map := map[string]int{
    "anons": 1, "ongoing": 2, "released": 3,
  }
  _, ok = status_map[e.Status]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Status = ""
  }

  season_map := map[string]int{
    "summer_2017": 1, "2016": 2, "2014_2016": 3, "199x": 4,
  }
  _, ok = season_map[e.Season]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Status = ""
  }

  s, _ := strconv.Atoi(e.Score)
  for i := 10; i <= s; i++ {
    e.Score = ""
  }

  rating_map := map[string]int{
    "none": 1, "g": 2, "pg": 3, "pg_13": 4,
    "r": 5, "r_plus": 6, "rx": 7,
  }
  _, ok = rating_map[e.Rating]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Rating = ""
  }

  v := url.Values{}
  v.Add("limit", e.Limit)
  v.Add("kind", e.Kind)
  v.Add("status", e.Status)
  v.Add("season", e.Season)
  v.Add("score", e.Score)
  v.Add("rating", e.Rating)

  return v.Encode()
}

// Limit  -> 50 maximum
// Kind   -> manga, manhwa, manhua,
//           light_novel, novel, one_shot, doujin
// Status -> anons, ongoing, released, paused, discontinued
// Season -> summer_2017, "spring_2016,fall_2016",
//           "2016,!winter_2016", 2016, 2014_2016, 199x
// Score  -> 9 maximum
func (e *Extra) OptionsManga() string {
  l, _ := strconv.Atoi(e.Limit)
  for i := 51; i <= l; i++ {
    e.Limit = "1"
  }

  var ok bool

  kind_map := map[string]int{
    "manga": 1, "manhwa": 2, "manhua": 3,
    "light_novel": 5, "novel": 6,
    "one_shot": 7, "doujin": 8,
  }
  _, ok = kind_map[e.Kind]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Kind = ""
  }

  status_map := map[string]int{
    "anons": 1, "ongoing": 2, "released": 3,
    "paused": 4, "discontinued": 5,
  }
  _, ok = status_map[e.Status]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Status = ""
  }

  season_map := map[string]int{
    "summer_2017": 1, "spring_2016,fall_2016": 2,
    "2016,!winter_2016": 3, "2016": 4,
    "2014_2016": 5, "199x": 6,
  }
  _, ok = season_map[e.Season]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    e.Status = ""
  }

  s, _ := strconv.Atoi(e.Score)
  for i := 10; i <= s; i++ {
    e.Score = ""
  }

  v := url.Values{}
  v.Add("limit", e.Limit)
  v.Add("kind", e.Kind)
  v.Add("status", e.Status)
  v.Add("season", e.Season)
  v.Add("score", e.Score)

  return v.Encode()
}

// Limit -> 30 maximum
func (el *ExtraLimit) OptionsClub() string {
  l, _ := strconv.Atoi(el.Limit)
  for i := 31; i <= l; i++ {
    el.Limit = "1"
  }

  v := url.Values{}
  v.Add("limit", el.Limit)

  return v.Encode()
}

// Censored -> true, false
// Set to false to allow hentai, yaoi and yuri
func (ec *ExtraCensored) OptionsCalendar() string {
  var ok bool

  censored_map := map[string]int{"true": 1, "false": 2}
  _, ok = censored_map[ec.Censored]
  if ok {
    time.Sleep(100 * time.Millisecond)
  } else {
    ec.Censored = "false"
  }

  v := url.Values{}
  v.Add("censored", ec.Censored)

  return v.Encode()
}

func (c *Configuration) NewGetRequest(search string) *http.Request {
  url := fmt.Sprintf(urlOrig, protocol, urlShiki, search)
  req, _ := http.NewRequest(http.MethodGet, url, nil)
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", bearer + c.PrivateKey)
  return req
}

// NOTES: search by user is case sensitive
func (c *Configuration) SearchUser(name string) (api.Users, error) {
  var u api.Users

  resp, err := client.Do(
    c.NewGetRequest("users/" + url.QueryEscape(name)),
  )
  if err != nil {
    return u, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return u, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return u, err
  }

  return u, nil
}

func (c *Configuration) SearchAnime(s string, r Result) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(s) + "&" + r.OptionsAnime(),
  ))
  if err != nil {
    return a, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return a, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchManga(s string, r Result) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(s) + "&" + r.OptionsManga(),
  ))
  if err != nil {
    return m, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return m, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

func (c *Configuration) FastIdAnime(s string) (int, error) {
  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(s),
  ))
  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()

  var a []api.Animes
  var aa api.Animes

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return 0, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return 0, err
  }

  for _, v := range a {
    aa = v
  }

  return aa.Id, nil
}

func (c *Configuration) FastIdManga(s string) (int, error) {
  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(s),
  ))
  if err != nil {
    return 0, err
  }

  var m []api.Mangas
  var mm api.Mangas

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return 0, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return 0, err
  }

  for _, v := range m {
    mm = v
  }

  return mm.Id, nil
}

func (c *Configuration) SearchAnimeScreenshots(i int) ([]api.AnimeScreenshots, error) {
  var s []api.AnimeScreenshots

  resp, err := client.Do(c.NewGetRequest(convertAnimeScreenshots(i)))
  if err != nil {
    return s, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return s, err
  }

  if err := json.Unmarshal(data, &s); err != nil {
    return nil, err
  }

  return s, nil
}

func (c *Configuration) SearchSimilarAnime(i int) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(convertSimilar(i, "animes")))
  if err != nil {
    return a, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return a, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchSimilarManga(i int) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest(convertSimilar(i, "mangas")))
  if err != nil {
    return m, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return m, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

func (c *Configuration) SearchRelatedAnime(i int) ([]api.RelatedAnimes, error) {
  var a []api.RelatedAnimes

  resp, err := client.Do(c.NewGetRequest(convertRelated(i, "animes")))
  if err != nil {
    return a, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return a, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchRelatedManga(i int) ([]api.RelatedMangas, error) {
  var m []api.RelatedMangas

  resp, err := client.Do(c.NewGetRequest(convertRelated(i, "mangas")))
  if err != nil {
    return m, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return m, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

func (c *Configuration) SearchClub(s string, r ResultLimit) ([]api.Clubs, error) {
  var cl []api.Clubs

  resp, err := client.Do(
    c.NewGetRequest("clubs?search=" + url.QueryEscape(s) + "&" + r.OptionsClub()),
  )
  if err != nil {
    return cl, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return cl, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return nil, err
  }

  return cl, nil
}

// NOTES: as a result, we return a complete list of all achievements.
// Next comes the filtering through "NekoSearch" and the error about obtaining
// specific achievements is already being processed there.
// See example in README.md
func (c *Configuration) SearchAchievement(i int) ([]api.Achievements, error) {
  var a []api.Achievements

  resp, err := client.Do(c.NewGetRequest(convertAchievements(i)))
  if err != nil {
    return a, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return a, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchAnimeVideos(i int) ([]api.AnimeVideos, error) {
  var v []api.AnimeVideos

  resp, err := client.Do(c.NewGetRequest(convertAnimeVideos(i)))
  if err != nil {
    return v, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return v, err
  }

  if err := json.Unmarshal(data, &v); err != nil {
    return nil, err
  }

  return v, nil
}

func (c *Configuration) SearchAnimeRoles(i int) ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(c.NewGetRequest(convertRoles(i, "animes")))
  if err != nil {
    return r, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return r, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, err
  }

  return r, nil
}

func (c *Configuration) SearchMangaRoles(i int) ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(c.NewGetRequest(convertRoles(i, "mangas")))
  if err != nil {
    return r, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return r, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, err
  }

  return r, nil
}

func (c *Configuration) SearchBans() ([]api.Bans, error) {
  var b []api.Bans

  resp, err := client.Do(c.NewGetRequest("bans"))
  if err != nil {
    return b, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return b, err
  }

  if err := json.Unmarshal(data, &b); err != nil {
    return nil, err
  }

  return b, nil
}

func (c *Configuration) SearchCalendar(r ResultCensored) ([]api.Calendar, error) {
  var ca []api.Calendar

  resp, err := client.Do(c.NewGetRequest(convertCalendar(r.OptionsCalendar())))
  if err != nil {
    return ca, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return ca, err
  }

  if err := json.Unmarshal(data, &ca); err != nil {
    return nil, err
  }

  return ca, nil
}
