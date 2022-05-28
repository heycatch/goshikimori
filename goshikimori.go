package goshikimori

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strings"
  "encoding/json"
  "net/url"
  "strconv"

  "github.com/vexilology/goshikimori/api"
)

const (
  bearer   = "Bearer "
  urlOrig  = "%s://%s/%s"
  protocol = "https"
  urlShiki = "shikimori.one/api"
)

var ok bool

var client = &http.Client{}

type Configuration struct {
  Application, AccessToken string
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

type ExtraAnimeRates struct {
  Limit, Status, Censored string
}

type ExtraMangaRates struct {
  Limit, Censored string
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

type ResultAnimeRates interface {
  OptionsAnimeRates() string
}

type ResultMangaRates interface {
  OptionsMangaRates() string
}

func Add(app, tok string) *Configuration {
  return &Configuration{Application: app, AccessToken: tok}
}

func convertAchievements(id int) string {
  return fmt.Sprintf("achievements?user_id=%d", id)
}

func convertAnimeScreenshots(id int) string {
  return fmt.Sprintf("animes/%d/screenshots", id)
}

func convertAnimeVideos(id int) string {
  return fmt.Sprintf("animes/%d/videos", id)
}

func convertRoles(id int, name string) string {
  return fmt.Sprintf("%s/%d/roles", name, id)
}

func convertSimilar(id int, name string) string {
  return fmt.Sprintf("%s/%d/similar", name, id)
}

func convertRelated(id int, name string) string {
  return fmt.Sprintf("%s/%d/related", name, id)
}

func convertCalendar(name string) string {
  return fmt.Sprintf("calendar?%s", name)
}

func convertUserFriends(id int) string {
  return fmt.Sprintf("users/%d/friends", id)
}

func convertUserClubs(id int) string {
  return fmt.Sprintf("users/%d/clubs", id)
}

func convertUserRates(id int, name, options string) string {
  return fmt.Sprintf("users/%d/%s?%s", id, name, options)
}

// String formatting for achievements search
func NekoSearch(name string) string {
  r := strings.Replace(strings.ToLower(name), " ", "_", -1)
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
  if l == 0 { e.Limit = "1" }
  for i := 51; i <= l; i++ {
    e.Limit = "1"
  }

  kind_map := map[string]int{
    "tv": 1, "movie": 2, "ova": 3, "ona": 4,
    "special": 5, "music": 6,
    "tv_13": 7, "tv_24": 8, "tv_48": 9,
  }
  _, ok = kind_map[e.Kind]
  if !ok { e.Kind = "" }

  status_map := map[string]int{
    "anons": 1, "ongoing": 2, "released": 3,
  }
  _, ok = status_map[e.Status]
  if !ok { e.Status = "" }

  season_map := map[string]int{
    "summer_2017": 1, "2016": 2, "2014_2016": 3, "199x": 4,
  }
  _, ok = season_map[e.Season]
  if !ok { e.Status = "" }

  s, _ := strconv.Atoi(e.Score)
  for i := 10; i <= s; i++ {
    e.Score = ""
  }

  rating_map := map[string]int{
    "none": 1, "g": 2, "pg": 3, "pg_13": 4,
    "r": 5, "r_plus": 6, "rx": 7,
  }
  _, ok = rating_map[e.Rating]
  if !ok { e.Rating = "" }

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
  if l == 0 { e.Limit = "1" }
  for i := 51; i <= l; i++ {
    e.Limit = "1"
  }

  kind_map := map[string]int{
    "manga": 1, "manhwa": 2, "manhua": 3,
    "light_novel": 5, "novel": 6,
    "one_shot": 7, "doujin": 8,
  }
  _, ok = kind_map[e.Kind]
  if !ok { e.Kind = "" }

  status_map := map[string]int{
    "anons": 1, "ongoing": 2, "released": 3,
    "paused": 4, "discontinued": 5,
  }
  _, ok = status_map[e.Status]
  if !ok { e.Status = "" }

  season_map := map[string]int{
    "summer_2017": 1, "spring_2016,fall_2016": 2,
    "2016,!winter_2016": 3, "2016": 4,
    "2014_2016": 5, "199x": 6,
  }
  _, ok = season_map[e.Season]
  if !ok { e.Status = "" }

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
  if l == 0 { el.Limit = "1" }
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
  censored_map := map[string]int{"true": 1, "false": 2}
  _, ok = censored_map[ec.Censored]
  if !ok { ec.Censored = "false" }

  v := url.Values{}
  v.Add("censored", ec.Censored)

  return v.Encode()
}

// Limit    -> 5000 maximum
// Status   -> planned, watching, rewatching, completed, on_hold, dropped
// Censored -> true, false
// Set to true to discard hentai, yaoi and yuri
func (ar *ExtraAnimeRates) OptionsAnimeRates() string {
  l, _ := strconv.Atoi(ar.Limit)
  if l == 0 { ar.Limit = "1" }
  for i := 5001; i <= l; i++ {
    ar.Limit = "1"
  }

  status_map := map[string]int{
    "planned": 1, "watching": 2,
    "rewatching": 3, "completed": 4,
    "on_hold": 5, "dropped": 6,
  }
  _, ok = status_map[ar.Status]
  if !ok { ar.Status = "watching" }

  censored_map := map[string]int{"true": 1, "false": 2}
  _, ok = censored_map[ar.Censored]
  if !ok { ar.Censored = "false" }

  v := url.Values{}
  v.Add("limit", ar.Limit)
  v.Add("status", ar.Status)
  v.Add("censored", ar.Censored)

  return v.Encode()
}

// Limit    -> 5000 maximum
// Censored -> true, false
// Set to true to discard hentai, yaoi and yuri
func (mr *ExtraMangaRates) OptionsMangaRates() string {
  l, _ := strconv.Atoi(mr.Limit)
  if l == 0 { mr.Limit = "1" }
  for i := 5001; i <= l; i++ {
    mr.Limit = "1"
  }

  censored_map := map[string]int{"true": 1, "false": 2}
  _, ok = censored_map[mr.Censored]
  if !ok { mr.Censored = "false" }

  v := url.Values{}
  v.Add("limit", mr.Limit)
  v.Add("censored", mr.Censored)

  return v.Encode()
}

func (c *Configuration) NewGetRequest(search string) *http.Request {
  full_url := fmt.Sprintf(urlOrig, protocol, urlShiki, search)
  req, _ := http.NewRequest(http.MethodGet, full_url, nil)
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", bearer + c.AccessToken)
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

func (c *Configuration) SearchUserFriends(id int) ([]api.UserFriends, error) {
  var uf []api.UserFriends

  resp, err := client.Do(c.NewGetRequest(convertUserFriends(id)))
  if err != nil {
    return uf, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return uf, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return nil, err
  }

  return uf, nil
}

func (c *Configuration) SearchUserClubs(id int) ([]api.Clubs, error) {
  var uc []api.Clubs

  resp, err := client.Do(c.NewGetRequest(convertUserClubs(id)))
  if err != nil {
    return uc, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return uc, err
  }

  if err := json.Unmarshal(data, &uc); err != nil {
    return nil, err
  }

  return uc, nil
}

func (c *Configuration) SearchUserAnimeRates(id int, r ResultAnimeRates) ([]api.UserAnimeRates, error) {
  var ar []api.UserAnimeRates

  resp, err := client.Do(c.NewGetRequest(
    convertUserRates(id, "anime_rates", r.OptionsAnimeRates()),
  ))
  if err != nil {
    return ar, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return ar, err
  }

  if err := json.Unmarshal(data, &ar); err != nil {
    return nil, err
  }

  return ar, nil
}

func (c *Configuration) SearchUserMangaRates(id int, r ResultMangaRates) ([]api.UserMangaRates, error) {
  var mr []api.UserMangaRates

  resp, err := client.Do(c.NewGetRequest(
    convertUserRates(id, "manga_rates", r.OptionsMangaRates()),
  ))
  if err != nil {
    return mr, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return mr, err
  }

  if err := json.Unmarshal(data, &mr); err != nil {
    return nil, err
  }

  return mr, nil
}

func (c *Configuration) WhoAmi() (api.Who, error) {
  var w api.Who

  resp, err := client.Do(c.NewGetRequest("users/whoami"))
  if err != nil {
    return w, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return w, err
  }

  if err := json.Unmarshal(data, &w); err != nil {
    return w, err
  }

  return w, nil
}

func (c *Configuration) SearchAnime(name string, r Result) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(name) + "&" + r.OptionsAnime(),
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

func (c *Configuration) SearchManga(name string, r Result) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(name) + "&" + r.OptionsManga(),
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

func (c *Configuration) FastIdAnime(name string) (int, error) {
  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(name),
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

func (c *Configuration) FastIdManga(name string) (int, error) {
  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(name),
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

func (c *Configuration) SearchAnimeScreenshots(id int) ([]api.AnimeScreenshots, error) {
  var s []api.AnimeScreenshots

  resp, err := client.Do(c.NewGetRequest(convertAnimeScreenshots(id)))
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

func (c *Configuration) SearchSimilarAnime(id int) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(convertSimilar(id, "animes")))
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

func (c *Configuration) SearchSimilarManga(id int) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest(convertSimilar(id, "mangas")))
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

func (c *Configuration) SearchRelatedAnime(id int) ([]api.RelatedAnimes, error) {
  var a []api.RelatedAnimes

  resp, err := client.Do(c.NewGetRequest(convertRelated(id, "animes")))
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

func (c *Configuration) SearchRelatedManga(id int) ([]api.RelatedMangas, error) {
  var m []api.RelatedMangas

  resp, err := client.Do(c.NewGetRequest(convertRelated(id, "mangas")))
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

func (c *Configuration) SearchClub(name string, r ResultLimit) ([]api.Clubs, error) {
  var cl []api.Clubs

  resp, err := client.Do(
    c.NewGetRequest("clubs?search=" + url.QueryEscape(name) + "&" + r.OptionsClub()),
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
func (c *Configuration) SearchAchievement(id int) ([]api.Achievements, error) {
  var a []api.Achievements

  resp, err := client.Do(c.NewGetRequest(convertAchievements(id)))
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

func (c *Configuration) SearchAnimeVideos(id int) ([]api.AnimeVideos, error) {
  var v []api.AnimeVideos

  resp, err := client.Do(c.NewGetRequest(convertAnimeVideos(id)))
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

func (c *Configuration) SearchAnimeRoles(id int) ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(c.NewGetRequest(convertRoles(id, "animes")))
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

func (c *Configuration) SearchMangaRoles(id int) ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(c.NewGetRequest(convertRoles(id, "mangas")))
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
