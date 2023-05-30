// Copyright (C) 2023 vexilology <hey_h0n3y@protonmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Comments are made in the style of "godoc" syntax support.
//
// More information can be found in the [examples] folder.
//
// [examples]: https://github.com/vexilology/goshikimori/blob/main/examples/
package goshikimori

import (
  "fmt"
  "net/http"
  "io"
  "strings"
  "encoding/json"
  "net/url"
  "strconv"
  "context"
  "time"
  "errors"
  "bytes"

  "github.com/vexilology/goshikimori/api"
  "github.com/vexilology/goshikimori/str"
)

const site = "shikimori.me/api"

var client = &http.Client{}

type Configuration struct {
  Application, AccessToken string
}

type FastId struct {
  Id   int
  Conf Configuration
  Err  error
}

type Options struct {
  Page, Limit, Kind, Status, Season, Score, Rating, Censored, Type, Target_id, Target_type string
}

type Result interface {
  OptionsAnime()           string
  OptionsManga()           string
  OptionsUsers()           string
  OptionsClub()            string
  OptionsCalendar()        string
  OptionsAnimeRates()      string
  OptionsMangaRates()      string
  OptionsUserHistory()     string
  OptionsMessages()        string
  OptionsPeople()          string
  OptionsClubInformation() string
}

// You need to enter the application and the private key.
//
// To register the application, follow the link from [OAuth].
//
// [OAuth]: https://github.com/vexilology/goshikimori#shikimori-documentation
func Add(app, tok string) *Configuration {
  return &Configuration{Application: app, AccessToken: tok}
}

// String formatting for achievements search. Check [example].
//
// [example]: https://github.com/vexilology/goshikimori/blob/main/examples/achievements/main.go
func NekoSearch(name string) string {
  return strings.Replace(strings.ToLower(name), " ", "_", -1)
}

func (o *Options) OptionsMessages() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 101 { o.Limit = "1" }

  target_map := map[string]int8{
    "inbox": 1, "private": 2, "sent": 3,
    "news": 4, "notifications": 5,
  }
  _, ok := target_map[o.Type]; if !ok {
    o.Type = "news"
  }

  v := url.Values{}
  v.Add("type", o.Type)
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)

  return v.Encode()
}

func (o *Options) OptionsUserHistory() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 101 { o.Limit = "1" }

  target_map := map[string]int8{"Anime": 1, "Manga": 2}
  _, ok := target_map[o.Target_type]; if !ok {
    o.Target_type = "Anime"
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  // We get an error if we do not process the request in this way.
  // json: cannot unmarshal string into Go value of type api.UserHistory
  if o.Target_id != "" { v.Add("target_id", o.Target_id) }
  v.Add("target_type", o.Target_type)

  return v.Encode()
}

func (o *Options) OptionsUsers() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 101 { o.Limit = "1" }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)

  return v.Encode()
}

func (o *Options) OptionsAnime() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 51 { o.Limit = "1" }

  kind_map := map[string]int8{
    "tv": 1, "movie": 2, "ova": 3, "ona": 4,
    "special": 5, "music": 6,
    "tv_13": 7, "tv_24": 8, "tv_48": 9,
  }
  _, ok_kind := kind_map[o.Kind]; if !ok_kind {
    o.Kind = ""
  }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = ""
  }

  season_map := map[string]int8{
    "summer_2017": 1, "2016": 2, "2014_2016": 3, "199x": 4,
  }
  _, ok_season := season_map[o.Season]; if !ok_season {
    o.Season = ""
  }

  s, _ := strconv.Atoi(o.Score)
  if s >= 10 { o.Score = "" }

  rating_map := map[string]int8{
    "none": 1, "g": 2, "pg": 3, "pg_13": 4,
    "r": 5, "r_plus": 6, "rx": 7,
  }
  _, ok_rating := rating_map[o.Rating]; if !ok_rating {
    o.Rating = ""
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("kind", o.Kind)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("score", o.Score)
  v.Add("rating", o.Rating)

  return v.Encode()
}

func (o *Options) OptionsManga() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 51 { o.Limit = "1" }

  kind_map := map[string]int8{
    "manga": 1, "manhwa": 2, "manhua": 3,
    "light_novel": 5, "novel": 6,
    "one_shot": 7, "doujin": 8,
  }
  _, ok_kind := kind_map[o.Kind]; if !ok_kind {
    o.Kind = ""
  }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3,
    "paused": 4, "discontinued": 5,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = ""
  }

  season_map := map[string]int8{
    "summer_2017": 1, "spring_2016,fall_2016": 2,
    "2016,!winter_2016": 3, "2016": 4,
    "2014_2016": 5, "199x": 6,
  }
  _, ok_season := season_map[o.Season]; if !ok_season {
    o.Season = ""
  }

  s, _ := strconv.Atoi(o.Score)
  if s >= 10 { o.Score = "" }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("kind", o.Kind)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("score", o.Score)

  return v.Encode()
}

func (o *Options) OptionsClub() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 31 { o.Limit = "1" }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)

  return v.Encode()
}

func (o *Options) OptionsCalendar() string {
  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok := censored_map[o.Censored]; if !ok {
    o.Censored = "false"
  }

  v := url.Values{}
  v.Add("censored", o.Censored)

  return v.Encode()
}

func (o *Options) OptionsAnimeRates() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 5001 { o.Limit = "1" }

  status_map := map[string]int8{
    "planned": 1, "watching": 2,
    "rewatching": 3, "completed": 4,
    "on_hold": 5, "dropped": 6,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = "watching"
  }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok_censored := censored_map[o.Censored]; if !ok_censored {
    o.Censored = "false"
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("status", o.Status)
  v.Add("censored", o.Censored)

  return v.Encode()
}

func (o *Options) OptionsMangaRates() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p == 0 || p >= 100001 { o.Page = "1" }
  if l == 0 || l >= 5001 { o.Limit = "1" }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok := censored_map[o.Censored]; if !ok {
    o.Censored = "false"
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("censored", o.Censored)

  return v.Encode()
}

func (o *Options) OptionsPeople() string {
  kind_map := map[string]int8{
    "seyu": 1, "mangaka": 2, "producer": 3,
  }
  _, ok := kind_map[o.Kind]; if !ok {
    o.Kind = "seyu"
  }

  v := url.Values{}
  v.Add("kind", o.Kind)

  return v.Encode()
}

func (o *Options) OptionsClubInformation() string {
  p, _ := strconv.Atoi(o.Page)

  if p == 0 || p >= 100001 { o.Page = "1" }

  v := url.Values{}
  v.Add("page", o.Page)

  return v.Encode()
}

// FIXME cancel does not work. the context is canceled at once.
// Duration works. When the time expires, the request is cancelled.
func ctx(number time.Duration) context.Context {
  duration := number * time.Second
  ctx, _ := context.WithTimeout(context.Background(), duration)
  return ctx
}

func (c *Configuration) NewGetRequest(search string) *http.Request {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  // ctx(10) -> query time 10 seconds,
  // in the future it will be possible to make the parameter dynamic.
  req, _ := http.NewRequestWithContext(ctx(10), http.MethodGet, custom_url, nil)
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", "Bearer " + c.AccessToken)
  return req
}

// To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func (c *Configuration) NewPostRequest(search string) *http.Request {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  // ctx(10) -> query time 10 seconds,
  // in the future it will be possible to make the parameter dynamic.
  req, _ := http.NewRequestWithContext(ctx(10), http.MethodPost, custom_url, nil)
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", "Bearer " + c.AccessToken)
  return req
}

// Custom POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func (c *Configuration) NewCustomPostRequest(search string, position int) *http.Request {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(`"new_index": "%d"`, position))
  // ctx(10) -> query time 10 seconds,
  // in the future it will be possible to make the parameter dynamic.
  req, _ := http.NewRequestWithContext(
    ctx(10), http.MethodPost, custom_url, bytes.NewBuffer(data),
  )
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", "Bearer " + c.AccessToken)
  req.Header.Set("Content-Type", "application/json")
  return req
}

// To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func (c *Configuration) NewDeleteRequest(search string) *http.Request {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  // ctx(10) -> query time 10 seconds,
  // in the future it will be possible to make the parameter dynamic.
  req, _ := http.NewRequestWithContext(ctx(10), http.MethodDelete, custom_url, nil)
  req.Header.Add("User-Agent", c.Application)
  req.Header.Add("Authorization", "Bearer " + c.AccessToken)
  return req
}

// Name: user name.
//
// Search by user is case sensitive.
func (c *Configuration) SearchUser(name string) (api.Users, error) {
  var u api.Users

  resp, err := client.Do(c.NewGetRequest("users/" + url.QueryEscape(name)))
  if err != nil {
    return u, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return u, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return u, err
  }

  return u, nil
}

// Name: user name.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//
// 'Options' settings
// 	- Page: 100000 maximum;
// 	- Limit: 100 maximum;
//
// Don't use Stats.Statuses.Anime and Stats.Statuses.Manga: empty slice.
func (c *Configuration) SearchUsers(name string, r Result) ([]api.Users, error) {
  var u []api.Users

  resp, err := client.Do(
    c.NewGetRequest("users?search=" + url.QueryEscape(name) + "&" + r.OptionsUsers()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return nil, err
  }

  return u, nil
}

// Id: user id.
func (c *Configuration) SearchUserFriends(id int) ([]api.UserFriends, error) {
  var uf []api.UserFriends

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "friends")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return nil, err
  }

  return uf, nil
}

// Id: user id.
func (c *Configuration) SearchUserClubs(id int) ([]api.Clubs, error) {
  var uc []api.Clubs

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "clubs")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uc); err != nil {
    return nil, err
  }

  return uc, nil
}

// Id: user id.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//	- Status: watching;
//	- Censored: false;
//
// 'Options' settings
//	- Page: 100000 maximum;
//	- Limit: 5000 maximum;
//	- Status: planned, watching, rewatching, completed, on_hold, dropped;
//	- Censored: true, false;
//
// Set to true to discard hentai, yaoi and yuri.
func (c *Configuration) SearchUserAnimeRates(id int, r Result) ([]api.UserAnimeRates, error) {
  var ar []api.UserAnimeRates

  resp, err := client.Do(c.NewGetRequest(
    str.ConvertUserRates(id, "anime_rates", r.OptionsAnimeRates()),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &ar); err != nil {
    return nil, err
  }

  return ar, nil
}

// Id: user id.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//	- Censored: false;
//
// 'Options' Settings
//	- Page: 100000 maximum;
//	- Limit: 5000 maximum;
//	- Censored: true, false;
//
// Set to true to discard hentai, yaoi and yuri.
func (c *Configuration) SearchUserMangaRates(id int, r Result) ([]api.UserMangaRates, error) {
  var mr []api.UserMangaRates

  resp, err := client.Do(c.NewGetRequest(
    str.ConvertUserRates(id, "manga_rates", r.OptionsMangaRates()),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &mr); err != nil {
    return nil, err
  }

  return mr, nil
}

// Id: user id.
func (c *Configuration) SearchUserFavourites(id int) (api.UserFavourites, error) {
  var uf api.UserFavourites

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "favourites")))
  if err != nil {
    return uf, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return uf, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return uf, err
  }

  return uf, nil
}

// Id: user id.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//	- Target_type: Anime;
//	- Target_id: option is hidden if empty;
//
// 'Options' settings
// 	- Page: 100000 maximum.
// 	- Limit: 100 maximum.
// 	- Target_id: id anime/manga in string format.
// 	- Target_type: Anime, Manga.
func (c *Configuration) SearchUserHistory(id int, r Result) ([]api.UserHistory, error) {
  var uh []api.UserHistory

  resp, err := client.Do(c.NewGetRequest(
    str.ConvertUserRates(id, "history", r.OptionsUserHistory()),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uh); err != nil {
    return nil, err
  }

  return uh, nil
}

// Id: user id.
func (c *Configuration) SearchUserBans(id int) ([]api.Bans, error) {
  var b []api.Bans

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "bans")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return w, err
  }

  if err := json.Unmarshal(data, &w); err != nil {
    return w, err
  }

  return w, nil
}

// Name: anime name.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//	- Kind: empty field;
//	- Status: empty field;
//	- Season: empty field;
//	- Score: empty field;
//	- Rating: empty field;
//
// 'Options' settings
//	- Page: 100000 maximum;
//	- Limit: 50 maximum;
//	- Order: check [RandomAnime];
//	- Type: "Deprecated";
//	- Kind: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48;
//	- Status: anons, ongoing, released;
//	- Season: summer_2017, 2016, 2014_2016, 199x;
//	- Score: 9 maximum;
//	- Rating: none, g, pg, pg_13, r, r_plus, rx;
//	- Search: default search;
//
// [RandomAnime]: https://github.com/vexilology/goshikimori/blob/main/examples/random/main.go
//
// FIXME
//	- Duration: not supported;
//	- Genre: not supported;
//	- Studio: not supported;
//	- Franchise: not supported;
//	- Censored: not supported;
//	- Mylist: not supported;
//	- Ids: not supported;
//	- Exclude_ids: not supported;
func (c *Configuration) SearchAnime(name string, r Result) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest(
    "animes?search=" + url.QueryEscape(name) + "&" + r.OptionsAnime(),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

// Name: manga name.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//	- Kind: empty field;
//	- Status: empty field;
//	- Season: empty field;
//	- Score: empty field;
//
// 'Options' settings
//	- Page: 100000 maximum;
//	- Limit: 50 maximum;
//	- Order: check [RandomManga];
//	- Type: "Deprecated";
//	- Kind: manga, manhwa, manhua, light_novel, novel, one_shot, doujin;
//	- Status: anons, ongoing, released, paused, discontinued;
//	- Season: summer_2017, "spring_2016,fall_2016", "2016,!winter_2016", 2016, 2014_2016, 199x;
//	- Score: 9 maximum;
//	- Search: default search;
//
// [RandomManga]: https://github.com/vexilology/goshikimori/blob/main/examples/random/main.go
//
// FIXME
//	- Genre: not supported;
//	- Publisher: not supported;
//	- Franchise: not supported;
//	- Censored: not supported;
//	- Mylist: not supported;
//	- Ids: not supported;
//	- Exclude_ids: not supported;
func (c *Configuration) SearchManga(name string, r Result) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest(
    "mangas?search=" + url.QueryEscape(name) + "&" + r.OptionsManga(),
  ))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

// Name: anime name.
func (c *Configuration) FastIdAnime(name string) *FastId {
  var a []api.Animes

  resp, err := client.Do(c.NewGetRequest("animes?search=" + url.QueryEscape(name)))
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(a) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil} }

  return &FastId{Id: a[0].Id, Conf: *c, Err: nil}
}

// Name: manga name.
func (c *Configuration) FastIdManga(name string) *FastId {
  var m []api.Mangas

  resp, err := client.Do(c.NewGetRequest("mangas?search=" + url.QueryEscape(name)))
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(m) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil} }

  return &FastId{Id: m[0].Id, Conf: *c, Err: nil}
}

// Name: club name.
func (c *Configuration) FastIdClub(name string) *FastId {
  var cl []api.Clubs

  resp, err := client.Do(c.NewGetRequest("clubs?search=" + url.QueryEscape(name)))
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(cl) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil} }

  return &FastId{Id: cl[0].Id, Conf: *c, Err: nil}
}

// Name: people name.
func (c *Configuration) FastIdPeople(name string) *FastId {
  var ap []api.AllPeople

  resp, err := client.Do(c.NewGetRequest("people/search?search=" + url.QueryEscape(name)))
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  if err := json.Unmarshal(data, &ap); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(ap) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil} }

  return &FastId{Id: ap[0].Id, Conf: *c, Err: nil}
}

// *Configuration.FastIdAnime(name string).SearchAnimeScreenshots()
func (f *FastId) SearchAnimeScreenshots() ([]api.AnimeScreenshots, error) {
  var s []api.AnimeScreenshots

  resp, err := client.Do(f.Conf.NewGetRequest(str.ConvertAnime(f.Id, "screenshots")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &s); err != nil {
    return nil, err
  }

  return s, nil
}

// *Configuration.FastIdAnime(name string).SearchAnimeFranchise()
func (f *FastId) SearchAnimeFranchise() (api.Franchise, error) {
  var ff api.Franchise

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertFranchise(f.Id, "animes")),
  )
  if err != nil {
    return ff, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ff, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, err
  }

  return ff, nil
}

// *Configuration.FastIdManga(name string).SearchMangaFranchise()
func (f *FastId) SearchMangaFranchise() (api.Franchise, error) {
  var ff api.Franchise

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertFranchise(f.Id, "mangas")),
  )
  if err != nil {
    return ff, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ff, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, err
  }

  return ff, nil
}

// *Configuraiton.FastIdAnime(name string).SearchAnimeExternalLinks()
func (f *FastId) SearchAnimeExternalLinks() ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertExternalLinks(f.Id, "animes")),
  )
  if err != nil {
    return nil, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, err
  }

  return el, nil
}

// *Configuraiton.FastIdManga(name string).SearchMangaExternalLinks()
func (f *FastId) SearchMangaExternalLinks() ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertExternalLinks(f.Id, "mangas")),
  )
  if err != nil {
    return nil, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, err
  }

  return el, nil
}

// *Configuraiton.FastIdAnime(name string).SearchSimilarAnime()
func (f *FastId) SearchSimilarAnime() ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertSimilar(f.Id, "animes")),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

// *Configuraiton.FastIdManga(name string).SearchSimilarManga()
func (f *FastId) SearchSimilarManga() ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertSimilar(f.Id, "mangas")),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

// *Configuraiton.FastIdAnime(name string).SearchRelatedAnime()
func (f *FastId) SearchRelatedAnime() ([]api.RelatedAnimes, error) {
  var a []api.RelatedAnimes

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertRelated(f.Id, "animes")),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

// *Configuraiton.FastIdManga(name string).SearchRelatedManga()
func (f *FastId) SearchRelatedManga() ([]api.RelatedMangas, error) {
  var m []api.RelatedMangas

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertRelated(f.Id, "mangas")),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

// Name: club name.
//
// If 'Options' empty fields
// 	- Page: 1;
// 	- Limit: 1;
//
// 'Options' settings
//	- Page: 100000 maximum;
//	- Limit: 30 maximum;
//	- Search: default search;
//
// If we set the limit=1, we will still have 2 results.
func (c *Configuration) SearchClub(name string, r Result) ([]api.Clubs, error) {
  var cl []api.Clubs

  resp, err := client.Do(
    c.NewGetRequest("clubs?search=" + url.QueryEscape(name) + "&" + r.OptionsClub()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return nil, err
  }

  return cl, nil
}

// If 'Options' empty fields
//	- Page: 1;
//
// 'Options' settings
//	- Page: 100000 maximum;
//
// *Configuration.FastIdClub(name string).SearchClubAnimes(r Result)
func (f *FastId) SearchClubAnimes(r Result) ([]api.Animes, error) {
  var a []api.Animes

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertClub(f.Id, "animes") + "?" + r.OptionsClubInformation()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

// If 'Options' empty fields
//	- Page: 1;
//
// 'Options' settings
//	- Page: 100000 maximum;
//
// *Configuration.FastIdClub(name string).SearchClubMangas(r Result)
func (f *FastId) SearchClubMangas(r Result) ([]api.Mangas, error) {
  var m []api.Mangas

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertClub(f.Id, "mangas") + "?" + r.OptionsClubInformation()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

// If 'Options' empty fields
//	- Page: 1;
//
// 'Options' settings
//	- Page: 100000 maximum;
//
// *Configuration.FastIdClub(name string).SearchClubCharacters(r Result)
func (f *FastId) SearchClubCharacters(r Result) ([]api.CharacterInfo, error) {
  var ci []api.CharacterInfo

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertClub(f.Id, "characters") + "?" + r.OptionsClubInformation()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &ci); err != nil {
    return nil, err
  }

  return ci, nil
}

// If 'Options' empty fields
//	- Page: 1;
//
// 'Options' settings
//	- Page: 100000 maximum;
//
// *Configuration.FastIdClub(name string).SearchClubClubs(r Result)
func (f *FastId) SearchClubClubs(r Result) ([]api.Clubs, error) {
  var cc []api.Clubs

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertClub(f.Id, "clubs") + "?" + r.OptionsClubInformation()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return nil, err
  }

  return cc, nil
}

// If 'Options' empty fields
//	- Page: 1;
//
// 'Options' settings
//	- Page: 100000 maximum;
//
// *Configuration.FastIdClub(name string).SearchClubCollections(r Result)
func (f *FastId) SearchClubCollections(r Result) ([]api.ClubCollections, error) {
  var cc []api.ClubCollections

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertClub(f.Id, "collections") + "?" + r.OptionsClubInformation()),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return nil, err
  }

  return cc, nil
}

// *Configuration.FastIdClub(name string).SearchClubMembers()
func (f *FastId) SearchClubMembers() ([]api.UserFriends, error) {
  var uf []api.UserFriends

  resp, err := client.Do(f.Conf.NewGetRequest(str.ConvertClub(f.Id, "members")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return nil, err
  }

  return uf, nil
}

// *Configuration.FastIdClub(name string).SearchClubImages()
func (f *FastId) SearchClubImages() ([]api.ClubImages, error) {
  var cm []api.ClubImages

  resp, err := client.Do(f.Conf.NewGetRequest(str.ConvertClub(f.Id, "images")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &cm); err != nil {
    return nil, err
  }

  return cm, nil
}

// You can only get a StatusCode.
//
// *Configuration.FastIdClub(name string).ClubJoin()
func (f *FastId) ClubJoin() (int, error) {
  resp, err := client.Do(
    f.Conf.NewPostRequest(str.ConvertClub(f.Id, "join")),
  )
  if err != nil {
    return 500, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// You can only get a StatusCode.
//
// *Configuration.FastIdClub(name string).ClubLeave()
func (f *FastId) ClubLeave() (int, error) {
  resp, err := client.Do(
    f.Conf.NewPostRequest(str.ConvertClub(f.Id, "leave")),
  )
  if err != nil {
    return 500, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// Id: user id.
//
// As a result, we return a complete list of all achievements.
//
// Next comes the filtering through "NekoSearch" and the error about obtaining
// specific achievements is already being processed there.
//
// Check [example].
//
// [example]: https://github.com/vexilology/goshikimori/blob/main/examples/achievements/main.go
func (c *Configuration) SearchAchievement(id int) ([]api.Achievements, error) {
  var a []api.Achievements

  resp, err := client.Do(c.NewGetRequest(str.ConvertAchievements(id)))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, err
  }

  return a, nil
}

// *Configuration.FastIdAnime(name string).SearchAnimeVideos()
func (f *FastId) SearchAnimeVideos() ([]api.AnimeVideos, error) {
  var v []api.AnimeVideos

  resp, err := client.Do(f.Conf.NewGetRequest(str.ConvertAnime(f.Id, "videos")))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &v); err != nil {
    return nil, err
  }

  return v, nil
}

// *Configuraiton.FastIdAnime(name string).SearchAnimeRoles()
func (f *FastId) SearchAnimeRoles() ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertRoles(f.Id, "animes")),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, err
  }

  return r, nil
}

// *Configuraiton.FastIdManga(name string).SearchMangaRoles()
func (f *FastId) SearchMangaRoles() ([]api.Roles, error) {
  var r []api.Roles

  resp, err := client.Do(
    f.Conf.NewGetRequest(str.ConvertRoles(f.Id, "mangas")),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &b); err != nil {
    return nil, err
  }

  return b, nil
}

// If 'Options' empty fields
// 	- Censored: false;
//
// 'Options' settings
//	- Censored: true, false;
//
// Set to false to allow hentai, yaoi and yuri.
func (c *Configuration) SearchCalendar(r Result) ([]api.Calendar, error) {
  var ca []api.Calendar

  resp, err := client.Do(c.NewGetRequest(str.ConvertCalendar(r.OptionsCalendar())))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return nil, err
  }

  return f, nil
}

// Id: user id.
func (c *Configuration) AddFriend(id int) (api.FriendRequest, error) {
  var f api.FriendRequest

  resp, err := client.Do(c.NewPostRequest(str.ConvertFriend(id)))
  if err != nil {
    return f, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return f, err
  }

  return f, nil
}

// Id: user id.
func (c *Configuration) RemoveFriend(id int) (api.FriendRequest, error) {
  var f api.FriendRequest

  resp, err := client.Do(c.NewDeleteRequest(str.ConvertFriend(id)))
  if err != nil {
    return f, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return f, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return f, err
  }

  return f, nil
}

// Id: user id.
func (c *Configuration) UserUnreadMessages(id int) (api.UnreadMessages, error) {
  var um api.UnreadMessages

  resp, err := client.Do(c.NewGetRequest(str.ConvertUser(id, "unread_messages")))
  if err != nil {
    return um, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return um, err
  }

  if err := json.Unmarshal(data, &um); err != nil {
    return um, err
  }

  return um, nil
}

// Id: user id.
//
// If 'Options' empty fields
// 	- Type: news;
// 	- Page: 1;
// 	- Limit: 1;
//
// 'Options' settings
// 	- Page: 100000 maximum;
// 	- Limit: 100 maximum;
//  - Type: inbox, private, sent, news, notifications;
func (c *Configuration) UserMessages(id int, r Result) ([]api.Messages, error) {
  var m []api.Messages

  resp, err := client.Do(c.NewGetRequest(str.ConvertMessages(id, r.OptionsMessages())))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
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

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, err
  }

  return m, nil
}

// Name: people name.
//
// If 'Options' empty fields
//	- Kind: seyu;
//
// 'Options' settings
//	- Page/Limit: not supported, idk why;
//	- Kind: seyu, mangaka, producer;
func (c *Configuration) SearchPeople(name string, r Result) ([]api.AllPeople, error) {
  var ap []api.AllPeople

  resp, err := client.Do(
    c.NewGetRequest("people/search?search=" + url.QueryEscape(name) + "&" + r.OptionsPeople()),
  )
  if err != nil {
    return nil, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &ap); err != nil {
    return nil, err
  }

  return ap, nil
}

// *Configuraiton.FastIdPeople(name string).People()
func (f *FastId) People() (api.People, error) {
  var p api.People

  resp, err := client.Do(f.Conf.NewGetRequest(str.ConvertPeople(f.Id)))
  if err != nil {
    return p, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return p, err
  }

  if err := json.Unmarshal(data, &p); err != nil {
    return p, err
  }

  return p, nil
}

// Linked_type: Anime, Manga, Ranobe, Person, Character.
//
// Kind(required when Linked_type is Person): common, seyu, mangaka, producer, person.
//
// *Configuraiton.FastIdAnime/FastIdManga(name string).FavoritesCreate(linked_type string, kind string)
func (f *FastId) FavoritesCreate(linked_type string, kind string) (api.Favorites, error) {
  var fa api.Favorites

  type_map := map[string]int8{"Anime": 1, "Manga": 2, "Ranobe": 3, "Person": 4, "Character": 5}
  _, ok_type := type_map[linked_type]
  if !ok_type { return fa, errors.New("incorrect string, try again and watch the upper case") }

  kind_map := map[string]int8{"common": 1, "seyu": 2, "mangaka": 3, "producer": 4, "person": 5}
  _, ok_kind := kind_map[kind]
  if !ok_kind { kind = "" }

  resp, err := client.Do(f.Conf.NewPostRequest(str.ConvertFavorites(linked_type, f.Id, kind)))
  if err != nil {
    return fa, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return fa, err
  }

  if err := json.Unmarshal(data, &fa); err != nil {
    return fa, err
  }

  return fa, nil
}

// Linked_type: Anime, Manga, Ranobe, Person, Character.
//
// *Configuraiton.FastIdAnime/FastIdManga(name string).FavoritesDelete(linked_type string)
func (f *FastId) FavoritesDelete(linked_type string) (api.Favorites, error) {
  var ff api.Favorites

  type_map := map[string]int8{"Anime": 1, "Manga": 2, "Ranobe": 3, "Person": 4, "Character": 5}
  _, ok_type := type_map[linked_type]
  if !ok_type { return ff, errors.New("incorrect string, try again and watch the upper case") }

  resp, err := client.Do(f.Conf.NewDeleteRequest(str.ConvertFavorites(linked_type, f.Id, "")))
  if err != nil {
    return ff, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ff, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, err
  }

  return ff, nil
}

// FIXME At the moment this function does not work. I get an error code - 422.
//
// Id: anime/manga/ranobe/person/character id.
//
// Position: a new position on the list, it starts from 0.
//
// You can only get a StatusCode.
//
// *Configuraiton.FastIdAnime/FastIdManga(name string).FavoritesReorder(position int)
func (f *FastId) FavoritesReorder(position int) (int, error) {
  resp, err := client.Do(
    f.Conf.NewCustomPostRequest(str.ConvertFavoritesReorder(f.Id), position),
  )
  if err != nil {
    return 500, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// Id: user id.
func (c *Configuration) AddIgnoreUser(id int) (int, api.Ignore, error) {
  var i api.Ignore

  resp, err := client.Do(c.NewPostRequest(str.ConvertIgnoreUser(id)))
  if err != nil {
    return 500, i, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return 500, i, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return 500, i, err
  }

  return resp.StatusCode, i, nil
}

// Id: user id.
func (c *Configuration) RemoveIgnoreUser(id int) (int, api.Ignore, error) {
  var i api.Ignore

  resp, err := client.Do(c.NewDeleteRequest(str.ConvertIgnoreUser(id)))
  if err != nil {
    return 500, i, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return 500, i, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return 500, i, err
  }

  return resp.StatusCode, i, nil
}

func (c *Configuration) Dialogs() ([]api.Dialogs, error) {
  var d []api.Dialogs

  resp, err := client.Do(c.NewGetRequest("dialogs"))
  if err != nil {
    return nil, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &d); err != nil {
    return nil, err
  }

  return d, nil
}

// Id: user id.
func (c *Configuration) SearchDialogs(id int) ([]api.SearchDialogs, error) {
  var sd []api.SearchDialogs

  resp, err := client.Do(c.NewGetRequest(str.ConvertDialogs(id)))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &sd); err != nil {
    return nil, err
  }

  return sd, nil
}

// Id: user id.
func (c *Configuration) DeleteDialogs(id int) (int, api.FriendRequest, error) {
  var fr api.FriendRequest

  resp, err := client.Do(c.NewDeleteRequest(str.ConvertDialogs(id)))
  if err != nil {
    return 500, fr, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return 500, fr, err
  }

  if err := json.Unmarshal(data, &fr); err != nil {
    return resp.StatusCode, fr, errors.New("не найдено ни одного сообщения для удаления")
  }

  return resp.StatusCode, fr, nil
}
