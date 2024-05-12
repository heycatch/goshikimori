package graphql

import (
  "fmt"
  "errors"
  "strings"
  "strconv"

  "github.com/heycatch/goshikimori/search"
  "github.com/heycatch/goshikimori/concat"
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
      if ok && page >= 1 { parameterOptions += ", page: " + strconv.Itoa(page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += ", limit: " + strconv.Itoa(limit) }
    case 2:
      score, ok := option.(int)
      if ok && score >= 1 && score <= 9 { parameterOptions += ", score: " + strconv.Itoa(score) }
    case 3:
      order, ok := option.(string)
      if ok && search.IndexInSlice(order, []string{
        "id", "ranked", "kind", "popularity", "name", "aired_on", "episodes", "status",
      }) != -1 { parameterOptions += ", order: " + order }
    case 4:
      kind, ok := option.(string)
      if ok && search.IndexInSlice(kind, []string{
        "tv", "movie", "ova", "ona", "special", "music",
        "tv_13", "tv_24", "tv_48", "!tv", "!movie",
        "!ova", "!ona", "!special", "!music", "!tv_13", "!tv_24", "!tv_48",
      }) != -1 { parameterOptions += ", kind: " + `"` + kind + `"` }
    case 5:
      status, ok := option.(string)
      if ok && search.IndexInSlice(status, []string{
        "anons", "ongoing", "released", "!anons", "!ongoing", "!released",
      }) != -1 { parameterOptions += ", status: " + `"` + status + `"` }
    case 6:
      season, ok := option.(string)
      if ok && search.IndexInSlice(season, []string{
        "2000_2010", "2010_2014", "2015_2019", "199x",
        "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
        "198x", "!198x", "2020_2021", "!2020_2021",
        "2022", "!2022", "2023", "!2023",
      }) != -1 { parameterOptions += ", season: " + `"` + season + `"` }
    case 7:
      duration, ok := option.(string)
      if ok && search.IndexInSlice(duration, []string{
        "S", "D", "F", "!S", "!D", "!F",
      }) != -1 { parameterOptions += ", duration: " + `"` + duration + `"` }
    case 8:
      rating, ok := option.(string)
      if ok && search.IndexInSlice(rating, []string{
        "none", "g", "pg", "pg_13", "r", "r_plus", "rx",
        "!g", "!pg", "!pg_13", "!r", "!r_plus", "!rx",
      }) != -1 { parameterOptions += ", rating: " + `"` + rating + `"` }
    case 9:
      mylist, ok := option.(string)
      if ok && search.IndexInSlice(mylist, []string{
        "planned", "watching", "rewatching", "completed", "on_hold", "dropped",
      }) != -1 { parameterOptions += ", mylist: " + `"` + mylist + `"` }
    case 10:
      censored, ok := option.(bool)
      if ok { parameterOptions += ", censored: " + strconv.FormatBool(censored) }
    case 11:
      genres_v2, ok_genre_v2 := option.([]int)
      genre := concat.MapGenresAnime(genres_v2)
      if ok_genre_v2 && genre != "" { parameterOptions += ", genre: " + `"` + genre + `"` }
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
      if ok && page >= 1 { parameterOptions += ", page: " + strconv.Itoa(page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += ", limit: " + strconv.Itoa(limit) }
    case 2:
      score, ok := option.(int)
      if ok && score >= 1 && score <= 9 { parameterOptions += ", score: " + strconv.Itoa(score) }
    case 3:
      order, ok := option.(string)
      if ok && search.IndexInSlice(order, []string{
        "id", "ranked", "kind", "popularity",
        "name", "aired_on", "volumes", "chapters", "status",
      }) != -1 { parameterOptions += ", order: " + order }
    case 4:
      kind, ok := option.(string)
      if ok && search.IndexInSlice(kind, []string {
        "manga", "manhwa", "manhua", "light_novel", "novel",
        "one_shot", "doujin", "!manga", "!manhwa", "!manhua",
        "!light_novel", "!novel", "!one_shot", "!doujin",
      }) != -1 { parameterOptions += ", kind: " + `"` + kind + `"` }
    case 5:
      status, ok := option.(string)
      if ok && search.IndexInSlice(status, []string{
        "anons", "ongoing", "released", "paused", "discontinued",
        "!anons", "!ongoing", "!released", "!paused", "!discontinued",
      }) != -1 { parameterOptions += ", status: " + `"` + status + `"` }
    case 6:
      season, ok := option.(string)
      if ok && search.IndexInSlice(season, []string{
        "2000_2010", "2010_2014", "2015_2019", "199x",
        "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
        "198x", "!198x", "2020_2021", "!2020_2021",
        "2022", "!2022", "2023", "!2023",
      }) != -1 { parameterOptions += ", season: " + `"` + season + `"` }
    case 7:
      mylist, ok := option.(string)
      if ok && search.IndexInSlice(mylist, []string{
        "planned", "watching", "rewatching", "completed", "on_hold", "dropped",
      }) != -1 { parameterOptions += ", mylist: " + `"` + mylist + `"` }
    case 8:
      censored, ok := option.(bool)
      if ok { parameterOptions += ", censored: " + strconv.FormatBool(censored) }
    case 9:
      genres_v2, ok_genre_v2 := option.([]int)
      genre := concat.MapGenresManga(genres_v2)
      if ok_genre_v2 && genre != "" { parameterOptions += ", genre: " + `"` + genre + `"` }
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
      if ok && page >= 1 { parameterOptions += ", page: " +  strconv.Itoa(page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += ", limit: " + strconv.Itoa(limit) }
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
      if ok && page >= 1 { parameterOptions += ", page: " + strconv.Itoa(page) }
    case 1:
      limit, ok := option.(int)
      if ok && limit >= 1 && limit <= 50 { parameterOptions += ", limit: " + strconv.Itoa(limit) }
    case 2:
      seyu, ok := option.(bool)
      if ok { parameterOptions += ", isSeyu: " + strconv.FormatBool(seyu) }
    case 3:
      mangaka, ok := option.(bool)
      if ok { parameterOptions += ", isMangaka: " + strconv.FormatBool(mangaka) }
    case 4:
      producer, ok := option.(bool)
      if ok { parameterOptions += ", isProducer: " + strconv.FormatBool(producer) }
    default:
      return "", errors.New("one of the parameters is entered incorrectly, check sequence or spelling errors")
    }
  }

  return fmt.Sprintf(`graphql?query={people(search: "%s"%s){%s}}`, name, parameterOptions, values), nil
}
