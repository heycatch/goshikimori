// Copyright (C) 2024 heycatch <andreyisback@yandex.ru>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Comments are made in the style of "godoc" syntax support.
//
// More information can be found in the [examples] folder.
//
// [examples]: https://github.com/heycatch/goshikimori/blob/master/examples/
package goshikimori

import (
  "net/http"
  "io"
  "encoding/json"
  "errors"
  "strconv"

  "github.com/heycatch/goshikimori/api"
  "github.com/heycatch/goshikimori/concat"
)

// Name: user name.
//
// Search by user is case sensitive.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (c *Configuration) SearchUser(name string) (api.Users, int, error) {
  var u api.Users
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 6(users/) + ?(name)
    concat.Url(32+len(name), []string{SITE, "users/", name}), 10,
  )
  if err != nil {
    return u, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return u, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return u, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return u, resp.StatusCode, err
  }

  return u, resp.StatusCode, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Name: user name.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// Don't use Stats.Statuses.Anime and Stats.Statuses.Manga: empty slice.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/users
func (c *Configuration) SearchUsers(name string, r Result) ([]api.Users, int, error) {
  var u []api.Users
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 100)

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 13(users?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(40+len(name)+len(opt), []string{
      SITE, "users?search=", name, "&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return nil, resp.StatusCode, err
  }

  return u, resp.StatusCode, nil
}

// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserFriends(r Result) ([]api.UserFriends, error) {
  var uf []api.UserFriends
  var client = &http.Client{}

  // 26(SITE) + 6(users/) + ?(id) + 9(/friends?) + ?(Result)
  opt := r.OptionsOnlyPageLimit(100000, 100)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 9(/friends?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/friends?" + opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserClubs() ([]api.Clubs, error) {
  var uc []api.Clubs
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 6(/clubs)
    concat.Url(38+len(str_id), []string{
      SITE, "users/", str_id, "/clubs",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Limit always returns +1 of the given number.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 5000 maximum;
//  - Censored: true, false;
//
//  - Status:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserAnimeRates(r Result) ([]api.UserAnimeRates, error) {
  var ar []api.UserAnimeRates
  var client = &http.Client{}

  opt := r.OptionsAnimeRates()
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 13(/anime_rates?) + ?(Result)
    concat.Url(45+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/anime_rates?" + opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Limit always returns +1 of the given number.
//
// 'Options' Settings:
//  - Page: 100000 maximum;
//  - Limit: 5000 maximum;
//  - Censored: true, false;
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserMangaRates(r Result) ([]api.UserMangaRates, error) {
  var mr []api.UserMangaRates
  var client = &http.Client{}

  opt := r.OptionsMangaRates()
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 13(/manga_rates?) + ?(Result)
    concat.Url(45+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/manga_rates?" + opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserFavourites() (api.UserFavourites, error) {
  var uf api.UserFavourites
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 11(/favourites)
    concat.Url(43+len(str_id), []string{
      SITE, "users/", str_id, "/favourites",
    }), 10,
  )
  if err != nil {
    return uf, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Limit always returns +1 of the given number.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//  - Target_id: id anime/manga/ranobe;
//
//  - Target_type:
//
//  > TARGET_TYPE_ANIME, TARGET_TYPE_MANGA;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserHistory(r Result) ([]api.UserHistory, error) {
  var uh []api.UserHistory
  var client = &http.Client{}

  opt := r.OptionsUserHistory()
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 9(/history?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/history?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserBans() ([]api.Bans, error) {
  var b []api.Bans
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 5(/bans)
    concat.Url(37+len(str_id), []string{
      SITE, "users/", str_id, "/bans",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/whoami
func (c *Configuration) WhoAmi() (api.Who, int, error) {
  var w api.Who
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 12(users/whoami)
    concat.Url(38, []string{
      SITE, "users/whoami",
    }), 10,
  )
  if err != nil {
    return w, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return w, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return w, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &w); err != nil {
    return w, resp.StatusCode, err
  }

  return w, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchAnime() (api.Anime, error) {
  var a api.Anime
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "animes/", str_id,
    }), 10,
  )
  if err != nil {
    return a, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return a, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return a, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return a, err
  }

  return a, nil
}

// Name: anime name.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 50 maximum;
//
//  - Order:
//
//  > ANIME_ORDER_ID, ANIME_ORDER_RANKED, ANIME_ORDER_KIND,
//  ANIME_ORDER_POPULARITY, ANIME_ORDER_NAME, ANIME_ORDER_AIRED_ON,
//  ANIME_ORDER_EPISODES, ANIME_ORDER_STATUS;
//
//  - Kind:
//
//  > ANIME_KIND_TV, ANIME_KIND_MOVIE, ANIME_KIND_OVA, ANIME_KIND_ONA,
//  ANIME_KIND_SPECIAL, ANIME_KIND_MUSIC, ANIME_KIND_TV_13, ANIME_KIND_TV_24,
//  ANIME_KIND_TV_48, ANIME_KIND_TV_NOT_EQUAL, ANIME_KIND_MOVIE_NOT_EQUAL,
//  ANIME_KIND_OVA_NOT_EQUAL, ANIME_KIND_ONA_NOT_EQUAL, ANIME_KIND_SPECIAL_NOT_EQUAL,
//  ANIME_KIND_MUSIC_NOT_EQUAL, ANIME_KIND_TV_13_NOT_EQUAL,
//  ANIME_KIND_TV_24_NOT_EQUAL, ANIME_KIND_TV_48_NOT_EQUAL;
//
//  - Status:
//
//  > ANIME_STATUS_ANONS, ANIME_STATUS_ONGOING, ANIME_STATUS_RELEASED,
//  ANIME_STATUS_ANONS_NOT_EQUAL, ANIME_STATUS_ONGOING_NOT_EQUAL,
//  ANIME_STATUS_RELEASED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Duration:
//
//  > ANIME_DURATION_S, ANIME_DURATION_D, ANIME_DURATION_F,
//  ANIME_DURATION_S_NOT_EQUAL, ANIME_DURATION_D_NOT_EQUAL,
//  ANIME_DURATION_F_NOT_EQUAL;
//
//  - Rating:
//
//  > ANIME_RATING_NONE, ANIME_RATING_G, ANIME_RATING_PG,
//  ANIME_RATING_PG_13, ANIME_RATING_R, ANIME_RATING_R_PLUS, ANIME_RATING_RX,
//  ANIME_RATING_G_NOT_EQUAL, ANIME_RATING_PG_NOT_EQUAL,
//  ANIME_RATING_PG_13_NOT_EQUAL, ANIME_RATING_R_NOT_EQUAL,
//  ANIME_RATING_R_PLUS_NOT_EQUAL, ANIME_RATING_RX_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Genre_v2: id search. Below is a list of all available genres by id:
//
//  > 1 (Action); 2 (Adventure); 3 (Cars); 4 (Comedy); 5 (Dementia); 6 (Demons); 7 (Mystery);
//  8 (Drama); 9 (Ecchi); 10 (Fantasy); 11 (Game); 12 (Hentai); 13 (Historical); 14 (Horror);
//  15 (Kids); 16 (Magic); 17 (Martial Arts); 18 (Mecha); 19 (Music); 20 (Parody); 21 (Samurai);
//  22 (Romance); 23 (School); 24 (Sci-Fi); 25 (Shoujo); 26 (Shoujo Ai); 27 (Shounen); 28 (Shounen Ai);
//  29 (Space); 30 (Sports); 31 (Super Power); 32 (Vampire); 33 (Yaoi); 34 (Yuri); 35 (Harem);
//  36 (Slice of Life); 37 (Supernatural); 38 (Military); 39 (Police); 40 (Psychological);
//  41 (Thriller); 42 (Seinen); 43 (Josei); 539 (Erotica); 541 (Work Life); 543 (Gourmet);
//
// Set to true to discard hentai, yaoi and yuri.
//
//  - Type: "deprecated";
//  - Studio: not supported;
//  - Franchise: not supported;
//  - Ids: not supported;
//  - Exclude_ids: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchAnimes(name string, r Result) ([]api.Animes, int, error) {
  var a []api.Animes
  var client = &http.Client{}

  opt := r.OptionsAnime()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(animes?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(41+len(name)+len(opt), []string{
      SITE, "animes?search=", name, "&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, resp.StatusCode, err
  }

  return a, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchManga() (api.Manga, error) {
  var m api.Manga
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "mangas/", str_id,
    }), 10,
  )
  if err != nil {
    return m, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return m, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return m, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, err
  }

  return m, nil
}

// Name: manga name.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 50 maximum;
//
//  - Order:
//
//  > MANGA_ORDER_ID, MANGA_ORDER_RANKED, MANGA_ORDER_KIND, MANGA_ORDER_POPULARITY,
//  MANGA_ORDER_NAME, MANGA_ORDER_AIRED_ON, MANGA_ORDER_VOLUMES,
//  MANGA_ORDER_CHAPTERS, MANGA_ORDER_STATUS;
//
//  - Kind:
//
//  > MANGA_KIND_MANGA, MANGA_KIND_MANHWA, MANGA_KIND_MANHUA, MANGA_KIND_LIGHT_NOVEL,
//  MANGA_KIND_NOVEL, MANGA_KIND_ONE_SHOT, MANGA_KIND_DOUJIN, MANGA_KIND_MANGA_NOT_EQUAL,
//  MANGA_KIND_MANHWA_NOT_EQUAL, MANGA_KIND_MANHUA_NOT_EQUAL, MANGA_KIND_LIGHT_NOVEL_NOT_EQUAL,
//  MANGA_KIND_NOVEL_NOT_EQUAL, MANGA_KIND_ONE_SHOT_NOT_EQUAL, MANGA_KIND_DOUJIN_NOT_EQUAL;
//
//  - Status:
//
//  > MANGA_STATUS_ANONS, MANGA_STATUS_ONGOING, MANGA_STATUS_RELEASED, MANGA_STATUS_PAUSED,
//  MANGA_STATUS_DISCONTINUED, MANGA_STATUS_ANONS_NOT_EQUAL, MANGA_STATUS_ONGOING_NOT_EQUAL,
//  MANGA_STATUS_RELEASED_NOT_EQUAL, MANGA_STATUS_PAUSED_NOT_EQUAL, MANGA_STATUS_DISCONTINUED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Genre_v2: id search. Below is a list of all available genres by id:
//
//  > 46 (Mystery); 47 (Shounen); 48 (Supernatural);
//  49 (Comedy); 50 (Drama); 51 (Ecchi); 52 (Seinen); 53 (Sci-Fi); 54 (Slice of Life); 55 (Shounen Ai);
//  56 (Action); 57 (Fantasy); 58 (Magic); 59 (Hentai); 60 (School); 61 (Doujinshi); 62 (Romance);
//  63 (Shoujo); 64 (Vampire); 65 (Yaoi); 66 (Martial Arts); 67 (Psychological); 68 (Adventure);
//  69 (Historical); 70 (Military); 71 (Harem); 72 (Demons); 73 (Shoujo Ai); 74 (Gender Bender);
//  75 (Yuri); 76 (Sports); 77 (Kids); 78 (Music); 79 (Game); 80 (Horror); 81 (Thriller);
//  82 (Super Power); 83 (Mecha); 84 (Cars); 85 (Space); 86 (Parody); 87 (Josei); 88 (Samurai);
//  89 (Police); 90 (Dementia); 540 (Erotica); 542 (Work Life); 544 (Gourmet);
//
// Set to true to discard hentai, yaoi and yuri.
//
//  - Type: "deprecated";
//  - Publisher: not supported;
//  - Franchise: not supported;
//  - Ids: not supported;
//  - Exclude_ids: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchMangas(name string, r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  opt := r.OptionsManga()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(mangas?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(41+len(name)+len(opt), []string{
      SITE, "mangas?search=", name, "&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchRanobe() (api.Manga, error) {
  var m api.Manga
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "ranobe/", str_id,
    }), 10,
  )
  if err != nil {
    return m, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return m, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return m, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, err
  }

  return m, nil
}

// Name: ranobe name.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 50 maximum;
//
//  - Order:
//
//  > MANGA_ORDER_ID, MANGA_ORDER_RANKED, MANGA_ORDER_KIND, MANGA_ORDER_POPULARITY,
//  MANGA_ORDER_NAME, MANGA_ORDER_AIRED_ON, MANGA_ORDER_VOLUMES,
//  MANGA_ORDER_CHAPTERS, MANGA_ORDER_STATUS;
//
//  - Status:
//
//  > MANGA_STATUS_ANONS, MANGA_STATUS_ONGOING, MANGA_STATUS_RELEASED, MANGA_STATUS_PAUSED,
//  MANGA_STATUS_DISCONTINUED, MANGA_STATUS_ANONS_NOT_EQUAL, MANGA_STATUS_ONGOING_NOT_EQUAL,
//  MANGA_STATUS_RELEASED_NOT_EQUAL, MANGA_STATUS_PAUSED_NOT_EQUAL, MANGA_STATUS_DISCONTINUED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Genre_v2: id search. Below is a list of all available genres by id:
//
//  > 46 (Mystery); 47 (Shounen); 48 (Supernatural);
//  49 (Comedy); 50 (Drama); 51 (Ecchi); 52 (Seinen); 53 (Sci-Fi); 54 (Slice of Life); 55 (Shounen Ai);
//  56 (Action); 57 (Fantasy); 58 (Magic); 59 (Hentai); 60 (School); 61 (Doujinshi); 62 (Romance);
//  63 (Shoujo); 64 (Vampire); 65 (Yaoi); 66 (Martial Arts); 67 (Psychological); 68 (Adventure);
//  69 (Historical); 70 (Military); 71 (Harem); 72 (Demons); 73 (Shoujo Ai); 74 (Gender Bender);
//  75 (Yuri); 76 (Sports); 77 (Kids); 78 (Music); 79 (Game); 80 (Horror); 81 (Thriller);
//  82 (Super Power); 83 (Mecha); 84 (Cars); 85 (Space); 86 (Parody); 87 (Josei); 88 (Samurai);
//  89 (Police); 90 (Dementia); 540 (Erotica); 542 (Work Life); 544 (Gourmet);
//
// Set to true to discard hentai, yaoi and yuri.
//
//  - Publisher: not supported;
//  - Franchise: not supported;
//  - Ids: not supported;
//  - Exclude_ids: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchRanobes(name string, r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  opt := r.OptionsRanobe()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(ranobe?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(41+len(name)+len(opt), []string{
      SITE, "ranobe?search=", name, "&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// Name: user name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdUser(name string) (*FastId, int, error) {
  var u api.Users
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 6(users/) + ?(name)
    concat.Url(32+len(name), []string{SITE, "users/", name,}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  return &FastId{Id: u.Id, Conf: *c, Err: err}, resp.StatusCode, nil
}

// Name: anime name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdAnime(name string) (*FastId, int, error) {
  var a []api.Animes
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(animes?search=) + ?(name)
    concat.Url(40+len(name), []string{SITE, "animes?search=", name}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(a) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, resp.StatusCode, nil }

  return &FastId{Id: a[0].Id, Conf: *c, Err: nil}, resp.StatusCode, nil
}

// Name: manga name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdManga(name string) (*FastId, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(mangas?search=) + ?(name)
    concat.Url(40+len(name), []string{SITE, "mangas?search=", name}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(m) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, resp.StatusCode, nil }

  return &FastId{Id: m[0].Id, Conf: *c, Err: nil}, resp.StatusCode, nil
}

// Name: ranobe name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdRanobe(name string) (*FastId, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(ranobe?search=) + ?(name)
    concat.Url(40+len(name), []string{SITE, "ranobe?search=", name}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(m) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, resp.StatusCode, nil }

  return &FastId{Id: m[0].Id, Conf: *c, Err: nil}, resp.StatusCode, nil
}

// Name: club name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdClub(name string) (*FastId, int, error) {
  var cl []api.Clubs
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 13(clubs?search=) + ?(name)
    concat.Url(39+len(name), []string{SITE, "clubs?search=", name}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(cl) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, resp.StatusCode, nil }

  return &FastId{Id: cl[0].Id, Conf: *c, Err: nil}, resp.StatusCode, nil
}

// Name: character name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdCharacter(name string) (*FastId, int, error) {
  var ch []api.CharacterInfo
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 25(characters/search?search=) + ?(name)
    concat.Url(51+len(name), []string{
      SITE, "characters/search?search=", name,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ch); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(ch) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, resp.StatusCode, nil }

  return &FastId{Id: ch[0].Id, Conf: *c, Err: nil}, resp.StatusCode, nil
}

// Name: people name.
//
// NOTES: There is a conflict with a long word in Latin. Everything is fine in Cyrillic.
// At the moment the problem has been solved by an additional check for unicode.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdPeople(name string) (*FastId, int, error) {
  var ap []api.AllPeople
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 21(people/search?search=) + ?(name)
    concat.Url(47+len(name), []string{
      SITE, "people/search?search=", languageCheck(name),
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ap); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, resp.StatusCode, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(ap) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, resp.StatusCode, nil }

  return &FastId{Id: ap[0].Id, Conf: *c, Err: nil}, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_screenshots
func (f *FastId) SearchAnimeScreenshots() ([]api.AnimeScreenshots, error) {
  var s []api.AnimeScreenshots
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 12(/screenshots)
    concat.Url(45+len(str_id), []string{
      SITE, "animes/", str_id, "/screenshots",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchAnimeFranchise() (api.Franchise, error) {
  var ff api.Franchise
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 10(/franchise)
    concat.Url(43+len(str_id), []string{
      SITE, "animes/", str_id, "/franchise",
    }), 10,
  )
  if err != nil {
    return ff, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchMangaFranchise() (api.Franchise, error) {
  var ff api.Franchise
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id) + 10(/franchise)
    concat.Url(43+len(str_id), []string{
      SITE, "mangas/", str_id, "/franchise",
    }), 10,
  )
  if err != nil {
    return ff, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchRanobeFranchise() (api.Franchise, error) {
  var ff api.Franchise
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id) + 10(/franchise)
    concat.Url(43+len(str_id), []string{
      SITE, "ranobe/", str_id, "/franchise",
    }), 10,
  )
  if err != nil {
    return ff, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchAnimeExternalLinks() ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 15(/external_links)
    concat.Url(48+len(str_id), []string{
      SITE, "animes/", str_id, "/external_links",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchMangaExternalLinks() ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id) + 15(/external_links)
    concat.Url(48+len(str_id), []string{
      SITE, "mangas/", str_id, "/external_links",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchRanobeExternalLinks() ([]api.ExternalLinks, error) {
  var el []api.ExternalLinks
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id) + 15(/external_links)
    concat.Url(48+len(str_id), []string{
      SITE, "ranobe/", str_id, "/external_links",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarAnime() ([]api.Animes, error) {
  var a []api.Animes
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 8(/similar)
    concat.Url(41+len(str_id), []string{
      SITE, "animes/", str_id, "/similar",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarManga() ([]api.Mangas, error) {
  var m []api.Mangas
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id) + 8(/similar)
    concat.Url(41+len(str_id), []string{
      SITE, "mangas/", str_id, "/similar",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarRanobe() ([]api.Mangas, error) {
  var m []api.Mangas
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id) + 8(/similar)
    concat.Url(41+len(str_id), []string{
      SITE, "ranobe/", str_id, "/similar",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedAnime() ([]api.RelatedAnimes, error) {
  var a []api.RelatedAnimes
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 8(/related)
    concat.Url(41+len(str_id), []string{
      SITE, "animes/", str_id, "/related",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedManga() ([]api.RelatedMangas, error) {
  var m []api.RelatedMangas
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id) + 8(/related)
    concat.Url(41+len(str_id), []string{
      SITE, "mangas/", str_id, "/related",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedRanobe() ([]api.RelatedMangas, error) {
  var m []api.RelatedMangas
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id) + 8(/related)
    concat.Url(41+len(str_id), []string{
      SITE, "ranobe/", str_id, "/related",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Limit always returns +1 of the given number.
//
// Name: club name.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (c *Configuration) SearchClubs(name string, r Result) ([]api.Clubs, int, error) {
  var cl []api.Clubs
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 30)

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 13(clubs?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(40+len(name)+len(opt), []string{
      SITE, "clubs?search=", name, "&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return nil, resp.StatusCode, err
  }

  return cl, resp.StatusCode, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubAnimes(r Result) ([]api.Animes, error) {
  var a []api.Animes
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/animes?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/animes?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: The limit does not work and always gives the maximum amount.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMangas(r Result) ([]api.Mangas, error) {
  var m []api.Mangas
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/mangas?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/mangas?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: The limit does not work and always gives the maximum amount.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubRanobe(r Result) ([]api.Mangas, error) {
  var m []api.Mangas
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/ranobe?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/ranobe?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: The limit does not work and always gives the maximum amount.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCharacters(r Result) ([]api.CharacterInfo, error) {
  var ci []api.CharacterInfo
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 12(/characters?) + ?(Result)
    concat.Url(44+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/characters?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: The limit does not work and always gives the maximum amount.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubClubs(r Result) ([]api.Clubs, error) {
  var cc []api.Clubs
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 7(/clubs?) + ?(Result)
    concat.Url(39+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/clubs?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: The limit does not work and always gives the maximum amount.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Page: 4 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCollections(r Result) ([]api.ClubCollections, error) {
  var cc []api.ClubCollections
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 4)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 13(/collections?) + ?(Result)
    concat.Url(45+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/collections?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Limit always returns +1 of the given number.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMembers(r Result) ([]api.UserFriends, error) {
  var uf []api.UserFriends
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 100)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 9(/members?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/members?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Limit always returns +1 of the given number.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubImages(r Result) ([]api.ClubImages, error) {
  var cm []api.ClubImages
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 100)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/images?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/images?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) ClubJoin() (int, error) {
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  post, cancel, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 5(/join)
    concat.Url(37+len(str_id), []string{
      SITE, "clubs/", str_id, "/join",
    }), 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) ClubLeave() (int, error) {
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  post, cancel, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 6(/leave)
    concat.Url(38+len(str_id), []string{
      SITE, "clubs/", str_id, "/leave",
    }), 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// As a result, we return a complete list of all achievements.
//
// Next comes the filtering through "NekoSearch" and the error about obtaining
// specific achievements is already being processed there.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/achievements
func (f *FastId) SearchAchievement() ([]api.Achievements, error) {
  var a []api.Achievements
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 21(achievements?user_id=) + ?(id)
    concat.Url(47+len(str_id), []string{
      SITE, "achievements?user_id=", str_id,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/video
func (f *FastId) SearchAnimeVideos() ([]api.AnimeVideos, error) {
  var v []api.AnimeVideos
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    concat.Url(40+len(str_id), []string{
      SITE, "animes/", str_id, "/videos",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/roles
func (f *FastId) SearchAnimeRoles() ([]api.Roles, error) {
  var r []api.Roles
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 6(/roles)
    concat.Url(39+len(str_id), []string{
      SITE, "animes/", str_id, "/roles",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/roles
func (f *FastId) SearchMangaRoles() ([]api.Roles, error) {
  var r []api.Roles
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id) + 6(/roles)
    concat.Url(39+len(str_id), []string{
      SITE, "mangas/", str_id, "/roles",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/roles
func (f *FastId) SearchRanobeRoles() ([]api.Roles, error) {
  var r []api.Roles
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id) + 6(/roles)
    concat.Url(39+len(str_id), []string{
      SITE, "ranobe/", str_id, "/roles",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/bans
func (c *Configuration) SearchBans() ([]api.Bans, int, error) {
  var b []api.Bans
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 4(bans)
    concat.Url(30, []string{SITE, "bans"}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &b); err != nil {
    return nil, resp.StatusCode, err
  }

  return b, resp.StatusCode, nil
}

// 'Options' settings:
//  - Censored: true, false;
//
// Set to false to allow hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/calendar
func (c *Configuration) SearchCalendar(r Result) ([]api.Calendar, int, error) {
  var ca []api.Calendar
  var client = &http.Client{}

  opt := r.OptionsCalendar()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(calendar?) + ?(Result)
    concat.Url(35+len(opt), []string{
      SITE, "calendar?", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ca); err != nil {
    return nil, resp.StatusCode, err
  }

  return ca, resp.StatusCode, nil
}

// name: GENRES_ANIME or GENRES_MANGA.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/genres
func (c *Configuration) SearchGenres(name string) ([]api.Genres, int, error) {
  var g []api.Genres
  var client = &http.Client{}


  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 12(genres?kind=) + ?(name)
    concat.Url(38+len(name), []string{SITE, "genres?kind=", name}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &g); err != nil {
    return nil, resp.StatusCode, err
  }

  return g, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/studios
func (c *Configuration) SearchStudios() ([]api.Studios, int, error) {
  var s []api.Studios
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 7(studios)
    concat.Url(33, []string{SITE, "studios"}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &s); err != nil {
    return nil, resp.StatusCode, err
  }

  return s, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/publishers
func (c *Configuration) SearchPublishers() ([]api.Publishers, int, error) {
  var p []api.Publishers
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 10(publishers)
    concat.Url(36, []string{SITE, "publishers"}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &p); err != nil {
    return nil, resp.StatusCode, err
  }

  return p, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/forums
func (c *Configuration) SearchForums() ([]api.Forums, int, error) {
  var f []api.Forums
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 6(forums)
    concat.Url(32, []string{SITE, "forums"}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return nil, resp.StatusCode, err
  }

  return f, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/add_remove_friend
func (f *FastId) AddFriend() (api.FriendRequest, error) {
  var ff api.FriendRequest
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  post, cancel, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(friends/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "friends/", str_id,
    }), 10,
  )
  if err != nil {
    return ff, err
  }
  defer cancel()

  resp, err := client.Do(post)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/add_remove_friend
func (f *FastId) RemoveFriend() (api.FriendRequest, error) {
  var ff api.FriendRequest
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  remove, cancel, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(friends/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "friends/", str_id,
    }), 10,
  )
  if err != nil {
    return ff, err
  }
  defer cancel()

  resp, err := client.Do(remove)
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

// Show current user unread messages counts.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/messages
func (f *FastId) UserUnreadMessages() (api.UnreadMessages, error) {
  var um api.UnreadMessages
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 16(/unread_messages)
    concat.Url(48+len(str_id), []string{
      SITE, "users/", str_id, "/unread_messages",
    }), 10,
  )
  if err != nil {
    return um, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//  - Type:
//
//  > MESSAGE_TYPE_INBOX, MESSAGE_TYPE_PRIVATE, MESSAGE_TYPE_SENT,
//  MESSAGE_TYPE_NEWS, MESSAGE_TYPE_NOTIFICATIONS;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/messages
func (f *FastId) UserMessages(r Result) ([]api.Messages, error) {
  var m []api.Messages
  var client = &http.Client{}

  opt := r.OptionsMessages()
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 10(/messages?) + ?(Result)
    concat.Url(42+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/messages?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsAnime() (api.Constants, int, error) {
  var ca api.Constants
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 15(constants/anime)
    concat.Url(41, []string{SITE, "constants/anime"}), 10,
  )
  if err != nil {
    return ca, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return ca, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ca, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ca); err != nil {
    return ca, resp.StatusCode, err
  }

  return ca, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsManga() (api.Constants, int, error) {
  var cm api.Constants
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 15(constants/manga)
    concat.Url(41, []string{SITE, "constants/manga"}), 10,
  )
  if err != nil {
    return cm, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return cm, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return cm, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &cm); err != nil {
    return cm, resp.StatusCode, err
  }

  return cm, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsUserRate() (api.ConstantsUserRate, int, error) {
  var ur api.ConstantsUserRate
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 19(constants/user_rate)
    concat.Url(45, []string{SITE, "constants/user_rate"}), 10,
  )
  if err != nil {
    return ur, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return ur, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ur, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ur); err != nil {
    return ur, resp.StatusCode, err
  }

  return ur, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsClub() (api.ConstantsClub, int, error) {
  var cc api.ConstantsClub
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(constants/club)
    concat.Url(40, []string{SITE, "constants/club"}), 10,
  )
  if err != nil {
    return cc, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return cc, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return cc, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return cc, resp.StatusCode, err
  }

  return cc, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsSmileys() ([]api.ConstantsSmileys, int, error) {
  var cs []api.ConstantsSmileys
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 17(constants/smileys)
    concat.Url(43, []string{SITE, "constants/smileys"}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &cs); err != nil {
    return nil, resp.StatusCode, err
  }

  return cs, resp.StatusCode, nil
}

// 'Options' settings:
//  - Limit: 50 maximum;
//  - Kind:
//
//  > ANIME_KIND_TV, ANIME_KIND_MOVIE, ANIME_KIND_OVA, ANIME_KIND_ONA,
//  ANIME_KIND_SPECIAL, ANIME_KIND_MUSIC, ANIME_KIND_TV_13, ANIME_KIND_TV_24,
//  ANIME_KIND_TV_48, ANIME_KIND_TV_NOT_EQUAL, ANIME_KIND_MOVIE_NOT_EQUAL,
//  ANIME_KIND_OVA_NOT_EQUAL, ANIME_KIND_ONA_NOT_EQUAL, ANIME_KIND_SPECIAL_NOT_EQUAL,
//  ANIME_KIND_MUSIC_NOT_EQUAL, ANIME_KIND_TV_13_NOT_EQUAL,
//  ANIME_KIND_TV_24_NOT_EQUAL, ANIME_KIND_TV_48_NOT_EQUAL;
//
//  - Status:
//
//  > ANIME_STATUS_ANONS, ANIME_STATUS_ONGOING, ANIME_STATUS_RELEASED,
//  ANIME_STATUS_ANONS_NOT_EQUAL, ANIME_STATUS_ONGOING_NOT_EQUAL,
//  ANIME_STATUS_RELEASED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Duration:
//
//  > ANIME_DURATION_S, ANIME_DURATION_D, ANIME_DURATION_F,
//  ANIME_DURATION_S_NOT_EQUAL, ANIME_DURATION_D_NOT_EQUAL,
//  ANIME_DURATION_F_NOT_EQUAL;
//
//  - Rating:
//
//  > ANIME_RATING_NONE, ANIME_RATING_G, ANIME_RATING_PG,
//  ANIME_RATING_PG_13, ANIME_RATING_R, ANIME_RATING_R_PLUS, ANIME_RATING_RX,
//  ANIME_RATING_G_NOT_EQUAL, ANIME_RATING_PG_NOT_EQUAL,
//  ANIME_RATING_PG_13_NOT_EQUAL, ANIME_RATING_R_NOT_EQUAL,
//  ANIME_RATING_R_PLUS_NOT_EQUAL, ANIME_RATING_RX_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Genre_v2: id search. Below is a list of all available genres by id:
//
//  > 1 (Action); 2 (Adventure); 3 (Cars); 4 (Comedy); 5 (Dementia); 6 (Demons); 7 (Mystery);
//  8 (Drama); 9 (Ecchi); 10 (Fantasy); 11 (Game); 12 (Hentai); 13 (Historical); 14 (Horror);
//  15 (Kids); 16 (Magic); 17 (Martial Arts); 18 (Mecha); 19 (Music); 20 (Parody); 21 (Samurai);
//  22 (Romance); 23 (School); 24 (Sci-Fi); 25 (Shoujo); 26 (Shoujo Ai); 27 (Shounen); 28 (Shounen Ai);
//  29 (Space); 30 (Sports); 31 (Super Power); 32 (Vampire); 33 (Yaoi); 34 (Yuri); 35 (Harem);
//  36 (Slice of Life); 37 (Supernatural); 38 (Military); 39 (Police); 40 (Psychological);
//  41 (Thriller); 42 (Seinen); 43 (Josei); 539 (Erotica); 541 (Work Life); 543 (Gourmet);
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/random
func (c *Configuration) RandomAnimes(r Result) ([]api.Animes, int, error) {
  var a []api.Animes
  var client = &http.Client{}

  opt := r.OptionsRandomAnime()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 20(animes?order=random&) + ?(Result)
    concat.Url(46+len(opt), []string{
      SITE, "animes?order=random&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, resp.StatusCode, err
  }

  return a, resp.StatusCode, nil
}

// 'Options' settings:
//  - Limit: 50 maximum;
//  - Kind:
//
//  > MANGA_KIND_MANGA, MANGA_KIND_MANHWA, MANGA_KIND_MANHUA, MANGA_KIND_LIGHT_NOVEL,
//  MANGA_KIND_NOVEL, MANGA_KIND_ONE_SHOT, MANGA_KIND_DOUJIN, MANGA_KIND_MANGA_NOT_EQUAL,
//  MANGA_KIND_MANHWA_NOT_EQUAL, MANGA_KIND_MANHUA_NOT_EQUAL, MANGA_KIND_LIGHT_NOVEL_NOT_EQUAL,
//  MANGA_KIND_NOVEL_NOT_EQUAL, MANGA_KIND_ONE_SHOT_NOT_EQUAL, MANGA_KIND_DOUJIN_NOT_EQUAL;
//
//  - Status:
//
//  > MANGA_STATUS_ANONS, MANGA_STATUS_ONGOING, MANGA_STATUS_RELEASED, MANGA_STATUS_PAUSED,
//  MANGA_STATUS_DISCONTINUED, MANGA_STATUS_ANONS_NOT_EQUAL, MANGA_STATUS_ONGOING_NOT_EQUAL,
//  MANGA_STATUS_RELEASED_NOT_EQUAL, MANGA_STATUS_PAUSED_NOT_EQUAL, MANGA_STATUS_DISCONTINUED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Genre_v2: id search. Below is a list of all available genres by id:
//
//  > 46 (Mystery); 47 (Shounen); 48 (Supernatural);
//  49 (Comedy); 50 (Drama); 51 (Ecchi); 52 (Seinen); 53 (Sci-Fi); 54 (Slice of Life); 55 (Shounen Ai);
//  56 (Action); 57 (Fantasy); 58 (Magic); 59 (Hentai); 60 (School); 61 (Doujinshi); 62 (Romance);
//  63 (Shoujo); 64 (Vampire); 65 (Yaoi); 66 (Martial Arts); 67 (Psychological); 68 (Adventure);
//  69 (Historical); 70 (Military); 71 (Harem); 72 (Demons); 73 (Shoujo Ai); 74 (Gender Bender);
//  75 (Yuri); 76 (Sports); 77 (Kids); 78 (Music); 79 (Game); 80 (Horror); 81 (Thriller);
//  82 (Super Power); 83 (Mecha); 84 (Cars); 85 (Space); 86 (Parody); 87 (Josei); 88 (Samurai);
//  89 (Police); 90 (Dementia); 540 (Erotica); 542 (Work Life); 544 (Gourmet);
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/random
func (c *Configuration) RandomMangas(r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  opt := r.OptionsRandomManga()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 20(mangas?order=random&) + ?(Result)
    concat.Url(46+len(opt), []string{
      SITE, "mangas?order=random&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// 'Options' settings:
//  - Limit: 50 maximum;
//  - Status:
//
//  > MANGA_STATUS_ANONS, MANGA_STATUS_ONGOING, MANGA_STATUS_RELEASED, MANGA_STATUS_PAUSED,
//  MANGA_STATUS_DISCONTINUED, MANGA_STATUS_ANONS_NOT_EQUAL, MANGA_STATUS_ONGOING_NOT_EQUAL,
//  MANGA_STATUS_RELEASED_NOT_EQUAL, MANGA_STATUS_PAUSED_NOT_EQUAL, MANGA_STATUS_DISCONTINUED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;

//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Genre_v2: id search. Below is a list of all available genres by id:
//
//  > 46 (Mystery); 47 (Shounen); 48 (Supernatural);
//  49 (Comedy); 50 (Drama); 51 (Ecchi); 52 (Seinen); 53 (Sci-Fi); 54 (Slice of Life); 55 (Shounen Ai);
//  56 (Action); 57 (Fantasy); 58 (Magic); 59 (Hentai); 60 (School); 61 (Doujinshi); 62 (Romance);
//  63 (Shoujo); 64 (Vampire); 65 (Yaoi); 66 (Martial Arts); 67 (Psychological); 68 (Adventure);
//  69 (Historical); 70 (Military); 71 (Harem); 72 (Demons); 73 (Shoujo Ai); 74 (Gender Bender);
//  75 (Yuri); 76 (Sports); 77 (Kids); 78 (Music); 79 (Game); 80 (Horror); 81 (Thriller);
//  82 (Super Power); 83 (Mecha); 84 (Cars); 85 (Space); 86 (Parody); 87 (Josei); 88 (Samurai);
//  89 (Police); 90 (Dementia); 540 (Erotica); 542 (Work Life); 544 (Gourmet);
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/random
func (c *Configuration) RandomRanobes(r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  opt := r.OptionsRandomRanobe()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 20(ranobe?order=random&) + ?(Result)
    concat.Url(46+len(opt), []string{
      SITE, "ranobe?order=random&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/character
func (f *FastId) SearchCharacter() (api.Character, error) {
  var ch api.Character
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 11(characters/) + ?(id)
    concat.Url(37+len(str_id), []string{
      SITE, "characters/", str_id,
    }), 10,
  )
  if err != nil {
    return ch, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return ch, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ch, err
  }

  if err := json.Unmarshal(data, &ch); err != nil {
    return ch, err
  }

  return ch, nil
}

// Name: character name.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/character
func (c *Configuration) SearchCharacters(name string) ([]api.CharacterInfo, int, error) {
  var ci []api.CharacterInfo
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 25(characters/search?search=) + ?(name)
    concat.Url(51+len(name), []string{SITE, "characters/search?search=", name}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ci); err != nil {
    return nil, resp.StatusCode, err
  }

  return ci, resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/people
func (f *FastId) SearchPeople() (api.People, error) {
  var p api.People
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(people/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "people/", str_id,
    }), 10,
  )
  if err != nil {
    return p, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// FIXME: Page and limit not supprted, idk why. Check later.
//
// Name: people name.
//
// 'Options' settings:
//  - Kind:
//
//  > PEOPLE_KIND_SEYU, PEOPLE_KIND_MANGAKA, PEOPLE_KIND_PRODUCER;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/people
func (c *Configuration) SearchPeoples(name string, r Result) ([]api.AllPeople, int, error) {
  var ap []api.AllPeople
  var client = &http.Client{}

  opt := r.OptionsPeople()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    concat.Url(48+len(name)+len(opt), []string{
      SITE, "people/search?search=", name, "&", opt,
    }), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ap); err != nil {
    return nil, resp.StatusCode, err
  }

  return ap, resp.StatusCode, nil
}

// Linked_type:
//
// > FAVORITES_LINKED_TYPE_ANIME, FAVORITES_LINKED_TYPE_MANGA,
// FAVORITES_LINKED_TYPE_RANOBE, FAVORITES_LINKED_TYPE_PERSON,
// FAVORITES_LINKED_TYPE_CHARACTER;
//
// Kind(required when Linked_type is Person):
//
// > FAVORITES_KIND_COMMON, FAVORITES_KIND_SEYU, FAVORITES_KIND_MANGAKA,
// FAVORITES_KIND_PRODUCER, FAVORITES_KIND_PERSON;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesCreate(linked_type string, kind string) (api.Favorites, error) {
  var fa api.Favorites
  var client = &http.Client{}

  if linked_type != FAVORITES_LINKED_TYPE_PERSON { kind = "" }

  str_id := strconv.Itoa(f.Id)

  post, cancel, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 10(favorites/) + ?(linked_type) + 1(/) + ?(id) + 1(/) + ?(kind)
    concat.Url(38+len(linked_type)+len(str_id)+len(kind), []string{
      SITE, "favorites/", linked_type, "/", str_id, "/", kind,
    }), 10,
  )
  if err != nil {
    return fa, err
  }
  defer cancel()

  resp, err := client.Do(post)
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

// Linked_type:
//
// > FAVORITES_LINKED_TYPE_ANIME, FAVORITES_LINKED_TYPE_MANGA,
// FAVORITES_LINKED_TYPE_RANOBE, FAVORITES_LINKED_TYPE_PERSON,
// FAVORITES_LINKED_TYPE_CHARACTER;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesDelete(linked_type string) (api.Favorites, error) {
  var ff api.Favorites
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  remove, cancel, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 10(favorites/) + ?(linked_type) + 1(/) + ?(id)
    concat.Url(37+len(linked_type)+len(str_id), []string{
      SITE, "favorites/", linked_type, "/", str_id,
    }), 10,
  )
  if err != nil {
    return ff, err
  }
  defer cancel()

  resp, err := client.Do(remove)
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

// FIXME: https://github.com/heycatch/goshikimori/issues/14
//
// Position: a new position on the list, it starts from 0.
//
// You can only get a StatusCode.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesReorder(position int) (int, error) {
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  post, cancel, err := NewReorderPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 10(favorites/) + ?(id) + 8(/reorder)
    concat.Url(44+len(str_id), []string{
      SITE, "favorites/", str_id, "/reorder",
    }), position, 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/ignore
func (f *FastId) AddIgnoreUser() (api.IgnoreUser, error) {
  var i api.IgnoreUser
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  post, cancel, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 9(v2/users/) + ?(id) + 7(/ignore)
    concat.Url(42+len(str_id), []string{
      SITE, "v2/users/", str_id, "/ignore",
    }), 10,
  )
  if err != nil {
    return i, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return i, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return i, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, err
  }

  return i, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/ignore
func (f *FastId) RemoveIgnoreUser() (api.IgnoreUser, error) {
  var i api.IgnoreUser
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  remove, cancel, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 9(v2/users/) + ?(id) + 7(/ignore)
    concat.Url(42+len(str_id), []string{
      SITE, "v2/users/", str_id, "/ignore",
    }), 10,
  )
  if err != nil {
    return i, err
  }
  defer cancel()

  resp, err := client.Do(remove)
  if err != nil {
    return i, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return i, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, err
  }

  return i, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (c *Configuration) Dialogs() ([]api.Dialogs, int, error) {
  var d []api.Dialogs
  var client = &http.Client{}

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 7(dialogs)
    concat.Url(33, []string{SITE, "dialogs"}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &d); err != nil {
    return nil, resp.StatusCode, err
  }

  return d, resp.StatusCode, nil
}

// When using FastIdUser()/SetFastId(), specify the user's nickname (not your own).
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (f *FastId) SearchDialogs() ([]api.SearchDialogs, error) {
  var sd []api.SearchDialogs
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(dialogs/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "dialogs/", str_id,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
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

// When using FastIdUser()/SetFastId(), specify the user's nickname (not your own).
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (f *FastId) DeleteDialogs() (api.FriendRequest, error) {
  var fr api.FriendRequest
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  remove, cancel, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(dialogs/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "dialogs/", str_id,
    }), 10,
  )
  if err != nil {
    return fr, err
  }
  defer cancel()

  resp, err := client.Do(remove)
  if err != nil {
    return fr, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return fr, err
  }

  if err := json.Unmarshal(data, &fr); err != nil {
    // Original error message from api/v1.
    return fr, errors.New("не найдено ни одного сообщения для удаления")
  }

  return fr, nil
}

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) UserBriefInfo() (api.Info, error) {
  var i api.Info
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 5(/info)
    concat.Url(37+len(str_id), []string{
      SITE, "users/", str_id, "/info",
    }), 10,
  )
  if err != nil {
    return i, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return i, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return i, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, err
  }

  return i, nil
}

// https://github.com/heycatch/goshikimori/issues/26
func (c *Configuration) SignOut() (string, int, error) {
  var client = &http.Client{}

  post, cancel, err := NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(users/sign_out)
    concat.Url(40, []string{
      SITE, "users/sign_out",
    }), 10,
  )
  if err != nil {
    return "", 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return "", resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return "", resp.StatusCode, err
  }

  return string(data), resp.StatusCode, nil
}

// Users having at least 1 completed animes and active during last month.
//
// Time to complete request increased to 40 seconds. Too big request.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/active_users
func (c *Configuration) ActiveUsers() ([]int, int, error) {
  var client = &http.Client{}
  ids := make([]int, 0)

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 18(stats/active_users)
    concat.Url(44, []string{
      SITE, "stats/active_users",
    }), 40,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ids); err != nil {
    return nil, resp.StatusCode, err
  }

  return ids, resp.StatusCode, nil
}

// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsAnime(r Result) ([]api.Topics, error) {
  var t []api.Topics
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(animes/) + ?(id) + 8(/topics?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "animes/", str_id, "/topics?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, err
  }

  return t, nil
}

// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsManga(r Result) ([]api.Topics, error) {
  var t []api.Topics
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(mangas/) + ?(id) + 8(/topics?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "mangas/", str_id, "/topics?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, err
  }

  return t, nil
}

// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsRanobe(r Result) ([]api.Topics, error) {
  var t []api.Topics
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 7(ranobe/) + ?(id) + 8(/topics?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "ranobe/", str_id, "/topics?", opt,
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, err
  }

  return t, nil
}

// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
//  - Forum:
//
//  > TOPIC_FORUM_ALL, TOPIC_FORUM_COSPLAY, TOPIC_FORUM_ANIMANGA, TOPIC_FORUM_SITE,
//  TOPIC_FORUM_GAMES, TOPIC_FORUM_VN, TOPIC_FORUM_CONTEST, TOPIC_FORUM_OFFTOPIC,
//  TOPIC_FORUM_CLUBS, TOPIC_FORUM_MYCLUBS, TOPIC_FORUM_CRITIQUES,
//  TOPIC_FORUM_NEWS, TOPIC_FORUM_COLLECTIONS, TOPIC_FORUM_ARTICLES;
//
//  - Linked_id: number without limit;
//  - Linked_type:
//
//  > TOPIC_LINKED_TYPE_ANIME, TOPIC_LINKED_TYPE_MANGA, TOPIC_LINKED_TYPE_RANOBE,
//  TOPIC_LINKED_TYPE_CHARACTER, TOPIC_LINKED_TYPE_PERSON, TOPIC_LINKED_TYPE_CLUB,
//  TOPIC_LINKED_TYPE_CLUBPAGE, TOPIC_LINKED_TYPE_CRITIQUE, TOPIC_LINKED_TYPE_REVIEW,
//  TOPIC_LINKED_TYPE_CONTEST, TOPIC_LINKED_TYPE_COSPLAYGALLYRY,
//  TOPIC_LINKED_TYPE_COLLECTION, TOPIC_LINKED_TYPE_ARTICLE;
//
// REMARK: linked_id and linked_type are only used together.
//
//  - Type: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopics(r Result) ([]api.Topics, int, error) {
  var t []api.Topics
  var client = &http.Client{}

  opt := r.OptionsTopics()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 7(topics?) + ?(Result)
    concat.Url(33+len(opt), []string{SITE, "topics?", opt}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, resp.StatusCode, err
  }

  return t, resp.StatusCode, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsUpdates(r Result) ([]api.TopicsUpdates, int, error) {
  var t []api.TopicsUpdates
  var client = &http.Client{}

  opt := r.OptionsOnlyPageLimit(100000, 30)

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 15(topics/updates?) + ?(Result)
    concat.Url(41+len(opt), []string{SITE, "topics/updates?", opt}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, resp.StatusCode, err
  }

  return t, resp.StatusCode, nil
}

// 'Options' settings:
//  - Limit: 10 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsHot(r Result) ([]api.Topics, int, error) {
  var t []api.Topics
  var client = &http.Client{}

  opt := r.OptionsTopicsHot()

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 11(topics/hot?) + ?(Result)
    concat.Url(37+len(opt), []string{SITE, "topics/hot?", opt}), 10,
  )
  if err != nil {
    return nil, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, resp.StatusCode, err
  }

  return t, resp.StatusCode, nil
}

// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsId(id int) (api.TopicsId, int, error) {
  var t api.TopicsId
  var client = &http.Client{}

  str_id := strconv.Itoa(id)

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 7(topics/) + ?(id)
    concat.Url(33+len(str_id), []string{SITE, "topics/", str_id}), 10,
  )
  if err != nil {
    return t, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return t, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return t, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return t, resp.StatusCode, err
  }

  return t, resp.StatusCode, nil
}

// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) AddIgnoreTopic(id int) (api.IgnoreTopic, int, error) {
  var i api.IgnoreTopic
  var client = &http.Client{}

  str_id := strconv.Itoa(id)

  post, cancel, err := NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 10(v2/topics/) + ?(id) + 7(/ignore)
    concat.Url(43+len(str_id), []string{
      SITE, "v2/topics/", str_id, "/ignore",
    }), 10,
  )
  if err != nil {
    return i, 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return i, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return i, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, resp.StatusCode, err
  }

  return i, resp.StatusCode, nil
}

// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) RemoveIgnoreTopic(id int) (api.IgnoreTopic, int, error) {
  var i api.IgnoreTopic
  var client = &http.Client{}

  str_id := strconv.Itoa(id)

  remove, cancel, err := NewDeleteRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 10(v2/topics/) + ?(id) + 7(/ignore)
    concat.Url(43+len(str_id), []string{
      SITE, "v2/topics/", str_id, "/ignore",
    }), 10,
  )
  if err != nil {
    return i, 0, err
  }
  defer cancel()

  resp, err := client.Do(remove)
  if err != nil {
    return i, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return i, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, resp.StatusCode, err
  }

  return i, resp.StatusCode, nil
}

// Schema: customized request.
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func (c *Configuration) SearchGraphql(schema string) (api.GraphQL, int, error) {
  var client = &http.Client{}
  var g api.GraphQL

  post, cancel, err := NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + ?(schema)
    concat.Url(26+len(schema), []string{SITE, schema}), 10,
  )
  if err != nil {
    return g, 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return g, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return g, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &g); err != nil {
    return g, resp.StatusCode, err
  }

  return g, resp.StatusCode, nil
}

// Id: message id.
//
// Ignore:
//   - Linked_type: nil;
//   - Linked: nil;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) ReadMessage(id int) (api.Messages, int, error) {
  var client = &http.Client{}
  var m api.Messages

  str_id := strconv.Itoa(id)

  get, cancel, err := NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(messages/) + ?(id)
    concat.Url(35+len(str_id), []string{SITE, "messages/", str_id}), 10,
  )
  if err != nil {
    return m, 0, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return m, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return m, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// From_id: your Id.
//
// To_id: the Id of the person you want to send the message to.
//
// Message: message text.
//
// Returns a status of 201.
//
// Ignore:
//   - Linked_type: nil;
//   - Linked: nil;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) SendMessage(from_id, to_id int, message string) (api.Messages, int, error) {
  var client = &http.Client{}
  var m api.Messages

  post, cancel, err := NewSendMessagePostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 8(messages)
    concat.Url(34, []string{SITE, "messages"}), message, from_id, to_id, 10,
  )
  if err != nil {
    return m, 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return m, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return m, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// Id: message id.
//
// Message: message text.
//
// Ignore:
//   - Linked_type: nil;
//   - Linked: nil;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) ChangeMessage(id int, message string) (api.Messages, int, error) {
  var client = &http.Client{}
  var m api.Messages

  str_id := strconv.Itoa(id)

  put, cancel, err := NewChangeMessagePutRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(messages/) + ?(id)
    concat.Url(35+len(str_id), []string{SITE, "messages/", str_id}), message, 10,
  )
  if err != nil {
    return m, 0, err
  }
  defer cancel()

  resp, err := client.Do(put)
  if err != nil {
    return m, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return m, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, resp.StatusCode, err
  }

  return m, resp.StatusCode, nil
}

// Id: message id.
//
// Only status 204 is returned.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) DeleteMessage(id int) (int, error) {
  var client = &http.Client{}

  str_id := strconv.Itoa(id)

  del, cancel, err := NewDeleteMessageDeleteRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(messages/) + ?(id)
    concat.Url(35+len(str_id), []string{SITE, "messages/", str_id}), 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(del)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// Ids: array of ids converted to a string.
//
// Is_read: mark a message as read or unread.
//
// 'Is_read' settings:
//  - 1 (read)
//  - 0 (unread)
//
// Only status 200 is returned.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) MarkReadMessages(ids string, is_read int) (int, error) {
  var client = &http.Client{}

  post, cancel, err := NewMarkReadPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 18(messages/mark_read)
    concat.Url(44, []string{
      SITE, "messages/mark_read",
    }), ids, is_read, 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// Name: unread message type.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_MESSAGES,
// UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// Empty array to be filled with ids for messages.
func (f *FastId) UnreadMessagesIds(name string) ([]int, error) {
  var um api.UnreadMessages
  var client = &http.Client{}

  str_id := strconv.Itoa(f.Id)

  get, cancel, err := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 16(/unread_messages)
    concat.Url(48+len(str_id), []string{
      SITE, "users/", str_id, "/unread_messages",
    }), 10,
  )
  if err != nil {
    return nil, err
  }
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &um); err != nil {
    return nil, err
  }

  switch name {
  case "messages":
    if um.Messages == 0 { return nil, errors.New("unread messages not found") }
    return make([]int, um.Messages), nil
  case "news":
    if um.News == 0 { return nil, errors.New("unread news not found") }
    return make([]int, um.News), nil
  case "notifications":
    if um.Notifications == 0 { return nil, errors.New("unread notifications not found") }
    return make([]int, um.Notifications), nil
  default:
    return nil, errors.New("wrong name... try messages, news or notifications")
  }
}

// Name: type in the mail.
//
// Returns a status of 200.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) ReadAllMessages(name string) (int, error) {
  var client = &http.Client{}

  post, cancel, err := NewReadDeleteAllPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 17(messages/read_all)
    concat.Url(43, []string{SITE, "messages/read_all"}), name, 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}

// Name: type in the mail.
//
// Returns a status of 200.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) DeleteAllMessages(name string) (int, error) {
  var client = &http.Client{}

  post, cancel, err := NewReadDeleteAllPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 19(messages/delete_all)
    concat.Url(45, []string{SITE, "messages/delete_all"}), name, 10,
  )
  if err != nil {
    return 0, err
  }
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}
