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
  "github.com/vexilology/goshikimori/str"
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
  Page, Limit, Kind, Status, Season, Score, Rating string
}

type ExtraLimit struct {
  Page, Limit string
}

type ExtraCensored struct {
  Censored string
}

type ExtraAnimeRates struct {
  Page, Limit, Status, Censored string
}

type ExtraMangaRates struct {
  Page, Limit, Censored string
}

type ExtraTargetType struct {
  Page, Limit, Target_id, Target_type string
}

type ExtraMessages struct {
  Page, Limit, Type string
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

type ResultMessages interface {
  OptionsMessages() string
}

func Add(app, tok string) *Configuration {
  return &Configuration{Application: app, AccessToken: tok}
}

// string formatting for achievements search
func NekoSearch(name string) string {
  r := strings.Replace(strings.ToLower(name), " ", "_", -1)
  return fmt.Sprintf("%s", r)
}

// Page  -> 100000 maximum
// Limit -> 100 maximum
// Type  -> inbox, private, sent, news, notifications
func (em *ExtraMessages) OptionsMessages() string {
  p, _ := strconv.Atoi(em.Page)
  l, _ := strconv.Atoi(em.Limit)

  if p == 0 { em.Page = "1" }
  if l == 0 { em.Limit = "1" }
  for i := 100001; i <= p; i++ {
    em.Page = "1"
  }
  for i := 101; i <= l; i++ {
    em.Limit = "1"
  }

  target_map := map[string]int8{
    "inbox": 1, "private": 2, "sent": 3,
    "news": 4, "notifications": 5,
  }
  _, ok = target_map[em.Type]
  if !ok { em.Type = "news" }

  v := url.Values{}
  v.Add("type", em.Type)
  v.Add("page", em.Page)
  v.Add("limit", em.Limit)

  return v.Encode()
}

// Page        -> 100000 maximum
// Limit       -> 100 maximum
// Target_id   -> id anime/manga in string format
// Target_type -> Anime, Manga
func (ett *ExtraTargetType) OptionsUserHistory() string {
  p, _ := strconv.Atoi(ett.Page)
  l, _ := strconv.Atoi(ett.Limit)

  if p == 0 { ett.Page = "1" }
  if l == 0 { ett.Limit = "1" }
  for i := 100001; i <= p; i++ {
    ett.Page = "1"
  }
  for i := 101; i <= l; i++ {
    ett.Limit = "1"
  }

  target_map := map[string]int8{"Anime": 1, "Manga": 2}
  _, ok = target_map[ett.Target_type]
  if !ok { ett.Target_type = "Anime" }

  v := url.Values{}
  v.Add("page", ett.Page)
  v.Add("limit", ett.Limit)
  // NOTES: we get an error if we do not process the request in this way
  // json: cannot unmarshal string into Go value of type api.UserHistory
  if ett.Target_id != "" { v.Add("target_id", ett.Target_id) }
  v.Add("target_type", ett.Target_type)

  return v.Encode()
}

// Page  -> 100000 maximum
// Limit -> 100 maximum
func (el *ExtraLimit) OptionsUsers() string {
  p, _ := strconv.Atoi(el.Page)
  l, _ := strconv.Atoi(el.Limit)

  if p == 0 { el.Page = "1" }
  if l == 0 { el.Limit = "1" }
  for i := 100001; i <= p; i++ {
    el.Page = "1"
  }
  for i := 101; i <= l; i++ {
    el.Limit = "1"
  }

  v := url.Values{}
  v.Add("page", el.Page)
  v.Add("limit", el.Limit)

  return v.Encode()
}

// Page   -> 100000 maximum
// Limit  -> 50 maximum
// Order  -> check RandomAnime()
// Type   -> "Deprecated"
// Kind   -> tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48
// Status -> anons, ongoing, released
// Season -> summer_2017, 2016, 2014_2016, 199x
// Score  -> 9 maximum
// FIXME  Duration -> not supported
// Rating -> none, g, pg, pg_13, r, r_plus, rx
// FIXME  Genre -> not supported
// FIXME  Studio -> not supported
// FIXME  Franchise -> not supported
// FIXME  Censored -> not supported
// FIXME  Mylist -> not supported
// FIXME  Ids -> not supported
// FIXME  Exclude_ids -> not supported
// Search -> default search
func (e *Extra) OptionsAnime() string {
  p, _ := strconv.Atoi(e.Page)
  l, _ := strconv.Atoi(e.Limit)

  if p == 0 { e.Page = "1" }
  if l == 0 { e.Limit = "1" }
  for i := 100001; i <= p; i++ {
    e.Page = "1"
  }
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
  v.Add("page", e.Page)
  v.Add("limit", e.Limit)
  v.Add("kind", e.Kind)
  v.Add("status", e.Status)
  v.Add("season", e.Season)
  v.Add("score", e.Score)
  v.Add("rating", e.Rating)

  return v.Encode()
}

// Page   -> 100000 maximum
// Limit  -> 50 maximum
// Order  -> check RandomManga()
// Type   -> "Deprecated"
// Kind   -> manga, manhwa, manhua,
//           light_novel, novel, one_shot, doujin
// Status -> anons, ongoing, released, paused, discontinued
// Season -> summer_2017, "spring_2016,fall_2016",
//           "2016,!winter_2016", 2016, 2014_2016, 199x
// Score  -> 9 maximum
// FIXME  Genre -> not supported
// FIXME  Publisher -> not supported
// FIXME  Franchise -> not supported
// FIXME  Censored -> not supported
// FIXME  Mylist -> not supported
// FIXME  Ids -> not supported
// FIXME  Exclude_ids -> not supported
// Search -> default search
func (e *Extra) OptionsManga() string {
  p, _ := strconv.Atoi(e.Page)
  l, _ := strconv.Atoi(e.Limit)

  if p == 0 { e.Page = "1" }
  if l == 0 { e.Limit = "1" }
  for i := 100001; i <= p; i++ {
    e.Page = "1"
  }
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
  v.Add("page", e.Page)
  v.Add("limit", e.Limit)
  v.Add("kind", e.Kind)
  v.Add("status", e.Status)
  v.Add("season", e.Season)
  v.Add("score", e.Score)

  return v.Encode()
}

// Page  -> 100000 maximum
// Limit -> 30 maximum
// Search -> default search
func (el *ExtraLimit) OptionsClub() string {
  p, _ := strconv.Atoi(el.Page)
  l, _ := strconv.Atoi(el.Limit)

  if p == 0 { el.Page = "1" }
  if l == 0 { el.Limit = "1" }
  for i := 100001; i <= p; i++ {
    el.Page = "1"
  }
  for i := 31; i <= l; i++ {
    el.Limit = "1"
  }

  v := url.Values{}
  v.Add("page", el.Page)
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

// Page     -> 100000 maximum
// Limit    -> 5000 maximum
// Status   -> planned, watching, rewatching, completed, on_hold, dropped
// Censored -> true, false
// Set to true to discard hentai, yaoi and yuri
func (ar *ExtraAnimeRates) OptionsAnimeRates() string {
  p, _ := strconv.Atoi(ar.Page)
  l, _ := strconv.Atoi(ar.Limit)

  if p == 0 { ar.Page = "1" }
  if l == 0 { ar.Limit = "1" }
  for i := 100001; i <= p; i++ {
    ar.Page = "1"
  }
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
  v.Add("page", ar.Page)
  v.Add("limit", ar.Limit)
  v.Add("status", ar.Status)
  v.Add("censored", ar.Censored)

  return v.Encode()
}

// Page     -> 100000 maximum
// Limit    -> 5000 maximum
// Censored -> true, false
// Set to true to discard hentai, yaoi and yuri
func (mr *ExtraMangaRates) OptionsMangaRates() string {
  p, _ := strconv.Atoi(mr.Page)
  l, _ := strconv.Atoi(mr.Limit)

  if p == 0 { mr.Page = "1" }
  if l == 0 { mr.Limit = "1" }
  for i := 100001; i <= p; i++ {
    mr.Page = "1"
  }
  for i := 5001; i <= l; i++ {
    mr.Limit = "1"
  }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok = censored_map[mr.Censored]
  if !ok { mr.Censored = "false" }

  v := url.Values{}
  v.Add("page", mr.Page)
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return nil, err
  }

  return u, nil
}

func (c *Configuration) SearchUserFriends(id int) ([]api.UserFriends, error) {
  var uf []api.UserFriends

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "friends")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return nil, err
  }

  return uf, nil
}

func (c *Configuration) SearchUserClubs(id int) ([]api.Clubs, error) {
  var uc []api.Clubs

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "clubs")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uc); err != nil {
    return nil, err
  }

  return uc, nil
}

func (c *Configuration) SearchUserAnimeRates(id int, r ResultAnimeRates) ([]api.UserAnimeRates, error) {
  var ar []api.UserAnimeRates

  resp, err := client.Do(c.NewGetRequest(
    str.ConvertUserRates(id, "anime_rates", r.OptionsAnimeRates()),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &ar); err != nil {
    return nil, err
  }

  return ar, nil
}

func (c *Configuration) SearchUserMangaRates(id int, r ResultMangaRates) ([]api.UserMangaRates, error) {
  var mr []api.UserMangaRates

  resp, err := client.Do(c.NewGetRequest(
    str.ConvertUserRates(id, "manga_rates", r.OptionsMangaRates()),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &mr); err != nil {
    return nil, err
  }

  return mr, nil
}

func (c *Configuration) SearchUserFavourites(id int) (api.UserFavourites, error) {
  var uf api.UserFavourites

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "favourites")))
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
    str.ConvertUserRates(id, "history", r.OptionsUserHistory()),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uh); err != nil {
    return nil, err
  }

  return uh, nil
}

func (c *Configuration) SearchUserBans(id int) ([]api.Bans, error) {
  var b []api.Bans

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "bans")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

func (c *Configuration) FastIdAnime(name string) (int, error) {
  var a []api.Animes
  var aa api.Animes

  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(name),
  ))
  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()

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
  var m []api.Mangas
  var mm api.Mangas

  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(name),
  ))
  if err != nil {
    return 0, err
  }

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

// FIXME: the output is always 2 results and
// we write the last number, limit not working
func (c *Configuration) FastIdClub(name string) (int, error) {
  var cl []api.Clubs
  var clcl api.Clubs

  resp, err := client.Do(c.NewGetRequest(
    "clubs?search=" + url.QueryEscape(name),
  ))
  if err != nil {
    return 0, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return 0, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return 0, err
  }

  for _, v := range cl {
    clcl = v
  }

  return clcl.Id, nil
}

func (c *Configuration) SearchAnimeScreenshots(id int) ([]api.AnimeScreenshots, error) {
  var s []api.AnimeScreenshots

  resp, err := client.Do(c.NewGetRequest(str.ConvertAnime(id, "screenshots")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &s); err != nil {
    return nil, err
  }

  return s, nil
}

func (c *Configuration) SearchAnimeFranchise(id int) (api.Franchise, error) {
  var f api.Franchise

  resp, err := client.Do(c.NewGetRequest(str.ConvertFranchise(id, "animes")))
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

  resp, err := client.Do(c.NewGetRequest(str.ConvertFranchise(id, "mangas")))
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

  resp, err := client.Do(c.NewGetRequest(str.ConvertExternalLinks(id, "animes")))
  if err != nil {
    return nil, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, err
  }

  return el, nil
}

func (c *Configuration) SearchMangaExternalLinks(id int) ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks

  resp, err := client.Do(c.NewGetRequest(str.ConvertExternalLinks(id, "mangas")))
  if err != nil {
    return nil, err
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, err
  }

  return el, nil
}

func (c *Configuration) SearchSimilarAnime(id int) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(str.ConvertSimilar(id, "animes")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchSimilarManga(id int) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest(str.ConvertSimilar(id, "mangas")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

func (c *Configuration) SearchRelatedAnime(id int) ([]api.RelatedAnimes, error) {
  var a []api.RelatedAnimes

  resp, err := client.Do(c.NewGetRequest(str.ConvertRelated(id, "animes")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchRelatedManga(id int) ([]api.RelatedMangas, error) {
  var m []api.RelatedMangas

  resp, err := client.Do(c.NewGetRequest(str.ConvertRelated(id, "mangas")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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

  resp, err := client.Do(c.NewGetRequest(str.ConvertAchievements(id)))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) SearchAnimeVideos(id int) ([]api.AnimeVideos, error) {
  var v []api.AnimeVideos

  resp, err := client.Do(c.NewGetRequest(str.ConvertAnime(id, "videos")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &v); err != nil {
    return nil, err
  }

  return v, nil
}

func (c *Configuration) SearchAnimeRoles(id int) ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(c.NewGetRequest(str.ConvertRoles(id, "animes")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, err
  }

  return r, nil
}

func (c *Configuration) SearchMangaRoles(id int) ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(c.NewGetRequest(str.ConvertRoles(id, "mangas")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &b); err != nil {
    return nil, err
  }

  return b, nil
}

func (c *Configuration) SearchCalendar(r ResultCensored) ([]api.Calendar, error) {
  var ca []api.Calendar

  resp, err := client.Do(c.NewGetRequest(str.ConvertCalendar(r.OptionsCalendar())))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return nil, err
  }

  return f, nil
}

func (c *Configuration) AddFriend(id int) (api.FriendRequest, error) {
  var f api.FriendRequest

  resp, err := client.Do(c.NewPostRequest(str.ConvertFriend(id)))
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

  resp, err := client.Do(c.NewDeleteRequest(str.ConvertFriend(id)))
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

func (c *Configuration) UserUnreadMessages(id int) (api.UnreadMessages, error) {
  var um api.UnreadMessages

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "unread_messages")))
  if err != nil {
    return um, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return um, err
  }

  if err := json.Unmarshal(data, &um); err != nil {
    return um, err
  }

  return um, nil
}

func (c *Configuration) UserMessages(id int, r ResultMessages) ([]api.Messages, error) {
  var m []api.Messages

  resp, err := client.Do(c.NewGetRequest(str.ConvertMessages(id, r.OptionsMessages())))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

func (c *Configuration) SearchConstantsAnime() (api.Constants, error) {
  var ca api.Constants

  resp, err := client.Do(c.NewGetRequest(str.ConvertConstants("anime")))
  if err != nil {
    return ca, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return ca, err
  }

  if err := json.Unmarshal(data, &ca); err != nil {
    return ca, err
  }

  return ca, nil
}

func (c *Configuration) SearchConstantsManga() (api.Constants, error) {
  var cm api.Constants

  resp, err := client.Do(c.NewGetRequest(str.ConvertConstants("manga")))
  if err != nil {
    return cm, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return cm, err
  }

  if err := json.Unmarshal(data, &cm); err != nil {
    return cm, err
  }

  return cm, nil
}

func (c *Configuration) SearchConstantsUserRate() (api.ConstantsUserRate, error) {
  var ur api.ConstantsUserRate

  resp, err := client.Do(c.NewGetRequest(str.ConvertConstants("user_rate")))
  if err != nil {
    return ur, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return ur, err
  }

  if err := json.Unmarshal(data, &ur); err != nil {
    return ur, err
  }

  return ur, nil
}

func (c *Configuration) SearchConstantsClub() (api.ConstantsClub, error) {
  var cc api.ConstantsClub

  resp, err := client.Do(c.NewGetRequest(str.ConvertConstants("club")))
  if err != nil {
    return cc, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return cc, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return cc, err
  }

  return cc, nil
}

func (c *Configuration) SearchConstantsSmileys() ([]api.ConstantsSmileys, error) {
  var cs []api.ConstantsSmileys

  resp, err := client.Do(c.NewGetRequest(str.ConvertConstants("smileys")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &cs); err != nil {
    return nil, err
  }

  return cs, nil
}

func (c *Configuration) RandomAnime() ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest("animes?order=random"))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

func (c *Configuration) RandomManga() ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest("mangas?order=random"))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}
