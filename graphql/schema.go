package graphql

import (
  "fmt"
  "errors"
  "strings"
)

// Available anime options:
//   - id malId name russian licenseNameRu english japanese synonyms kind rating score status episodes episodesAired duration airedOn{year month day date} releasedOn{year month day date} url season
//   - poster{id originalUrl mainUrl}
//   - fansubbers fandubbers licensors createdAt updatedAt nextEpisodeAt isCensored
//   - genres{id name russian kind}
//   - studios{id name imageUrl}
//   - personRoles{id rolesRu rolesEn person{id name poster{id}}}
//   - characterRoles{id rolesRu rolesEn character{id name poster{id}}}
//   - related{id anime{id name} manga{id name} relationRu relationEn}
//   - videos{id url name kind}
//   - screenshots{id originalUrl x166Url x332Url}
//   - scoresStats{score count}
//   - statusesStats{status count}
//   - description descriptionHtml descriptionSource
//
// Available manga options:
//   - id malId name russian licenseNameRu english japanese synonyms kind score status volumes chapters airedOn{year month day date} releasedOn{year month day date} url
//   - poster{id originalUrl mainUrl}
//   - licensors createdAt updatedAt isCensored
//   - genres{id name russian kind}
//   - publishers{id name}
//   - personRoles{id rolesRu rolesEn person{id name poster{id}}}
//   - characterRoles{id rolesRu rolesEn character{id name poster{id}}}
//   - related{id anime{id name} manga{id name} relationRu relationEn}
//   - scoresStats{score count}
//   - statusesStats{status count}
//   - description descriptionHtml descriptionSource
//
// Available character options:
//   - id malId name russian japanese synonyms url createdAt updatedAt isAnime isManga isRanobe
//   - poster{id originalUrl mainUrl}
//   - description descriptionHtml descriptionSource
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
// 	- Limit: 1;
//	- Score: 1;
//  - Order: empty field;
//	- Kind: empty field;
//	- Status: empty field;
//	- Season: empty field;
//	- Duration: empty field;
//	- Rating: empty field;
//	- Mylist: empty field;
//	- Censored: false;
//
// 'Options' settings:
//	- Limit: 50 maximum;
//	- Order: id, ranked, kind, popularity, name, aired_on, episodes, statust; random has been moved to a separate function, check [RandomAnime];
//	- Kind: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48, !tv, !movie, !ova, !ona, !special, !music, !tv_13, !tv_24, !tv_48;
//	- Status: anons, ongoing, released, !anons, !ongoing, !released;
//	- Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//	- Score: 1-9 maximum;
//  - Duration: S, D, F, !S, !D, !F;
//	- Rating: none, g, pg, pg_13, r, r_plus, rx, !g, !pg, !pg_13, !r, !r_plus, !rx;
//	- Censored: true, false;
//	- Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func AnimeSchema(values string, name string, options ...interface{}) (string, error) {
  var parameterOptions string

  for i, option := range options {
    switch i {
    case 0:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += fmt.Sprintf(", limit: %d", limit) }
    case 1:
      score, ok := option.(int)
      if ok && score >= 1 && score <= 9 { parameterOptions += fmt.Sprintf(", score: %d", score) }
    case 2:
      order, ok := option.(string)
      order_map := map[string]int8{
        "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
        "name": 5, "aired_on": 6, "episodes": 7, "status": 8,
      }
      _, ok_order := order_map[order]
      if ok && ok_order { parameterOptions += fmt.Sprintf(", order: %s", order) } // this parameter(string) without the quotation marks.
    case 3:
      kind, ok := option.(string)
      kind_map := map[string]int8{
        "tv": 1, "movie": 2, "ova": 3, "ona": 4, "special": 5, "music": 6,
        "tv_13": 7, "tv_24": 8, "tv_48": 9, "!tv": 10, "!movie": 11,
        "!ova": 12, "!ona": 13, "!special": 14, "!music": 15, "!tv_13": 16,
        "!tv_24": 17, "!tv_48": 18,
      }
      _, ok_kind := kind_map[kind]
      if ok && ok_kind { parameterOptions += fmt.Sprintf(`, kind: "%s"`, kind) }
    case 4:
      status, ok := option.(string)
      status_map := map[string]int8{
        "anons": 1, "ongoing": 2, "released": 3,
        "!anons": 4, "!ongoing": 5, "!released": 6,
      }
      _, ok_status := status_map[status]
      if ok && ok_status { parameterOptions += fmt.Sprintf(`, status: "%s"`, status) }
    case 5:
      season, ok := option.(string)
      season_map := map[string]int8{
        "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
        "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
        "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
        "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
      }
      _, ok_season := season_map[season]
      if ok && ok_season { parameterOptions += fmt.Sprintf(`, season: "%s"`, season) }
    case 6:
      duration, ok := option.(string)
      duration_map := map[string]int8{"S": 1, "D": 2, "F": 3, "!S": 4, "!D": 5, "!F": 6}
      _, ok_duration := duration_map[duration]
      if ok && ok_duration { parameterOptions += fmt.Sprintf(`, duration: "%s"`, duration) }
    case 7:
      rating, ok := option.(string)
      rating_map := map[string]int8{
        "none": 1, "g": 2, "pg": 3, "pg_13": 4,
        "r": 5, "r_plus": 6, "rx": 7, "!g": 8, "!pg": 9,
        "!pg_13": 10, "!r": 11, "!r_plus": 12, "!rx": 13,
      }
      _, ok_rating := rating_map[rating]
      if ok && ok_rating { parameterOptions += fmt.Sprintf(`, rating: "%s"`, rating) }
    case 8:
      mylist, ok := option.(string)
      mylist_map := map[string]int8{
        "planned": 1, "watching": 2, "rewatching": 3,
        "completed": 4, "on_hold": 5, "dropped": 6,
      }
      _, ok_mylist := mylist_map[mylist]
      if ok && ok_mylist { parameterOptions += fmt.Sprintf(`, mylist: "%s"`, mylist) }
    case 9:
      censored, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(`, censored: %t`, censored) }
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
// 	- Limit: 1;
//  - Order: empty field;
//	- Kind: empty field;
//	- Status: empty field;
//	- Season: empty field;
//	- Score: empty field;
//  - Censored: false;
//  - Mylist: empty field;
//
// 'Options' settings:
//	- Limit: 50 maximum;
//	- Order: id, ranked, kind, popularity, name, aired_on, volumes, chapters, status; random has been moved to a separate function, check [RandomManga];
//	- Kind: manga, manhwa, manhua, light_novel, novel, one_shot, doujin, !manga, !manhwa, !manhua, !light_novel, !novel, !one_shot, !doujin;
//	- Status: anons, ongoing, released, paused, discontinued, !anons, !ongoing, !released, !paused, !discontinued;
//	- Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//	- Score: 1-9 maximum;
//	- Censored: true, false;
//	- Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func MangaSchema(values string, name string, options ...interface{}) (string, error) {
  var parameterOptions string

  for i, option := range options {
    switch i {
    case 0:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += fmt.Sprintf(", limit: %d", limit) }
    case 1:
      score, ok := option.(int)
      if ok && score >= 1 && score <= 9 { parameterOptions += fmt.Sprintf(", score: %d", score) }
    case 2:
      order, ok := option.(string)
      order_map := map[string]int8{
        "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
        "name": 5, "aired_on": 6, "volumes": 7,
        "chapters": 8, "status": 9,
      }
      _, ok_order := order_map[order]
      if ok && ok_order { parameterOptions += fmt.Sprintf(", order: %s", order) } // this parameter(string) without the quotation marks.
    case 3:
      kind, ok := option.(string)
      kind_map := map[string]int8{
        "manga": 1, "manhwa": 2, "manhua": 3, "light_novel": 4, "novel": 5,
        "one_shot": 6, "doujin": 7, "!manga": 8, "!manhwa": 9, "!manhua": 10,
        "!light_novel": 11, "!novel": 12, "!one_shot": 13, "!doujin": 14,
      }
      _, ok_kind := kind_map[kind]
      if ok && ok_kind { parameterOptions += fmt.Sprintf(`, kind: "%s"`, kind) }
    case 4:
      status, ok := option.(string)
      status_map := map[string]int8{
        "anons": 1, "ongoing": 2, "released": 3, "paused": 4, "discontinued": 5,
        "!anons": 6, "!ongoing": 7, "!released": 8, "!paused": 9, "!discontinued": 10,
      }
      _, ok_status := status_map[status]
      if ok && ok_status { parameterOptions += fmt.Sprintf(`, status: "%s"`, status) }
    case 5:
      season, ok := option.(string)
      season_map := map[string]int8{
        "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
        "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
        "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
        "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
      }
      _, ok_season := season_map[season]
      if ok && ok_season { parameterOptions += fmt.Sprintf(`, season: "%s"`, season) }
    case 6:
      mylist, ok := option.(string)
      mylist_map := map[string]int8{
        "planned": 1, "watching": 2, "rewatching": 3,
        "completed": 4, "on_hold": 5, "dropped": 6,
      }
      _, ok_mylist := mylist_map[mylist]
      if ok && ok_mylist { parameterOptions += fmt.Sprintf(`, mylist: "%s"`, mylist) }
    case 7:
      censored, ok := option.(bool)
      if ok { parameterOptions += fmt.Sprintf(`, censored: %t`, censored) }
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
// 	- Limit: 1;
//
// 'Options' settings:
//  - Page: >= 1;
//	- Limit: 50 maximum;
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
