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
  "net/url"
  "errors"

  "github.com/heycatch/goshikimori/api"
  "github.com/heycatch/goshikimori/str"
  "github.com/heycatch/goshikimori/req"
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "users/" + url.QueryEscape(name), 10,
  )
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

// Name: user name.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "users?search=" + url.QueryEscape(name) + "&" + r.OptionsUsers(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUser(f.Id, "friends?" + r.OptionsUsers()), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUser(f.Id, "clubs"), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Status: watching;
//  - Censored: false;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 5000 maximum;
//  - Status: planned, watching, rewatching, completed, on_hold, dropped;
//  - Censored: true, false;
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserAnimeRates(r Result) ([]api.UserAnimeRates, error) {
  var ar []api.UserAnimeRates
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUserRates(f.Id, "anime_rates", r.OptionsAnimeRates()), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Censored: false;
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUserRates(f.Id, "manga_rates", r.OptionsMangaRates()), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUser(f.Id, "favourites"), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Target_type: Anime;
//  - Target_id: option is hidden if empty;
//
// 'Options' settings:
//  - Page: 100000 maximum.
//  - Limit: 100 maximum.
//  - Target_id: id anime/manga in string format.
//  - Target_type: Anime, Manga.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserHistory(r Result) ([]api.UserHistory, error) {
  var uh []api.UserHistory
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUserRates(f.Id, "history", r.OptionsUserHistory()), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUser(f.Id, "bans"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "users/whoami", 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertSearchById("animes", f.Id), 10,
  )
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
// Exclamation mark(!) indicates ignore.
//
// If you use the 'order' parameter, you don't need to enter the name of the anime.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Order: empty field;
//  - Kind: empty field;
//  - Status: empty field;
//  - Season: empty field;
//  - Score: empty field;
//  - Duration: empty field;
//  - Rating: empty field;
//  - Censored: false;
//  - Mylist: empty field;
//  - Genre_v2: empty field;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 50 maximum;
//  - Order: id, ranked, kind, popularity, name, aired_on, episodes, status; random has been moved to a separate function, check [RandomAnime];
//  - Kind: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48, !tv, !movie, !ova, !ona, !special, !music, !tv_13, !tv_24, !tv_48;
//  - Status: anons, ongoing, released, !anons, !ongoing, !released;
//  - Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//  - Score: 1-9 maximum;
//  - Duration: S, D, F, !S, !D, !F;
//  - Rating: none, g, pg, pg_13, r, r_plus, rx, !g, !pg, !pg_13, !r, !r_plus, !rx;
//  - Censored: true(string), false(string);
//  - Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//  - Search: default search;
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
// [RandomAnime]: https://github.com/heycatch/goshikimori/blob/master/examples/random
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchAnimes(name string, r Result) ([]api.Animes, int, error) {
  var a []api.Animes
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "animes?search=" + url.QueryEscape(name) + "&" + r.OptionsAnime(), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertSearchById("mangas", f.Id), 10,
  )
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
// Exclamation mark(!) indicates ignore.
//
// If you use the 'order' parameter, you don't need to enter the name of the manga.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Order: empty field;
//  - Kind: empty field;
//  - Status: empty field;
//  - Season: empty field;
//  - Score: empty field;
//  - Censored: false;
//  - Mylist: empty field;
//  - Genre_v2: empty field;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 50 maximum;
//  - Order: id, ranked, kind, popularity, name, aired_on, volumes, chapters, status; random has been moved to a separate function, check [RandomManga];
//  - Kind: manga, manhwa, manhua, light_novel, novel, one_shot, doujin, !manga, !manhwa, !manhua, !light_novel, !novel, !one_shot, !doujin;
//  - Status: anons, ongoing, released, paused, discontinued, !anons, !ongoing, !released, !paused, !discontinued;
//  - Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//  - Score: 1-9 maximum;
//  - Censored: true(string), false(string);
//  - Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//  - Search: default search;
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
// [RandomManga]: https://github.com/heycatch/goshikimori/blob/master/examples/random
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchMangas(name string, r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "mangas?search=" + url.QueryEscape(name) + "&" + r.OptionsManga(), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertSearchById("ranobe", f.Id), 10,
  )
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
// Exclamation mark(!) indicates ignore.
//
// If you use the 'order' parameter, you don't need to enter the name of the ranobe.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Order: empty field;
//  - Status: empty field;
//  - Season: empty field;
//  - Score: empty field;
//  - Censored: false;
//  - Mylist: empty field;
//  - Genre_v2: empty field;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 50 maximum;
//  - Order: id, ranked, kind, popularity, name, aired_on, volumes, chapters, status; random has been moved to a separate function, check [RandomRanobe];
//  - Status: anons, ongoing, released, paused, discontinued, !anons, !ongoing, !released, !paused, !discontinued;
//  - Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//  - Score: 1-9 maximum;
//  - Censored: true(string), false(string);
//  - Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//  - Search: default search;
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
// [RandomRanobe]: https://github.com/heycatch/goshikimori/blob/master/examples/random
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchRanobes(name string, r Result) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "ranobe?search=" + url.QueryEscape(name) + "&" + r.OptionsRanobe(), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "users/" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "animes?search=" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "mangas?search=" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "ranobe?search=" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "clubs?search=" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "characters/search?search=" + url.QueryEscape(name), 10,
  )
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
// Search by user is case sensitive.
func (c *Configuration) FastIdPeople(name string) (*FastId, int, error) {
  var ap []api.AllPeople
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "people/search?search=" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertAnime(f.Id, "screenshots"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFranchise(f.Id, "animes"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFranchise(f.Id, "mangas"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFranchise(f.Id, "ranobe"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertExternalLinks(f.Id, "animes"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertExternalLinks(f.Id, "mangas"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertSimilar(f.Id, "animes"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertSimilar(f.Id, "mangas"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertSimilar(f.Id, "ranobe"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertRelated(f.Id, "animes"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertRelated(f.Id, "mangas"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertRelated(f.Id, "ranobe"), 10,
  )
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

// Name: club name.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//  - Search: default search;
//
// If we set the limit=1, we will still have 2 results.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (c *Configuration) SearchClubs(name string, r Result) ([]api.Clubs, int, error) {
  var cl []api.Clubs
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "clubs?search=" + url.QueryEscape(name) + "&" + r.OptionsClub(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubAnimes(r Result) ([]api.Animes, error) {
  var a []api.Animes
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "animes") + "?" + r.OptionsClubInformation(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMangas(r Result) ([]api.Mangas, error) {
  var m []api.Mangas
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "mangas") + "?" + r.OptionsClubInformation(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCharacters(r Result) ([]api.CharacterInfo, error) {
  var ci []api.CharacterInfo
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "characters") + "?" + r.OptionsClubInformation(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubClubs(r Result) ([]api.Clubs, error) {
  var cc []api.Clubs
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "clubs") + "?" + r.OptionsClubInformation(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCollections(r Result) ([]api.ClubCollections, error) {
  var cc []api.ClubCollections
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "collections") + "?" + r.OptionsClubInformation(), 10,
  )
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMembers() ([]api.UserFriends, error) {
  var uf []api.UserFriends
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "members"), 10,
  )
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
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubImages() ([]api.ClubImages, error) {
  var cm []api.ClubImages
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "images"), 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "join"), 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertClub(f.Id, "leave"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertAchievements(f.Id), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertAnime(f.Id, "videos"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertRoles(f.Id, "animes"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertRoles(f.Id, "mangas"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, "bans", 10,
  )
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

// If 'Options' empty fields:
//  - Censored: false;
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
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertCalendar(r.OptionsCalendar()), 10,
  )
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

// name: anime or manga.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/genres
func (c *Configuration) SearchGenres(name string) ([]api.Genres, int, error) {
  var g []api.Genres
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, str.ConvertGenres(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, "studios", 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, "publishers", 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, "forums", 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFriend(f.Id), 10,
  )
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

  remove, cancel := req.NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFriend(f.Id), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUser(f.Id, "unread_messages"), 10,
  )
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

// If 'Options' empty fields:
//  - Type: news;
//  - Page: 1;
//  - Limit: 1;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 100 maximum;
//  - Type: inbox, private, sent, news, notifications;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/messages
func (f *FastId) UserMessages(r Result) ([]api.Messages, error) {
  var m []api.Messages
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertMessages(f.Id, r.OptionsMessages()), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertConstants("anime"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertConstants("manga"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertConstants("user_rate"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertConstants("club"), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertConstants("smileys"), 10,
  )
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

// Limit: number of results obtained;
//
// Minimum: 1; Maximum: 50;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/random
func (c *Configuration) RandomAnimes(limit int) ([]api.Animes, int, error) {
  var a []api.Animes
  var client = &http.Client{}

  if limit < 1 || limit > 50 { limit = 1 }

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, str.ConvertRandom("animes", limit), 10,
  )
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

// Limit: number of results obtained;
//
// Minimum: 1; Maximum: 50;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/random
func (c *Configuration) RandomMangas(limit int) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  if limit < 1 || limit > 50 { limit = 1 }

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, str.ConvertRandom("mangas", limit), 10,
  )
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

// Limit: number of results obtained;
//
// Minimum: 1; Maximum: 50;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/random
func (c *Configuration) RandomRanobes(limit int) ([]api.Mangas, int, error) {
  var m []api.Mangas
  var client = &http.Client{}

  if limit < 1 || limit > 50 { limit = 1 }

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, str.ConvertRandom("ranobe", limit), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertCharacters(f.Id), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "characters/search?search=" + url.QueryEscape(name), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertPeople(f.Id), 10,
  )
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

// Name: people name.
//
// If 'Options' empty fields:
//  - Kind: seyu;
//
// 'Options' settings:
//  - Page/Limit: not supported, idk why;
//  - Kind: seyu, mangaka, producer;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/people
func (c *Configuration) SearchPeoples(name string, r Result) ([]api.AllPeople, int, error) {
  var ap []api.AllPeople
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "people/search?search=" + url.QueryEscape(name) + "&" + r.OptionsPeople(), 10,
  )
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

// Linked_type: Anime, Manga, Ranobe, Person, Character.
//
// Kind(required when Linked_type is Person): common, seyu, mangaka, producer, person.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesCreate(linked_type string, kind string) (api.Favorites, error) {
  var fa api.Favorites
  var client = &http.Client{}

  type_map := map[string]int8{"Anime": 1, "Manga": 2, "Ranobe": 3, "Person": 4, "Character": 5}
  _, ok_type := type_map[linked_type]
  if !ok_type { return fa, errors.New("incorrect string, try again and watch the upper case") }

  kind_map := map[string]int8{"common": 1, "seyu": 2, "mangaka": 3, "producer": 4, "person": 5}
  _, ok_kind := kind_map[kind]
  if !ok_kind { kind = "" }

  post, cancel := req.NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFavorites(linked_type, f.Id, kind), 10,
  )
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

// Linked_type: Anime, Manga, Ranobe, Person, Character.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesDelete(linked_type string) (api.Favorites, error) {
  var ff api.Favorites
  var client = &http.Client{}

  type_map := map[string]int8{"Anime": 1, "Manga": 2, "Ranobe": 3, "Person": 4, "Character": 5}
  _, ok_type := type_map[linked_type]
  if !ok_type { return ff, errors.New("incorrect string, try again and watch the upper case") }

  remove, cancel := req.NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFavorites(linked_type, f.Id, ""), 10,
  )
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

  post, cancel := req.NewReorderPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertFavoritesReorder(f.Id), position, 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertIgnoreUser(f.Id), 10,
  )
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

  remove, cancel := req.NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertIgnoreUser(f.Id), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken, "dialogs", 10,
  )
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (f *FastId) SearchDialogs() ([]api.SearchDialogs, error) {
  var sd []api.SearchDialogs
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertDialogs(f.Id), 10,
  )
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/dialogs
func (f *FastId) DeleteDialogs() (api.FriendRequest, error) {
  var fr api.FriendRequest
  var client = &http.Client{}

  remove, cancel := req.NewDeleteRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertDialogs(f.Id), 10,
  )
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

  // errors.New(...) original error message from api/v1.
  if err := json.Unmarshal(data, &fr); err != nil {
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUserBriefInfo(f.Id), 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    "users/sign_out", 10,
  )
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
  var ids []int
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "stats/active_users", 40,
  )
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return ids, resp.StatusCode, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return ids, resp.StatusCode, err
  }

  if err := json.Unmarshal(data, &ids); err != nil {
    return ids, resp.StatusCode, err
  }

  return ids, resp.StatusCode, nil
}

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertTopicsType(f.Id, "animes") + "?" + r.OptionsClub(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertTopicsType(f.Id, "mangas") + "?" + r.OptionsClub(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//
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

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertTopicsType(f.Id, "ranobe") + "?" + r.OptionsClub(), 10,
  )
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

// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Forum: all;
//  - Linked_id: empty field;
//  - Linked_type: empty field;
//
// 'Options' settings:
//  - Page: 100000 maximum;
//  - Limit: 30 maximum;
//  - Forum: cosplay, animanga, site, games, vn, contests, offtopic, clubs, my_clubs, critiques, news, collections, articles;
//  - Linked_id: number without limit;
//  - Linked_type: Anime, Manga, Ranobe, Character, Person, Club, ClubPage, Critique, Review, Contest, CosplayGallery, Collection, Article;
//
// REMARK: linked_id and linked_type are only used together.
//
//  - Type: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
//
// TODO: Add implementation of creating/updating/deleting a topic. This is not needed at this stage.
func (c *Configuration) SearchTopics(r Result) ([]api.Topics, int, error) {
  var t []api.Topics
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "topics?" + r.OptionsTopics(), 10,
  )
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsUpdates(r Result) ([]api.TopicsUpdates, int, error) {
  var t []api.TopicsUpdates
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "topics/updates?" + r.OptionsClub(), 10,
  )
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

// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsHot(r Result) ([]api.Topics, int, error) {
  var t []api.Topics
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    "topics/hot?" + r.OptionsTopicsHot(), 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertTopicsId(id), 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertIgnoreTopic(id), 10,
  )
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

  remove, cancel := req.NewDeleteRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertIgnoreTopic(id), 10,
  )
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

  post, cancel := req.NewPostRequestWithCancel(
    c.Application, c.AccessToken, schema, 10,
  )
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

  get, cancel := req.NewGetRequestWithCancel(
    c.Application, c.AccessToken,
    str.ConvertMessage(id), 10,
  )
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

  post, cancel := req.NewSendMessagePostRequestWithCancel(
    c.Application, c.AccessToken, "messages", message, from_id, to_id, 10,
  )
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

  put, cancel := req.NewChangeMessagePutRequestWithCancel(
    c.Application, c.AccessToken, str.ConvertMessage(id), message, 10,
  )
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

  del, cancel := req.NewDeleteMessageDeleteRequestWithCancel(
    c.Application, c.AccessToken, str.ConvertMessage(id), 10,
  )
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

  post, cancel := req.NewMarkReadPostRequestWithCancel(
    c.Application, c.AccessToken, "messages/mark_read", ids, is_read, 10,
  )
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
//  - news
//  - notifications
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) ReadAllMessages(name string) (int, error) {
  var client = &http.Client{}

  post, cancel := req.NewReadDeleteAllPostRequestWithCancel(
    c.Application, c.AccessToken, "messages/read_all", name, 10,
  )
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
//  - news
//  - notifications
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/message
func (c *Configuration) DeleteAllMessages(name string) (int, error) {
  var client = &http.Client{}

  post, cancel := req.NewReadDeleteAllPostRequestWithCancel(
    c.Application, c.AccessToken, "messages/delete_all", name, 10,
  )
  defer cancel()

  resp, err := client.Do(post)
  if err != nil {
    return resp.StatusCode, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil
}
