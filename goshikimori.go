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
  "github.com/vexilology/goshikimori/transform"
)

const (
  bearer   = "Bearer "
  protocol = "https"
  urlshiki = "shikimori.one/api"
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

type ExtraTargetType struct {
  Limit, Target_type string
}

type Result interface {
  OptionsAnime() string
  OptionsManga() string
}

type ResultLimit interface {
  OptionsUsers() string
  OptionsClub()  string
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

type ResultUserHistory interface {
  OptionsUserHistory() string
}

func Add(app, tok string) *Configuration {
  return &Configuration{Application: app, AccessToken: tok}
}

// string formatting for achievements search
func NekoSearch(name string) string {
  r := strings.Replace(strings.ToLower(name), " ", "_", -1)
  return fmt.Sprintf("%s", r)
}

// Limit       -> 100
// Target_type -> Anime, Manga
func (ett *ExtraTargetType) OptionsUserHistory() string {
  l, _ := strconv.Atoi(ett.Limit)
  if l == 0 { ett.Limit = "1" }
  for i := 101; i <= l; i++ {
    ett.Limit = "1"
  }

  target_map := map[string]int8{"Anime": 1, "Manga": 2}
  _, ok = target_map[ett.Target_type]
  if !ok { ett.Target_type = "Anime" }

  v := url.Values{}
  v.Add("limit", ett.Limit)
  v.Add("target_type", ett.Target_type)

  return v.Encode()
}

// Limit -> 100
func (el *ExtraLimit) OptionsUsers() string {
  l, _ := strconv.Atoi(el.Limit)
  if l == 0 { el.Limit = "1" }
  for i := 101; i <= l; i++ {
    el.Limit = "1"
  }

  v := url.Values{}
  v.Add("limit", el.Limit)

  return v.Encode()
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

  kind_map := map[string]int8{
    "tv": 1, "movie": 2, "ova": 3, "ona": 4,
    "special": 5, "music": 6,
    "tv_13": 7, "tv_24": 8, "tv_48": 9,
  }
  _, ok = kind_map[e.Kind]
  if !ok { e.Kind = "" }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3,
  }
  _, ok = status_map[e.Status]
  if !ok { e.Status = "" }

  season_map := map[string]int8{
    "summer_2017": 1, "2016": 2, "2014_2016": 3, "199x": 4,
  }
  _, ok = season_map[e.Season]
  if !ok { e.Status = "" }

  s, _ := strconv.Atoi(e.Score)
  for i := 10; i <= s; i++ {
    e.Score = ""
  }

  rating_map := map[string]int8{
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

  kind_map := map[string]int8{
    "manga": 1, "manhwa": 2, "manhua": 3,
    "light_novel": 5, "novel": 6,
    "one_shot": 7, "doujin": 8,
  }
  _, ok = kind_map[e.Kind]
  if !ok { e.Kind = "" }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3,
    "paused": 4, "discontinued": 5,
  }
  _, ok = status_map[e.Status]
  if !ok { e.Status = "" }

  season_map := map[string]int8{
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
  censored_map := map[string]int8{"true": 1, "false": 2}
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

  status_map := map[string]int8{
    "planned": 1, "watching": 2,
    "rewatching": 3, "completed": 4,
    "on_hold": 5, "dropped": 6,
  }
  _, ok = status_map[ar.Status]
  if !ok { ar.Status = "watching" }

  censored_map := map[string]int8{"true": 1, "false": 2}
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

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok = censored_map[mr.Censored]
  if !ok { mr.Censored = "false" }

  v := url.Values{}
  v.Add("limit", mr.Limit)
  v.Add("censored", mr.Censored)

  return v.Encode()
}

func (c *Configuration) NewGetRequest(search string) *http.Request {
  custom_url := fmt.Sprintf("%s://%s/%s", protocol, urlshiki, search)
  req, _ := http.NewRequest(http.MethodGet, custom_url, nil)
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", bearer + c.AccessToken)
  return req
}

func (c *Configuration) NewPostRequest(search string) *http.Request {
  custom_url := fmt.Sprintf("%s://%s/%s", protocol, urlshiki, search)
  data := url.Values{} // empty data
  req, _ := http.NewRequest(http.MethodPost, custom_url, strings.NewReader(data.Encode()))
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", bearer + c.AccessToken)
  return req
}

func (c *Configuration) NewDeleteRequest(search string) *http.Request {
  custom_url := fmt.Sprintf("%s://%s/%s", protocol, urlshiki, search)
  data := url.Values{} // empty data
  req, _ := http.NewRequest(http.MethodDelete, custom_url, strings.NewReader(data.Encode()))
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

// don't use Stats.Statuses.Anime and Stats.Statuses.Manga: empty slice
func (c *Configuration) SearchUsers(name string, r ResultLimit) ([]api.Users, error) {
  var u []api.Users

  resp, err := client.Do(
    c.NewGetRequest("users?search=" + url.QueryEscape(name) + "&" + r.OptionsUsers()),
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
    return nil, err
  }

  return u, nil
}

func (c *Configuration) SearchUserFriends(id int) ([]api.UserFriends, error) {
  var uf []api.UserFriends

  resp, err := client.Do(c.NewGetRequest(transform.ConvertUser(id, "friends")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertUser(id, "clubs")))
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
    transform.ConvertUserRates(id, "anime_rates", r.OptionsAnimeRates()),
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
    transform.ConvertUserRates(id, "manga_rates", r.OptionsMangaRates()),
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

func (c *Configuration) SearchUserFavourites(id int) (api.UserFavourites, error) {
  var uf api.UserFavourites

  resp, err := client.Do(c.NewGetRequest(transform.ConvertUser(id, "favourites")))
  if err != nil {
    return uf, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return uf, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return uf, err
  }

  return uf, nil
}

func (c *Configuration) SearchUserHistory(id int, r ResultUserHistory) ([]api.UserHistory, error) {
  var uh []api.UserHistory

  resp, err := client.Do(c.NewGetRequest(
    transform.ConvertUserRates(id, "history", r.OptionsUserHistory()),
  ))
  if err != nil {
    return uh, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return uh, err
  }

  if err := json.Unmarshal(data, &uh); err != nil {
    return nil, err
  }

  return uh, nil
}

func (c *Configuration) SearchUserBans(id int) ([]api.Bans, error) {
  var b []api.Bans

  resp, err := client.Do(c.NewGetRequest(transform.ConvertUser(id, "bans")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertAnime(id, "screenshots")))
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

func (c *Configuration) SearchAnimeFranchise(id int) (api.Franchise, error) {
  var f api.Franchise

  resp, err := client.Do(c.NewGetRequest(transform.ConvertFranchise(id, "animes")))
  if err != nil {
    return f, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return f, err
  }

  return f, nil
}

func (c *Configuration) SearchMangaFranchise(id int) (api.Franchise, error) {
  var f api.Franchise

  resp, err := client.Do(c.NewGetRequest(transform.ConvertFranchise(id, "mangas")))
  if err != nil {
    return f, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return f, err
  }

  return f, nil
}

func (c *Configuration) SearchAnimeExternalLinks(id int) ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks

  resp, err := client.Do(c.NewGetRequest(transform.ConvertExternalLinks(id, "animes")))
  if err != nil {
    return el, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return el, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, err
  }

  return el, nil
}

func (c *Configuration) SearchMangaExternalLinks(id int) ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks

  resp, err := client.Do(c.NewGetRequest(transform.ConvertExternalLinks(id, "mangas")))
  if err != nil {
    return el, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return el, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, err
  }

  return el, nil
}

func (c *Configuration) SearchSimilarAnime(id int) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(transform.ConvertSimilar(id, "animes")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertSimilar(id, "mangas")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertRelated(id, "animes")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertRelated(id, "mangas")))
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

// as a result, we return a complete list of all achievements.
// next comes the filtering through "NekoSearch" and the error about obtaining
// specific achievements is already being processed there.
// see example in README.md
func (c *Configuration) SearchAchievement(id int) ([]api.Achievements, error) {
  var a []api.Achievements

  resp, err := client.Do(c.NewGetRequest(transform.ConvertAchievements(id)))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertAnime(id, "videos")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertRoles(id, "animes")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertRoles(id, "mangas")))
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

  resp, err := client.Do(c.NewGetRequest(transform.ConvertCalendar(r.OptionsCalendar())))
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

func (c *Configuration) SearchGenres() ([]api.Genres, error) {
  var g []api.Genres

  resp, err := client.Do(c.NewGetRequest("genres"))
  if err != nil {
    return g, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return g, err
  }

  if err := json.Unmarshal(data, &g); err != nil {
    return nil, err
  }

  return g, nil
}

func (c *Configuration) SearchStudios() ([]api.Studios, error) {
  var s []api.Studios

  resp, err := client.Do(c.NewGetRequest("studios"))
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

func (c *Configuration) SearchPublishers() ([]api.Publishers, error) {
  var p []api.Publishers

  resp, err := client.Do(c.NewGetRequest("publishers"))
  if err != nil {
    return p, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return p, err
  }

  if err := json.Unmarshal(data, &p); err != nil {
    return nil, err
  }

  return p, nil
}

func (c *Configuration) SearchForums() ([]api.Forums, error) {
  var f []api.Forums

  resp, err := client.Do(c.NewGetRequest("forums"))
  if err != nil {
    return f, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return nil, err
  }

  return f, nil
}

func (c *Configuration) AddFriend(id int) (api.FriendRequest, error) {
  var f api.FriendRequest

  resp, err := client.Do(c.NewPostRequest(transform.ConvertFriend(id)))
  if err != nil {
    return f, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return f, err
  }

  return f, nil
}

func (c *Configuration) RemoveFriend(id int) (api.FriendRequest, error) {
  var f api.FriendRequest

  resp, err := client.Do(c.NewDeleteRequest(transform.ConvertFriend(id)))
  if err != nil {
    return f, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return f, err
  }

  return f, nil
}
