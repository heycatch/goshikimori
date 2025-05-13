// Copyright (C) 2025 heycatch <andreyisback@yandex.ru>.
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
  "encoding/json"
  "errors"
  "net/url"
  "strconv"

  "github.com/heycatch/goshikimori/api"
  "github.com/heycatch/goshikimori/concat"
)

// Only the application needs to be specified in SetConfiguration().
//
// Name: user name.
//
// Search by user is case sensitive.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (c *Configuration) SearchUser(name string) (api.Users, int, error) {
  var u api.Users

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 6(users/) + ?(name)
    concat.Url(32+len(name), []string{SITE, "users/", url.QueryEscape(name)}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return u, status, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return u, status, err
  }

  return u, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
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

  opt := r.OptionsOnlyPageLimit(100000, 100)

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 13(users?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(40+len(name)+len(opt), []string{
      SITE, "users?search=", url.QueryEscape(name), "&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return u, status, err
  }

  return u, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserFriends(r Result) ([]api.UserFriends, int, error) {
  var uf []api.UserFriends

  // 26(SITE) + 6(users/) + ?(id) + 9(/friends?) + ?(Result)
  opt := r.OptionsOnlyPageLimit(100000, 100)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 9(/friends?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/friends?" + opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return nil, status, err
  }

  return uf, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserClubs() ([]api.Clubs, int, error) {
  var uc []api.Clubs

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 6(/clubs)
    concat.Url(38+len(str_id), []string{
      SITE, "users/", str_id, "/clubs",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &uc); err != nil {
    return nil, status, err
  }

  return uc, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
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
func (f *FastId) SearchUserAnimeRates(r Result) ([]api.UserAnimeRates, int, error) {
  var ar []api.UserAnimeRates

  opt := r.OptionsAnimeRates()
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 13(/anime_rates?) + ?(Result)
    concat.Url(45+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/anime_rates?" + opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ar); err != nil {
    return nil, status, err
  }

  return ar, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
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
func (f *FastId) SearchUserMangaRates(r Result) ([]api.UserMangaRates, int, error) {
  var mr []api.UserMangaRates

  opt := r.OptionsMangaRates()
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 13(/manga_rates?) + ?(Result)
    concat.Url(45+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/manga_rates?" + opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &mr); err != nil {
    return nil, status, err
  }

  return mr, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserFavourites() (api.UserFavourites, int, error) {
  var uf api.UserFavourites

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 11(/favourites)
    concat.Url(43+len(str_id), []string{
      SITE, "users/", str_id, "/favourites",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return uf, status, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return uf, status, err
  }

  return uf, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
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
func (f *FastId) SearchUserHistory(r Result) ([]api.UserHistory, int, error) {
  var uh []api.UserHistory

  opt := r.OptionsUserHistory()
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 9(/history?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/history?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &uh); err != nil {
    return nil, status, err
  }

  return uh, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserBans() ([]api.Bans, int, error) {
  var b []api.Bans

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 5(/bans)
    concat.Url(37+len(str_id), []string{
      SITE, "users/", str_id, "/bans",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &b); err != nil {
    return nil, status, err
  }

  return b, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/whoami
func (c *Configuration) WhoAmi() (api.Who, int, error) {
  var w api.Who

  data, status, err := NewGetRequestWithCancelAndBearer(
    c.Application, c.AccessToken,
    // 26(SITE) + 12(users/whoami)
    concat.Url(38, []string{SITE, "users/whoami"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return w, status, err
  }

  if err := json.Unmarshal(data, &w); err != nil {
    return w, status, err
  }

  return w, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchAnime() (api.Anime, int, error) {
  var a api.Anime

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "animes/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return a, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return a, status, err
  }

  return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsAnime()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(animes?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(41+len(name)+len(opt), []string{
      SITE, "animes?search=", url.QueryEscape(name), "&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, status, err
  }

  return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchManga() (api.Manga, int, error) {
  var m api.Manga

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "mangas/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return m, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsManga()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(mangas?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(41+len(name)+len(opt), []string{
      SITE, "mangas?search=", url.QueryEscape(name), "&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchRanobe() (api.Manga, int, error) {
  var m api.Manga

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "ranobe/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return m, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsRanobe()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(ranobe?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(41+len(name)+len(opt), []string{
      SITE, "ranobe?search=", url.QueryEscape(name), "&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: user name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdUser(name string) (*FastId, int, error) {
  var u api.Users

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 6(users/) + ?(name)
    concat.Url(32+len(name), []string{
      SITE, "users/", url.QueryEscape(name)}), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &u); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  return &FastId{Id: u.Id, Conf: *c, Err: err}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: anime name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdAnime(name string) (*FastId, int, error) {
  var a []api.Animes

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(animes?search=) + ?(name)
    concat.Url(40+len(name), []string{
      SITE, "animes?search=", url.QueryEscape(name)}), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(a) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil }

  return &FastId{Id: a[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: manga name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdManga(name string) (*FastId, int, error) {
  var m []api.Mangas

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(mangas?search=) + ?(name)
    concat.Url(40+len(name), []string{
      SITE, "mangas?search=", url.QueryEscape(name)}), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(m) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil }

  return &FastId{Id: m[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: ranobe name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdRanobe(name string) (*FastId, int, error) {
  var m []api.Mangas

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(ranobe?search=) + ?(name)
    concat.Url(40+len(name), []string{
      SITE, "ranobe?search=", url.QueryEscape(name)}), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(m) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil }

  return &FastId{Id: m[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: club name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdClub(name string) (*FastId, int, error) {
  var cl []api.Clubs

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 13(clubs?search=) + ?(name)
    concat.Url(39+len(name), []string{
      SITE, "clubs?search=", url.QueryEscape(name)}), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(cl) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil }

  return &FastId{Id: cl[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: character name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdCharacter(name string) (*FastId, int, error) {
  var ch []api.CharacterInfo

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 25(characters/search?search=) + ?(name)
    concat.Url(51+len(name), []string{
      SITE, "characters/search?search=", url.QueryEscape(name),
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ch); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(ch) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil }

  return &FastId{Id: ch[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: people name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdPeople(name string) (*FastId, int, error) {
  var ap []api.AllPeople
  // testing

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 21(people/search?search=) + ?(name)
    concat.Url(47+len(name), []string{
      SITE, "people/search?search=", url.QueryEscape(name),
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ap); err != nil {
    return &FastId{Id: 0, Conf: *c, Err: err}, status, err
  }

  // if len == 0; we get panic: runtime error.
  // To avoid a crash, process the error here.
  //
  // There is no point in processing the error. there is no place to catch it.
  if len(ap) == 0 { return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil }

  return &FastId{Id: ap[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_screenshots
func (f *FastId) SearchAnimeScreenshots() ([]api.AnimeScreenshots, int, error) {
  var s []api.AnimeScreenshots

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 12(/screenshots)
    concat.Url(45+len(str_id), []string{
      SITE, "animes/", str_id, "/screenshots",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &s); err != nil {
    return nil, status, err
  }

  return s, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchAnimeFranchise() (api.Franchise, int, error) {
  var ff api.Franchise

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 10(/franchise)
    concat.Url(43+len(str_id), []string{
      SITE, "animes/", str_id, "/franchise",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ff, status, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, status, err
  }

  return ff, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchMangaFranchise() (api.Franchise, int, error) {
  var ff api.Franchise

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id) + 10(/franchise)
    concat.Url(43+len(str_id), []string{
      SITE, "mangas/", str_id, "/franchise",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ff, status, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, status, err
  }

  return ff, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchRanobeFranchise() (api.Franchise, int, error) {
  var ff api.Franchise

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id) + 10(/franchise)
    concat.Url(43+len(str_id), []string{
      SITE, "ranobe/", str_id, "/franchise",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ff, status, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, status, err
  }

  return ff, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchAnimeExternalLinks() ([]api.ExternalLinks, int, error) {
  var el []api.ExternalLinks

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 15(/external_links)
    concat.Url(48+len(str_id), []string{
      SITE, "animes/", str_id, "/external_links",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, status, err
  }

  return el, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchMangaExternalLinks() ([]api.ExternalLinks, int, error) {
  var el []api.ExternalLinks

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id) + 15(/external_links)
    concat.Url(48+len(str_id), []string{
      SITE, "mangas/", str_id, "/external_links",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, status, err
  }

  return el, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchRanobeExternalLinks() ([]api.ExternalLinks, int, error) {
  var el []api.ExternalLinks

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id) + 15(/external_links)
    concat.Url(48+len(str_id), []string{
      SITE, "ranobe/", str_id, "/external_links",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &el); err != nil {
    return nil, status, err
  }

  return el, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarAnime() ([]api.Animes, int, error) {
  var a []api.Animes

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 8(/similar)
    concat.Url(41+len(str_id), []string{
      SITE, "animes/", str_id, "/similar",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, status, err
  }

  return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarManga() ([]api.Mangas, int, error) {
  var m []api.Mangas

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id) + 8(/similar)
    concat.Url(41+len(str_id), []string{
      SITE, "mangas/", str_id, "/similar",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarRanobe() ([]api.Mangas, int, error) {
  var m []api.Mangas

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id) + 8(/similar)
    concat.Url(41+len(str_id), []string{
      SITE, "ranobe/", str_id, "/similar",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedAnime() ([]api.RelatedAnimes, int, error) {
  var a []api.RelatedAnimes

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 8(/related)
    concat.Url(41+len(str_id), []string{
      SITE, "animes/", str_id, "/related",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, status, err
  }

  return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedManga() ([]api.RelatedMangas, int, error) {
  var m []api.RelatedMangas

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id) + 8(/related)
    concat.Url(41+len(str_id), []string{
      SITE, "mangas/", str_id, "/related",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedRanobe() ([]api.RelatedMangas, int, error) {
  var m []api.RelatedMangas

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id) + 8(/related)
    concat.Url(41+len(str_id), []string{
      SITE, "ranobe/", str_id, "/related",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
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

  opt := r.OptionsOnlyPageLimit(100000, 30)

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 13(clubs?search=) + ?(name) + 1(&) + ?(Result)
    concat.Url(40+len(name)+len(opt), []string{
      SITE, "clubs?search=", url.QueryEscape(name), "&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &cl); err != nil {
    return nil, status, err
  }

  return cl, status, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubAnimes(r Result) ([]api.Animes, int, error) {
  var a []api.Animes

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/animes?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/animes?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, status, err
  }

  return a, status, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMangas(r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/mangas?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/mangas?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubRanobe(r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/ranobe?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/ranobe?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCharacters(r Result) ([]api.CharacterInfo, int, error) {
  var ci []api.CharacterInfo

  opt := r.OptionsOnlyPageLimit(100000, 20)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 12(/characters?) + ?(Result)
    concat.Url(44+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/characters?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ci); err != nil {
    return nil, status, err
  }

  return ci, status, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubClubs(r Result) ([]api.Clubs, int, error) {
  var cc []api.Clubs

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 7(/clubs?) + ?(Result)
    concat.Url(39+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/clubs?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return nil, status, err
  }

  return cc, status, nil
}

// FIXME: The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Page: 4 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCollections(r Result) ([]api.ClubCollections, int, error) {
  var cc []api.ClubCollections

  opt := r.OptionsOnlyPageLimit(100000, 4)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 13(/collections?) + ?(Result)
    concat.Url(45+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/collections?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return nil, status, err
  }

  return cc, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMembers(r Result) ([]api.UserFriends, int, error) {
  var uf []api.UserFriends

  opt := r.OptionsOnlyPageLimit(100000, 100)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 9(/members?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/members?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &uf); err != nil {
    return nil, status, err
  }

  return uf, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubImages(r Result) ([]api.ClubImages, int, error) {
  var cm []api.ClubImages

  opt := r.OptionsOnlyPageLimit(100000, 100)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(clubs/) + ?(id) + 8(/images?) + ?(Result)
    concat.Url(40+len(str_id)+len(opt), []string{
      SITE, "clubs/", str_id, "/images?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &cm); err != nil {
    return nil, status, err
  }

  return cm, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) ClubJoin() (int, error) {
  str_id := strconv.Itoa(f.Id)

  _, status, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 5(/join)
    concat.Url(37+len(str_id), []string{
      SITE, "clubs/", str_id, "/join",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return status, err
  }

  return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) ClubLeave() (int, error) {
  str_id := strconv.Itoa(f.Id)

  _, status, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(clubs/) + ?(id) + 6(/leave)
    concat.Url(38+len(str_id), []string{
      SITE, "clubs/", str_id, "/leave",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return status, err
  }

  return status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// As a result, we return a complete list of all achievements.
//
// Next comes the filtering through "NekoSearch" and the error about obtaining
// specific achievements is already being processed there.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/achievements
func (f *FastId) SearchAchievement() ([]api.Achievements, int, error) {
  var a []api.Achievements

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 21(achievements?user_id=) + ?(id)
    concat.Url(47+len(str_id), []string{
      SITE, "achievements?user_id=", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, status, err
  }

  return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/video
func (f *FastId) SearchAnimeVideos() ([]api.AnimeVideos, int, error) {
  var v []api.AnimeVideos

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    concat.Url(40+len(str_id), []string{
      SITE, "animes/", str_id, "/videos",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &v); err != nil {
    return nil, status, err
  }

  return v, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/roles
func (f *FastId) SearchAnimeRoles() ([]api.Roles, int, error) {
  var r []api.Roles

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 6(/roles)
    concat.Url(39+len(str_id), []string{
      SITE, "animes/", str_id, "/roles",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, status, err
  }

  return r, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/roles
func (f *FastId) SearchMangaRoles() ([]api.Roles, int, error) {
  var r []api.Roles

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id) + 6(/roles)
    concat.Url(39+len(str_id), []string{
      SITE, "mangas/", str_id, "/roles",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, status, err
  }

  return r, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/roles
func (f *FastId) SearchRanobeRoles() ([]api.Roles, int, error) {
  var r []api.Roles

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id) + 6(/roles)
    concat.Url(39+len(str_id), []string{
      SITE, "ranobe/", str_id, "/roles",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &r); err != nil {
    return nil, status, err
  }

  return r, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/bans
func (c *Configuration) SearchBans() ([]api.Bans, int, error) {
  var b []api.Bans

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 4(bans)
    concat.Url(30, []string{SITE, "bans"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &b); err != nil {
    return nil, status, err
  }

  return b, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsCalendar()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 9(calendar?) + ?(Result)
    concat.Url(35+len(opt), []string{SITE, "calendar?", opt}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ca); err != nil {
    return nil, status, err
  }

  return ca, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name:
//
// > GENRES_ANIME, GENRES_MANGA;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/genres
func (c *Configuration) SearchGenres(name string) ([]api.Genres, int, error) {
  var g []api.Genres

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 12(genres?kind=) + ?(name)
    concat.Url(38+len(name), []string{SITE, "genres?kind=", name}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &g); err != nil {
    return nil, status, err
  }

  return g, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/studios
func (c *Configuration) SearchStudios() ([]api.Studios, int, error) {
  var s []api.Studios

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 7(studios)
    concat.Url(33, []string{SITE, "studios"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &s); err != nil {
    return nil, status, err
  }

  return s, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/publishers
func (c *Configuration) SearchPublishers() ([]api.Publishers, int, error) {
  var p []api.Publishers

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 10(publishers)
    concat.Url(36, []string{SITE, "publishers"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &p); err != nil {
    return nil, status, err
  }

  return p, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/forums
func (c *Configuration) SearchForums() ([]api.Forums, int, error) {
  var f []api.Forums

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 6(forums)
    concat.Url(32, []string{SITE, "forums"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &f); err != nil {
    return nil, status, err
  }

  return f, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/add_remove_friend
func (f *FastId) AddFriend() (api.FriendRequest, int, error) {
  var ff api.FriendRequest

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(friends/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "friends/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ff, status, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, status, err
  }

  return ff, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/add_remove_friend
func (f *FastId) RemoveFriend() (api.FriendRequest, int, error) {
  var ff api.FriendRequest

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(friends/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "friends/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ff, status, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, status, err
  }

  return ff, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Show current user unread messages counts.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/messages
func (f *FastId) UserUnreadMessages() (api.UnreadMessages, int, error) {
  var um api.UnreadMessages

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancelAndBearer(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 16(/unread_messages)
    concat.Url(48+len(str_id), []string{
      SITE, "users/", str_id, "/unread_messages",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return um, status, err
  }

  if err := json.Unmarshal(data, &um); err != nil {
    return um, status, err
  }

  return um, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
func (f *FastId) UserMessages(r Result) ([]api.Messages, int, error) {
  var m []api.Messages

  opt := r.OptionsMessages()
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancelAndBearer(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 10(/messages?) + ?(Result)
    concat.Url(42+len(str_id)+len(opt), []string{
      SITE, "users/", str_id, "/messages?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsAnime() (api.Constants, int, error) {
  var ca api.Constants

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 15(constants/anime)
    concat.Url(41, []string{SITE, "constants/anime"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return ca, status, err
  }

  if err := json.Unmarshal(data, &ca); err != nil {
    return ca, status, err
  }

  return ca, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsManga() (api.Constants, int, error) {
  var cm api.Constants

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 15(constants/manga)
    concat.Url(41, []string{SITE, "constants/manga"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return cm, status, err
  }

  if err := json.Unmarshal(data, &cm); err != nil {
    return cm, status, err
  }

  return cm, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsUserRate() (api.ConstantsUserRate, int, error) {
  var ur api.ConstantsUserRate

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 19(constants/user_rate)
    concat.Url(45, []string{SITE, "constants/user_rate"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return ur, status, err
  }

  if err := json.Unmarshal(data, &ur); err != nil {
    return ur, status, err
  }

  return ur, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsClub() (api.ConstantsClub, int, error) {
  var cc api.ConstantsClub

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 14(constants/club)
    concat.Url(40, []string{SITE, "constants/club"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return cc, status, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return cc, status, err
  }

  return cc, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsSmileys() ([]api.ConstantsSmileys, int, error) {
  var cs []api.ConstantsSmileys

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 17(constants/smileys)
    concat.Url(43, []string{SITE, "constants/smileys"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &cs); err != nil {
    return nil, status, err
  }

  return cs, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsRandomAnime()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 20(animes?order=random&) + ?(Result)
    concat.Url(46+len(opt), []string{
      SITE, "animes?order=random&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &a); err != nil {
    return nil, status, err
  }

  return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsRandomManga()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 20(mangas?order=random&) + ?(Result)
    concat.Url(46+len(opt), []string{
      SITE, "mangas?order=random&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsRandomRanobe()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 20(ranobe?order=random&) + ?(Result)
    concat.Url(46+len(opt), []string{
      SITE, "ranobe?order=random&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return nil, status, err
  }

  return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/character
func (f *FastId) SearchCharacter() (api.Character, int, error) {
  var ch api.Character

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 11(characters/) + ?(id)
    concat.Url(37+len(str_id), []string{
      SITE, "characters/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ch, status, err
  }

  if err := json.Unmarshal(data, &ch); err != nil {
    return ch, status, err
  }

  return ch, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: character name.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/character
func (c *Configuration) SearchCharacters(name string) ([]api.CharacterInfo, int, error) {
  var ci []api.CharacterInfo

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 25(characters/search?search=) + ?(name)
    concat.Url(51+len(name), []string{SITE,
      "characters/search?search=", url.QueryEscape(name)}),
      MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ci); err != nil {
    return nil, status, err
  }

  return ci, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/people
func (f *FastId) SearchPeople() (api.People, int, error) {
  var p api.People

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(people/) + ?(id)
    concat.Url(33+len(str_id), []string{
      SITE, "people/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return p, status, err
  }

  if err := json.Unmarshal(data, &p); err != nil {
    return p, status, err
  }

  return p, status, nil
}

// FIXME: Page and limit not supprted, idk why. Check later.
//
// Only the application needs to be specified in SetConfiguration().
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

  opt := r.OptionsPeople()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    concat.Url(48+len(name)+len(opt), []string{
      SITE, "people/search?search=", url.QueryEscape(name), "&", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ap); err != nil {
    return nil, status, err
  }

  return ap, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
func (f *FastId) FavoritesCreate(linked_type, kind string) (api.Favorites, int, error) {
  var fa api.Favorites

  if linked_type != FAVORITES_LINKED_TYPE_PERSON { kind = "" }

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 10(favorites/) + ?(linked_type) + 1(/) + ?(id) + 1(/) + ?(kind)
    concat.Url(38+len(linked_type)+len(str_id)+len(kind), []string{
      SITE, "favorites/", linked_type, "/", str_id, "/", kind,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return fa, status, err
  }

  if err := json.Unmarshal(data, &fa); err != nil {
    return fa, status, err
  }

  return fa, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Linked_type:
//
// > FAVORITES_LINKED_TYPE_ANIME, FAVORITES_LINKED_TYPE_MANGA,
// FAVORITES_LINKED_TYPE_RANOBE, FAVORITES_LINKED_TYPE_PERSON,
// FAVORITES_LINKED_TYPE_CHARACTER;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesDelete(linked_type string) (api.Favorites, int, error) {
  var ff api.Favorites

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 10(favorites/) + ?(linked_type) + 1(/) + ?(id)
    concat.Url(37+len(linked_type)+len(str_id), []string{
      SITE, "favorites/", linked_type, "/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return ff, status, err
  }

  if err := json.Unmarshal(data, &ff); err != nil {
    return ff, status, err
  }

  return ff, status, nil
}

// FIXME: https://github.com/heycatch/goshikimori/issues/14
//
// In SetConfiguration(), you must specify the application and the token.
//
// Position: a new position on the list, it starts from 0.
//
// You can only get a StatusCode.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesReorder(position int) (int, error) {
  str_id := strconv.Itoa(f.Id)

  _, status, err := NewReorderPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 10(favorites/) + ?(id) + 8(/reorder)
    concat.Url(44+len(str_id), []string{
      SITE, "favorites/", str_id, "/reorder",
    }), position, MAX_EXPECTATION,
  )
  if err != nil {
    return status, err
  }

  return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/ignore
func (f *FastId) AddIgnoreUser() (api.IgnoreUser, int, error) {
  var i api.IgnoreUser

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 9(v2/users/) + ?(id) + 7(/ignore)
    concat.Url(42+len(str_id), []string{
      SITE, "v2/users/", str_id, "/ignore",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return i, status, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, status, err
  }

  return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/ignore
func (f *FastId) RemoveIgnoreUser() (api.IgnoreUser, int, error) {
  var i api.IgnoreUser

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 9(v2/users/) + ?(id) + 7(/ignore)
    concat.Url(42+len(str_id), []string{
      SITE, "v2/users/", str_id, "/ignore",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return i, status, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, status, err
  }

  return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (c *Configuration) Dialogs() ([]api.Dialogs, int, error) {
  var d []api.Dialogs

  data, status, err := NewGetRequestWithCancelAndBearer(
    c.Application, c.AccessToken,
    // 26(SITE) + 7(dialogs)
    concat.Url(33, []string{SITE, "dialogs"}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &d); err != nil {
    return nil, status, err
  }

  return d, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// When using FastIdUser()/SetFastId(), specify the user's nickname (not your own).
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (f *FastId) SearchDialogs() ([]api.SearchDialogs, int, error) {
  var sd []api.SearchDialogs

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancelAndBearer(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(dialogs/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "dialogs/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &sd); err != nil {
    return nil, status, err
  }

  return sd, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// When using FastIdUser()/SetFastId(), specify the user's nickname (not your own).
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (f *FastId) DeleteDialogs() (api.FriendRequest, int, error) {
  var fr api.FriendRequest

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 8(dialogs/) + ?(id)
    concat.Url(34+len(str_id), []string{
      SITE, "dialogs/", str_id,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return fr, status, err
  }

  if err := json.Unmarshal(data, &fr); err != nil {
    // Original error message from api/v1.
    return fr, status, errors.New("не найдено ни одного сообщения для удаления")
  }

  return fr, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) UserBriefInfo() (api.Info, int, error) {
  var i api.Info

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 6(users/) + ?(id) + 5(/info)
    concat.Url(37+len(str_id), []string{
      SITE, "users/", str_id, "/info",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return i, status, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, status, err
  }

  return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// https://github.com/heycatch/goshikimori/issues/26
func (c *Configuration) SignOut() (string, int, error) {
  data, status, err := NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 14(users/sign_out)
    concat.Url(40, []string{
      SITE, "users/sign_out",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return string(data), status, err
  }

  return string(data), status, nil
}

// If we get a json unmarshal type error, it is a server-side error, namely:
// ["PG::DiskFull: ERROR:  could not resize shared memory segment \"/PostgreSQL.1559179908\" to 16777216 bytes: No space left on device\n"]
//
// Only the application needs to be specified in SetConfiguration().
//
// Users having at least 1 completed animes and active during last month.
//
// Time to complete request increased to 40 seconds. Too big request.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/active_users
func (c *Configuration) ActiveUsers() ([]int, int, error) {
  ids := make([]int, 0)

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 18(stats/active_users)
    concat.Url(44, []string{
      SITE, "stats/active_users",
    }), CUSTOM_MAX_EXPECTATION_ACTIVE_USERS,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &ids); err != nil {
    return nil, status, err
  }

  return ids, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsAnime(r Result) ([]api.Topics, int, error) {
  var t []api.Topics

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(animes/) + ?(id) + 8(/topics?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "animes/", str_id, "/topics?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, status, err
  }

  return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsManga(r Result) ([]api.Topics, int, error) {
  var t []api.Topics

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(mangas/) + ?(id) + 8(/topics?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "mangas/", str_id, "/topics?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, status, err
  }

  return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsRanobe(r Result) ([]api.Topics, int, error) {
  var t []api.Topics

  opt := r.OptionsOnlyPageLimit(100000, 30)
  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancel(
    f.Conf.Application,
    // 26(SITE) + 7(ranobe/) + ?(id) + 8(/topics?) + ?(Result)
    concat.Url(41+len(str_id)+len(opt), []string{
      SITE, "ranobe/", str_id, "/topics?", opt,
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, status, err
  }

  return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
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

  opt := r.OptionsTopics()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 7(topics?) + ?(Result)
    concat.Url(33+len(opt), []string{SITE, "topics?", opt}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, status, err
  }

  return t, status, nil
}

// FIXME: Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
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

  opt := r.OptionsOnlyPageLimit(100000, 30)

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 15(topics/updates?) + ?(Result)
    concat.Url(41+len(opt), []string{SITE, "topics/updates?", opt}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, status, err
  }

  return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Limit: 10 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsHot(r Result) ([]api.Topics, int, error) {
  var t []api.Topics

  opt := r.OptionsTopicsHot()

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 11(topics/hot?) + ?(Result)
    concat.Url(37+len(opt), []string{SITE, "topics/hot?", opt}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return nil, status, err
  }

  return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsId(id int) (api.TopicsId, int, error) {
  var t api.TopicsId

  str_id := strconv.Itoa(id)

  data, status, err := NewGetRequestWithCancel(
    c.Application,
    // 26(SITE) + 7(topics/) + ?(id)
    concat.Url(33+len(str_id), []string{SITE, "topics/", str_id}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return t, status, err
  }

  if err := json.Unmarshal(data, &t); err != nil {
    return t, status, err
  }

  return t, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) AddIgnoreTopic(id int) (api.IgnoreTopic, int, error) {
  var i api.IgnoreTopic

  str_id := strconv.Itoa(id)

  data, status, err := NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 10(v2/topics/) + ?(id) + 7(/ignore)
    concat.Url(43+len(str_id), []string{
      SITE, "v2/topics/", str_id, "/ignore",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return i, 0, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, status, err
  }

  return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) RemoveIgnoreTopic(id int) (api.IgnoreTopic, int, error) {
  var i api.IgnoreTopic

  str_id := strconv.Itoa(id)

  data, status, err := NewDeleteRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 10(v2/topics/) + ?(id) + 7(/ignore)
    concat.Url(43+len(str_id), []string{
      SITE, "v2/topics/", str_id, "/ignore",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return i, 0, err
  }

  if err := json.Unmarshal(data, &i); err != nil {
    return i, status, err
  }

  return i, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Schema: customized request.
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func (c *Configuration) SearchGraphql(schema string) (api.GraphQL, int, error) {
  var g api.GraphQL

  data, status, err := NewGraphQLPostRequestWithCancel(
    c.Application,
    // 26(SITE) + ?(schema)
    concat.Url(26+len(schema), []string{SITE, schema}),
    CUSTOM_MAX_EXPECTATION_GRAPHQL,
  )
  if err != nil {
    return g, status, err
  }

  if err := json.Unmarshal(data, &g); err != nil {
    return g, status, err
  }

  return g, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
  var m api.Messages

  str_id := strconv.Itoa(id)

  data, status, err := NewGetRequestWithCancelAndBearer(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(messages/) + ?(id)
    concat.Url(35+len(str_id), []string{SITE, "messages/", str_id}), MAX_EXPECTATION,
  )
  if err != nil {
    return m, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, status, err
  }

  return m, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
  var m api.Messages

  data, status, err := NewSendMessagePostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 8(messages)
    concat.Url(34, []string{SITE, "messages"}),
    message, from_id, to_id, MAX_EXPECTATION,
  )
  if err != nil {
    return m, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, status, err
  }

  return m, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
  var m api.Messages

  str_id := strconv.Itoa(id)

  data, status, err := NewChangeMessagePutRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(messages/) + ?(id)
    concat.Url(35+len(str_id), []string{SITE, "messages/", str_id}),
    message, MAX_EXPECTATION,
  )
  if err != nil {
    return m, status, err
  }

  if err := json.Unmarshal(data, &m); err != nil {
    return m, status, err
  }

  return m, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Id: message id.
//
// Only status 204 is returned.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) DeleteMessage(id int) (int, error) {
  str_id := strconv.Itoa(id)

  _, status, err := NewDeleteMessageDeleteRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 9(messages/) + ?(id)
    concat.Url(35+len(str_id), []string{SITE, "messages/", str_id}),
    MAX_EXPECTATION,
  )
  if err != nil {
    return status, err
  }

  return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
  _, status, err := NewMarkReadPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 18(messages/mark_read)
    concat.Url(44, []string{
      SITE, "messages/mark_read",
    }), ids, is_read, MAX_EXPECTATION,
  )
  if err != nil {
    return status, err
  }

  return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Name: unread message type.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_MESSAGES,
// UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// Empty array to be filled with ids for messages.
func (f *FastId) UnreadMessagesIds(name string) ([]int, int, error) {
  var um api.UnreadMessages

  str_id := strconv.Itoa(f.Id)

  data, status, err := NewGetRequestWithCancelAndBearer(
    f.Conf.Application, f.Conf.AccessToken,
    // 26(SITE) + 6(users/) + ?(id) + 16(/unread_messages)
    concat.Url(48+len(str_id), []string{
      SITE, "users/", str_id, "/unread_messages",
    }), MAX_EXPECTATION,
  )
  if err != nil {
    return nil, status, err
  }

  if err := json.Unmarshal(data, &um); err != nil {
    return nil, status, err
  }

  switch name {
  case "messages":
    if um.Messages == 0 { return nil, status, errors.New("unread messages not found") }
    return make([]int, um.Messages), status, nil
  case "news":
    if um.News == 0 { return nil, status, errors.New("unread news not found") }
    return make([]int, um.News), status, nil
  case "notifications":
    if um.Notifications == 0 { return nil, status, errors.New("unread notifications not found") }
    return make([]int, um.Notifications), status, nil
  default:
    return nil, status, errors.New("wrong name... try messages, news or notifications")
  }
}

// In SetConfiguration(), you must specify the application and the token.
//
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
  _, status, err := NewReadDeleteAllPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 17(messages/read_all)
    concat.Url(43, []string{SITE, "messages/read_all"}),
    name, MAX_EXPECTATION,
  )
  if err != nil {
    return status, err
  }

  return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
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
  _, status, err := NewReadDeleteAllPostRequestWithCancel(
    c.Application, c.AccessToken,
    // 26(SITE) + 19(messages/delete_all)
    concat.Url(45, []string{SITE, "messages/delete_all"}),
    name, MAX_EXPECTATION,
  )
  if err != nil {
    return 0, err
  }

  return status, nil
}
