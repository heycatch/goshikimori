package graphql

import (
  "fmt"
  "errors"
)

func AnimeSchema(name string, options ...interface{}) (string, error) {
  var parameterOptions, parameterRequests string

  values := []string{
    "id malId name russian licenseNameRu english japanese synonyms kind rating score status episodes episodesAired duration airedOn{year month day date} releasedOn{year month day date} url ",
    "poster{id originalUrl mainUrl} ",
    "fansubbers fandubbers licensors createdAt updatedAt isCensored ",
    "genres{id name russian kind} ",
    "studios{id name imageUrl} ",
    "personRoles{id rolesRu rolesEn person{id name poster{id}}} ",
    "characterRoles{id rolesRu rolesEn character{id name poster{id}}} ",
    "related{id anime{id name} manga{id name} relationRu relationEn} ",
    "videos{id url name kind} ",
    "screenshots{id originalUrl x166Url x332Url} ",
    "scoresStats{score count} ",
    "statusesStats{status count} ",
    "description descriptionHtml descriptionSource",
  }

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

  for i := 0; i < len(values); i++ { parameterRequests += values[i] }

  return fmt.Sprintf(`graphql?query={animes(search: "%s"%s){%s}}`, name, parameterOptions, parameterRequests), nil
}

func MangaSchema(name string, options ...interface{}) (string, error) {
  var parameterOptions, parameterRequests string

  values := []string{
    "id malId name russian licenseNameRu english japanese synonyms kind score status volumes chapters airedOn{year month day date} releasedOn{year month day date} url ",
    "poster{id originalUrl mainUrl} ",
    "licensors createdAt updatedAt isCensored ",
    "genres{id name russian kind} ",
    "publishers{id name} ",
    "personRoles{id rolesRu rolesEn person{id name poster{id}}} ",
    "characterRoles{id rolesRu rolesEn character{id name poster{id}}} ",
    "related{id anime{id name} manga{id name} relationRu relationEn} ",
    "scoresStats{score count} ",
    "statusesStats{status count} ",
    "description descriptionHtml descriptionSource",
  }

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

  for i := 0; i < len(values); i++ { parameterRequests += values[i] }

  return fmt.Sprintf(`graphql?query={mangas(search: "%s"%s){%s}}`, name, parameterOptions, parameterRequests), nil
}
