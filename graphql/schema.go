package graphql

import (
  "fmt"
  "errors"
  "strings"
)

// Available anime options:
//  - id malId name russian licenseNameRu english japanese synonyms kind rating score status episodes episodesAired duration airedOn{year month day date} releasedOn{year month day date} url season
//  - poster{id originalUrl mainUrl}
//  - fansubbers fandubbers licensors createdAt updatedAt nextEpisodeAt isCensored
//  - genres{id name russian kind}
//  - studios{id name imageUrl}
//  - personRoles{id rolesRu rolesEn person{id name poster{id}}}
//  - characterRoles{id rolesRu rolesEn character{id name poster{id}}}
//  - related{id anime{id name} manga{id name} relationRu relationEn}
//  - videos{id url name kind}
//  - screenshots{id originalUrl x166Url x332Url}
//  - scoresStats{score count}
//  - statusesStats{status count}
//  - description descriptionHtml descriptionSource
//
// Available manga options:
//  - id malId name russian licenseNameRu english japanese synonyms kind score status volumes chapters airedOn{year month day date} releasedOn{year month day date} url
//  - poster{id originalUrl mainUrl}
//  - licensors createdAt updatedAt isCensored
//  - genres{id name russian kind}
//  - publishers{id name}
//  - personRoles{id rolesRu rolesEn person{id name poster{id}}}
//  - characterRoles{id rolesRu rolesEn character{id name poster{id}}}
//  - related{id anime{id name} manga{id name} relationRu relationEn}
//  - scoresStats{score count}
//  - statusesStats{status count}
//  - description descriptionHtml descriptionSource
//
// Available character options:
//  - id malId name russian japanese synonyms url createdAt updatedAt isAnime isManga isRanobe
//  - poster{id originalUrl mainUrl}
//  - description descriptionHtml descriptionSource
//
// Available people options:
//  - id malId name russian japanese synonyms url isSeyu isMangaka isProducer website createdAt updatedAt
//  - birthOn{year month day date} deceasedOn{year month day date}
//  - poster{id originalUrl mainUrl}
func Values(input ...string) string {
  var res string

  // We always return 1, even if the slice is empty.
  // But we are not allowed to return an empty value,
  // so an additional check is present and returns id.
  if len(input) == 1 && len(strings.TrimSpace(input[0])) == 0 { return "id" }

  // TODO: add keyword checks and the keyword is "full" to add everything.
  for i := 0; i < len(input); i++ { res += input[i] + " " }

  return res
}

// Values: parameters we want to receive from the server.
//
// Name: anime name.
//
// If you use the 'order' parameter, you don't need to enter the name of the anime.
//
// Exclamation mark(!) indicates ignore.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - Score: 1;
//  - Order: empty field;
//  - Kind: empty field;
//  - Status: empty field;
//  - Season: empty field;
//  - Duration: empty field;
//  - Rating: empty field;
//  - Mylist: empty field;
//  - Censored: false;
//  - Genre_v2: empty field;
//
// 'Options' settings:
//  - Page: >= 1;
//  - Limit: 50 maximum;
//  - Order: id, ranked, kind, popularity, name, aired_on, episodes, statust; random has been moved to a separate function, check [RandomAnime];
//  - Kind: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48, !tv, !movie, !ova, !ona, !special, !music, !tv_13, !tv_24, !tv_48;
//  - Status: anons, ongoing, released, !anons, !ongoing, !released;
//  - Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//  - Score: 1-9 maximum;
//  - Duration: S, D, F, !S, !D, !F;
//  - Rating: none, g, pg, pg_13, r, r_plus, rx, !g, !pg, !pg_13, !r, !r_plus, !rx;
//  - Censored: true, false;
//  - Mylist: planned, watching, rewatching, completed, on_hold, dropped;
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
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func AnimeSchema(values string, name string, options ...interface{}) (string, error) {
  var parameterOptions string

  for i, option := range options {
    switch i {
    case 0:
      page, ok := option.(int)
      if ok && page >= 1 { parameterOptions += fmt.Sprintf(", page: %d", page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += fmt.Sprintf(", limit: %d", limit) }
    case 2:
      score, ok := option.(int)
      if ok && score >= 1 && score <= 9 { parameterOptions += fmt.Sprintf(", score: %d", score) }
    case 3:
      order, ok := option.(string)
      order_map := map[string]int8{
        "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
        "name": 5, "aired_on": 6, "episodes": 7, "status": 8,
      }
      _, ok_order := order_map[order]
      if ok && ok_order { parameterOptions += fmt.Sprintf(", order: %s", order) } // this parameter(string) without the quotation marks.
    case 4:
      kind, ok := option.(string)
      kind_map := map[string]int8{
        "tv": 1, "movie": 2, "ova": 3, "ona": 4, "special": 5, "music": 6,
        "tv_13": 7, "tv_24": 8, "tv_48": 9, "!tv": 10, "!movie": 11,
        "!ova": 12, "!ona": 13, "!special": 14, "!music": 15, "!tv_13": 16,
        "!tv_24": 17, "!tv_48": 18,
      }
      _, ok_kind := kind_map[kind]
      if ok && ok_kind { parameterOptions += fmt.Sprintf(`, kind: "%s"`, kind) }
    case 5:
      status, ok := option.(string)
      status_map := map[string]int8{
        "anons": 1, "ongoing": 2, "released": 3,
        "!anons": 4, "!ongoing": 5, "!released": 6,
      }
      _, ok_status := status_map[status]
      if ok && ok_status { parameterOptions += fmt.Sprintf(`, status: "%s"`, status) }
    case 6:
      season, ok := option.(string)
      season_map := map[string]int8{
        "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
        "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
        "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
        "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
      }
      _, ok_season := season_map[season]
      if ok && ok_season { parameterOptions += fmt.Sprintf(`, season: "%s"`, season) }
    case 7:
      duration, ok := option.(string)
      duration_map := map[string]int8{"S": 1, "D": 2, "F": 3, "!S": 4, "!D": 5, "!F": 6}
      _, ok_duration := duration_map[duration]
      if ok && ok_duration { parameterOptions += fmt.Sprintf(`, duration: "%s"`, duration) }
    case 8:
      rating, ok := option.(string)
      rating_map := map[string]int8{
        "none": 1, "g": 2, "pg": 3, "pg_13": 4,
        "r": 5, "r_plus": 6, "rx": 7, "!g": 8, "!pg": 9,
        "!pg_13": 10, "!r": 11, "!r_plus": 12, "!rx": 13,
      }
      _, ok_rating := rating_map[rating]
      if ok && ok_rating { parameterOptions += fmt.Sprintf(`, rating: "%s"`, rating) }
    case 9:
      mylist, ok := option.(string)
      mylist_map := map[string]int8{
        "planned": 1, "watching": 2, "rewatching": 3,
        "completed": 4, "on_hold": 5, "dropped": 6,
      }
      _, ok_mylist := mylist_map[mylist]
      if ok && ok_mylist { parameterOptions += fmt.Sprintf(`, mylist: "%s"`, mylist) }
    case 10:
      censored, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(`, censored: %t`, censored) }
    case 11:
      genres_v2, ok_genre_v2 := option.([]int)
      genres := map[int]string{
        1: "1-Action", 2: "2-Adventure", 3: "3-Cars", 4: "4-Comedy",
        5: "5-Dementia", 6: "6-Demons", 7: "7-Mystery", 8: "8-Drama",
        9: "9-Ecchi", 10: "10-Fantasy", 11: "11-Game", 12: "12-Hentai",
        13: "13-Historical", 14: "14-Horror", 15: "15-Kids", 16: "16-Magic",
        17: "17-Martial Arts", 18: "18-Mecha", 19: "19-Music", 20: "20-Parody",
        21: "21-Samurai", 22: "22-Romance", 23: "23-School",
        24: "24-Sci-Fi", 25: "25-Shoujo", 26: "26-Shoujo Ai", 27: "27-Shounen",
        28: "28-Shounen Ai", 29: "29-Space", 30: "30-Sports", 31: "31-Super Power",
        32: "32-Vampire", 33: "33-Yaoi", 34: "34-Yuri", 35: "35-Harem",
        36: "36-Slice of Life", 37: "37-Supernatural", 38: "38-Military",
        39: "39-Police", 40: "40-Psychological", 41: "41-Thriller",
        42: "42-Seinen", 43: "43-Josei",
        539: "539-Erotica", 541: "541-Work Life", 543: "543-Gourmet",
      }
      var genre_v2 string
      for i := 0; i < len(genres_v2); i++ {
        _, ok := genres[genres_v2[i]]; if ok {
          genre_v2 += genres[genres_v2[i]] + ","
        }
      }
      if len(genre_v2) != 0 { genre_v2 = genre_v2[:len(genre_v2)-1] }
      if ok_genre_v2 { parameterOptions += fmt.Sprintf(`, genre: "%s"`, genre_v2) }
    default:
      return "", errors.New("one of the parameters is entered incorrectly, check sequence or spelling errors")
    }
  }

  return fmt.Sprintf(`graphql?query={animes(search: "%s"%s){%s}}`, name, parameterOptions, values), nil
}


// Values: parameters we want to receive from the server.
//
// Name: manga name.
//
// If you use the 'order' parameter, you don't need to enter the name of the anime.
//
// Exclamation mark(!) indicates ignore.
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
//  - Page: >= 1;
//  - Limit: 50 maximum;
//  - Order: id, ranked, kind, popularity, name, aired_on, volumes, chapters, status; random has been moved to a separate function, check [RandomManga];
//  - Kind: manga, manhwa, manhua, light_novel, novel, one_shot, doujin, !manga, !manhwa, !manhua, !light_novel, !novel, !one_shot, !doujin;
//  - Status: anons, ongoing, released, paused, discontinued, !anons, !ongoing, !released, !paused, !discontinued;
//  - Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//  - Score: 1-9 maximum;
//  - Censored: true, false;
//  - Mylist: planned, watching, rewatching, completed, on_hold, dropped;
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
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func MangaSchema(values string, name string, options ...interface{}) (string, error) {
  var parameterOptions string

  for i, option := range options {
    switch i {
    case 0:
      page, ok := option.(int)
      if ok && page >= 1 { parameterOptions += fmt.Sprintf(", page: %d", page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += fmt.Sprintf(", limit: %d", limit) }
    case 2:
      score, ok := option.(int)
      if ok && score >= 1 && score <= 9 { parameterOptions += fmt.Sprintf(", score: %d", score) }
    case 3:
      order, ok := option.(string)
      order_map := map[string]int8{
        "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
        "name": 5, "aired_on": 6, "volumes": 7,
        "chapters": 8, "status": 9,
      }
      _, ok_order := order_map[order]
      if ok && ok_order { parameterOptions += fmt.Sprintf(", order: %s", order) } // this parameter(string) without the quotation marks.
    case 4:
      kind, ok := option.(string)
      kind_map := map[string]int8{
        "manga": 1, "manhwa": 2, "manhua": 3, "light_novel": 4, "novel": 5,
        "one_shot": 6, "doujin": 7, "!manga": 8, "!manhwa": 9, "!manhua": 10,
        "!light_novel": 11, "!novel": 12, "!one_shot": 13, "!doujin": 14,
      }
      _, ok_kind := kind_map[kind]
      if ok && ok_kind { parameterOptions += fmt.Sprintf(`, kind: "%s"`, kind) }
    case 5:
      status, ok := option.(string)
      status_map := map[string]int8{
        "anons": 1, "ongoing": 2, "released": 3, "paused": 4, "discontinued": 5,
        "!anons": 6, "!ongoing": 7, "!released": 8, "!paused": 9, "!discontinued": 10,
      }
      _, ok_status := status_map[status]
      if ok && ok_status { parameterOptions += fmt.Sprintf(`, status: "%s"`, status) }
    case 6:
      season, ok := option.(string)
      season_map := map[string]int8{
        "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
        "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
        "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
        "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
      }
      _, ok_season := season_map[season]
      if ok && ok_season { parameterOptions += fmt.Sprintf(`, season: "%s"`, season) }
    case 7:
      mylist, ok := option.(string)
      mylist_map := map[string]int8{
        "planned": 1, "watching": 2, "rewatching": 3,
        "completed": 4, "on_hold": 5, "dropped": 6,
      }
      _, ok_mylist := mylist_map[mylist]
      if ok && ok_mylist { parameterOptions += fmt.Sprintf(`, mylist: "%s"`, mylist) }
    case 8:
      censored, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(`, censored: %t`, censored) }
    case 9:
      genres_v2, ok_genre_v2 := option.([]int)
      genres := map[int]string{
        46: "46-Mystery", 47: "47-Shounen",
        48: "48-Supernatural", 49: "49-Comedy", 50: "50-Drama", 51: "51-Ecchi",
        52: "52-Seinen", 53: "53-Sci-Fi", 54: "54-Slice of Life", 55: "55-Shounen Ai",
        56: "56-Action", 57: "57-Fantasy", 58: "58-Magic", 59: "59-Hentai",
        60: "60-School", 61: "61-Doujinshi", 62: "62-Romance", 63: "63-Shoujo",
        64: "64-Vampire", 65: "65-Yaoi", 66: "66-Martial Arts", 67: "67-Psychological",
        68: "68-Adventure", 69: "69-Historical", 70: "70-Military", 71: "71-Harem",
        72: "72-Demons", 73: "73-Shoujo Ai", 74: "74-Gender Bender", 75: "75-Yuri",
        76: "76-Sports", 77: "77-Kids", 78: "78-Music", 79: "79-Game", 80: "80-Horror",
        81: "81-Thriller", 82: "82-Super Power", 83: "83-Mecha", 84: "84-Cars",
        85: "85-Space", 86: "86-Parody", 87: "87-Josei", 88: "88-Samurai", 89: "89-Police",
        90: "90-Dementia", 540: "540-Erotica", 542: "542-Work Life", 544: "544-Gourmet",
      }
      var genre_v2 string
      for i := 0; i < len(genres_v2); i++ {
        _, ok := genres[genres_v2[i]]; if ok {
          genre_v2 += genres[genres_v2[i]] + ","
        }
      }
      if len(genre_v2) != 0 { genre_v2 = genre_v2[:len(genre_v2)-1] }
      if ok_genre_v2 { parameterOptions += fmt.Sprintf(`, genre: "%s"`, genre_v2) }
    default:
      return "", errors.New("one of the parameters is entered incorrectly, check sequence or spelling errors")
    }
  }

  return fmt.Sprintf(`graphql?query={mangas(search: "%s"%s){%s}}`, name, parameterOptions, values), nil
}

// Values: parameters we want to receive from the server.
//
// Name: character name.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//
// 'Options' settings:
//  - Page: >= 1;
//  - Limit: 50 maximum;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func CharacterSchema(values string, name string, options ...interface{}) (string, error) {
  var parameterOptions string

  for i, option := range options {
    switch i {
    case 0:
      page, ok := option.(int)
      if ok && page >= 1 { parameterOptions += fmt.Sprintf(", page: %d", page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += fmt.Sprintf(", limit: %d", limit) }
    default:
      return "", errors.New("one of the parameters is entered incorrectly, check sequence or spelling errors")
    }
  }

  return fmt.Sprintf(`graphql?query={characters(search: "%s"%s){%s}}`, name, parameterOptions, values), nil
}

// Values: parameters we want to receive from the server.
//
// Name: people name.
//
// If 'Options' empty fields:
//  - Page: 1;
//  - Limit: 1;
//  - isSeyu: true;
//  - isMangaka: true;
//  - isProducer: true;
//
// 'Options' settings:
//  - Page: >= 1;
//  - Limit: 50 maximum;
//  - isSeyu: true, false;
//  - isMangaka: true, false;
//  - isProducer: true, false;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func PeopleSchema(values string, name string, options ...interface{}) (string, error) {
  var parameterOptions string

  for i, option := range options {
    switch i {
    case 0:
      page, ok := option.(int)
      if ok && page >= 1 { parameterOptions += fmt.Sprintf(", page: %d", page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += fmt.Sprintf(", limit: %d", limit) }
    case 2:
      seyu, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(", isSeyu: %t", seyu) }
    case 3:
      mangaka, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(", isMangaka: %t", mangaka) }
    case 4:
      producer, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(", isProducer: %t", producer) }
    default:
      return "", errors.New("one of the parameters is entered incorrectly, check sequence or spelling errors")
    }
  }

  return fmt.Sprintf(`graphql?query={people(search: "%s"%s){%s}}`, name, parameterOptions, values), nil
}
