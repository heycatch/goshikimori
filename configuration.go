package goshikimori

import (
  "net/url"
  "strconv"

  "github.com/heycatch/goshikimori/search"
  "github.com/heycatch/goshikimori/concat"
)

type Configuration struct {
  Application, AccessToken string
}

type FastId struct {
  Id   int
  Conf Configuration
  Err  error
}

// Getting an id(anime, manga, ranobe, user, person, group).
//
// More information can be found in the [example1] and [example2].
//
// [example1]: https://github.com/heycatch/goshikimori/blob/master/examples/custom_fastid
// [example2]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func (f *FastId) GetFastId() int { return f.Id }

// To create a custom id(anime, manga, ranobe, user, person, group).
//
// More information can be found in the [example1] and [example2].
//
// [example1]: https://github.com/heycatch/goshikimori/blob/master/examples/custom_fastid
// [example2]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func (c *Configuration) SetFastId(id int) *FastId {
  return &FastId{Id: id, Conf: *c, Err: nil}
}

// Getting the configuration.
//
// More information can be found in the [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func (c *Configuration) GetConfiguration() (string, string) {
  return c.Application, c.AccessToken
}

// You need to enter the application and the private key.
//
// To register the application, follow the link from [OAuth].
//
// More information can be found in the [example].
//
// [OAuth]: https://github.com/heycatch/goshikimori#shikimori-documentation
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/getter_setter
func SetConfiguration(appname, token string) *Configuration {
  return &Configuration{Application: appname, AccessToken: token}
}

type Options struct {
  Order, Kind, Status, Season, Rating,
  Type, Target_type, Duration, Mylist,
  Forum, Linked_type string
  Page, Limit, Score, Linked_id, Target_id int
  Censored bool
  Genre_v2 []int
}

type Result interface {
  OptionsAnime()           string
  OptionsManga()           string
  OptionsRanobe()          string
  OptionsUsers()           string
  OptionsClub()            string
  OptionsCalendar()        string
  OptionsAnimeRates()      string
  OptionsMangaRates()      string
  OptionsUserHistory()     string
  OptionsMessages()        string
  OptionsPeople()          string
  OptionsClubAnimeManga()  string
  OptionsClubCollections() string
  OptionsTopics()          string
  OptionsTopicsHot()       string
  OptionsRandomAnime()     string
  OptionsRandomManga()     string
  OptionsRandomRanobe()    string
}

// V2 implementation of OptionsTopics().
//
// BenchmarkTopicsV1-4   203890   6134 ns/op   488 B/op   14 allocs/op
//
// BenchmarkTopicsV2-4   745861   1884 ns/op   280 B/op   10 allocs/op
func (o *Options) OptionsTopics() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 31 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  search.LinearComplexity(&o.Forum, "all", []string{
    "cosplay", "animanga", "site", "games", "vn",
    "contests", "offtopic", "clubs", "my_clubs",
    "critiques", "news", "collections", "articles",
  })
  v.Add("forum", o.Forum)

  search.LinearComplexity(&o.Linked_type, "", []string{
    "Anime", "Manga", "Ranobe", "Character", "Person",
    "Club", "ClubPage", "Critique", "Review",
    "Contest", "CosplayGallery", "Collection", "Article",
  })
  // linked_id and linked_type are only used together.
  if o.Linked_id >= 1 && o.Linked_type != "" {
    v.Add("linked_id", strconv.Itoa(o.Linked_id))
    v.Add("linked_type", o.Linked_type)
  }

  return v.Encode()
}

// V2 implementation of OptionsMessages().
//
// BenchmarkMessagesV1-4    471354   2375 ns/op   312 B/op   10 allocs/op
//
// BenchmarkMessagesV2-4   1000000   1092 ns/op   152 B/op    7 allocs/op
func (o *Options) OptionsMessages() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 101 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  search.LinearComplexity(&o.Type, "news", []string{
    "inbox", "private", "sent", "news", "notifications",
  })
  v.Add("type", o.Type)

  return v.Encode()
}

// V2 implementation of OptionsUserHistory().
//
// BenchmarkUserHistoryV1-4   453085   3211 ns/op   408 B/op   12 allocs/op
//
// BenchmarkUserHistoryV2-4   949827   2529 ns/op   248 B/op    9 allocs/op
func (o *Options) OptionsUserHistory() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 101 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  search.LinearComplexity(&o.Target_type, "Anime", []string{"Anime", "Manga"})
  v.Add("target_type", o.Target_type)

  if o.Target_id > 0 { v.Add("target_id", strconv.Itoa(o.Target_id)) }

  return v.Encode()
}

func (o *Options) OptionsUsers() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 101 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  return v.Encode()
}

// V2 implementation of OptionsAnime().
//
// BenchmarkAnimeV1-4   74109   16041 ns/op   4053 B/op   32 allocs/op
//
// BenchmarkAnimeV2-4   198708   6785 ns/op   1727 B/op   22 allocs/op
func (o *Options) OptionsAnime() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 51 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  search.LinearComplexity(&o.Order, "", []string{
    "id", "ranked", "kind", "popularity",
    "name", "aired_on", "episodes", "status",
  })
  v.Add("order", o.Order)

  search.LinearComplexity(&o.Kind, "", []string{
    "tv", "movie", "ova", "ona", "special", "music",
    "tv_13", "tv_24", "tv_48", "!tv", "!movie", "!ova",
    "!ona", "!special", "!music", "!tv_13", "!tv_24", "!tv_48",
  })
  v.Add("kind", o.Kind)

  search.LinearComplexity(&o.Status, "", []string{
    "anons", "ongoing", "released", "!anons", "!ongoing", "!released",
  })
  v.Add("status", o.Status)

  search.LinearComplexity(&o.Season, "", []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  v.Add("season", o.Season)

  search.LinearComplexity(&o.Rating, "", []string{
    "none", "g", "pg", "pg_13",
    "r", "r_plus", "rx", "!g", "!pg",
    "!pg_13", "!r", "!r_plus", "!rx",
  })
  v.Add("rating", o.Rating)

  search.LinearComplexity(&o.Duration, "", []string{
    "S", "D", "F", "!S", "!D", "!F",
  })
  v.Add("duration", o.Duration)

  search.LinearComplexity(&o.Mylist, "", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresAnime(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

// V2 implementation of OptionsManga().
//
// BenchmarkMangaV1-4    99231   13817 ns/op   3662 B/op   27 allocs/op
//
// BenchmarkMangaV2-4   228412   5485 ns/op    1391 B/op   20 allocs/op
func (o *Options) OptionsManga() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 51 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  search.LinearComplexity(&o.Order, "", []string{
    "id", "ranked", "kind", "popularity", "name",
    "aired_on", "volumes", "chapters", "status",
  })
  v.Add("order", o.Order)

  search.LinearComplexity(&o.Kind, "",[]string{
    "manga", "manhwa", "manhua", "light_novel", "novel",
    "one_shot", "doujin", "!manga", "!manhwa", "!manhua",
    "!light_novel", "!novel", "!one_shot", "!doujin",
  })
  v.Add("kind", o.Kind)

  search.LinearComplexity(&o.Status, "",[]string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  })
  v.Add("status", o.Status)

  search.LinearComplexity(&o.Season, "", []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  v.Add("season", o.Season)

  search.LinearComplexity(&o.Mylist, "", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresManga(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

// V2 implementation of OptionsRanobe().
//
// BenchmarkRanobeV1-4   116786   10772 ns/op   2847 B/op   23 allocs/op
//
// BenchmarkRanobeV2-4   341701    3751 ns/op    616 B/op   17 allocs/op
func (o *Options) OptionsRanobe() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 51 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  search.LinearComplexity(&o.Order, "", []string{
    "id", "ranked", "kind", "popularity", "name",
    "aired_on", "volumes", "chapters", "status",
  })
  v.Add("order", o.Order)

  search.LinearComplexity(&o.Status, "", []string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  })
  v.Add("status", o.Status)

  search.LinearComplexity(&o.Season, "", []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  v.Add("season", o.Season)

  search.LinearComplexity(&o.Mylist, "", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresManga(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

func (o *Options) OptionsClub() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 31 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  return v.Encode()
}

func (o *Options) OptionsCalendar() string {
  v := url.Values{}

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

// V2 implementation of AnimeRates().
//
// BenchmarkRanobe-4   406456   3327 ns/op   408 B/op   12 allocs/op
//
// BenchmarkRanobe-4   958911   2595 ns/op   248 B/op    9 allocs/op
func (o *Options) OptionsAnimeRates() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 5001 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  search.LinearComplexity(
    &o.Status, "watching", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("status", o.Status)

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

func (o *Options) OptionsMangaRates() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 5001 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

// V2 implementation of OptionsPeople().
//
// BenchmarkPeopleV1-4    537405    1929 ns/op   216 B/op   7 allocs/op
//
// BenchmarkPeopleV2-4   1936758   797.7 ns/op    56 B/op   4 allocs/op
func (o *Options) OptionsPeople() string {
  v := url.Values{}

  search.LinearComplexity(
    &o.Kind, "seyu",
    []string{"seyu", "mangaka", "producer"},
  )
  v.Add("kind", o.Kind)

  return v.Encode()
}

func (o *Options) OptionsClubAnimeManga() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 21 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  return v.Encode()
}

func (o *Options) OptionsClubCollections() string {
  v := url.Values{}

  if o.Page <= 1 || o.Page >= 100001 {
    v.Add("page", "1")
  } else {
    v.Add("page", strconv.Itoa(o.Page))
  }
  if o.Limit <= 1 || o.Limit >= 5 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  return v.Encode()
}

func (o *Options) OptionsTopicsHot() string {
  v := url.Values{}

  if o.Limit <= 1 || o.Limit >= 11 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }

  return v.Encode()
}

func (o *Options) OptionsRandomAnime() string {
  v := url.Values{}

  if o.Limit <= 1 || o.Limit >= 51 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  search.LinearComplexity(&o.Kind, "", []string{
    "tv", "movie", "ova", "ona", "special", "music",
    "tv_13", "tv_24", "tv_48", "!tv", "!movie", "!ova",
    "!ona", "!special", "!music", "!tv_13", "!tv_24", "!tv_48",
  })
  v.Add("kind", o.Kind)

  search.LinearComplexity(&o.Status, "", []string{
    "anons", "ongoing", "released", "!anons", "!ongoing", "!released",
  })
  v.Add("status", o.Status)

  search.LinearComplexity(&o.Season, "", []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  v.Add("season", o.Season)

  search.LinearComplexity(&o.Rating, "", []string{
    "none", "g", "pg", "pg_13",
    "r", "r_plus", "rx", "!g", "!pg",
    "!pg_13", "!r", "!r_plus", "!rx",
  })
  v.Add("rating", o.Rating)

  search.LinearComplexity(&o.Duration, "", []string{
    "S", "D", "F", "!S", "!D", "!F",
  })
  v.Add("duration", o.Duration)

  search.LinearComplexity(&o.Mylist, "", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresAnime(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

func (o *Options) OptionsRandomManga() string {
  v := url.Values{}

  if o.Limit <= 1 || o.Limit >= 51 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  search.LinearComplexity(&o.Kind, "",[]string{
    "manga", "manhwa", "manhua", "light_novel", "novel",
    "one_shot", "doujin", "!manga", "!manhwa", "!manhua",
    "!light_novel", "!novel", "!one_shot", "!doujin",
  })
  v.Add("kind", o.Kind)

  search.LinearComplexity(&o.Status, "",[]string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  })
  v.Add("status", o.Status)

  search.LinearComplexity(&o.Season, "", []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  v.Add("season", o.Season)

  search.LinearComplexity(&o.Mylist, "", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresManga(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

func (o *Options) OptionsRandomRanobe() string {
  v := url.Values{}

  if o.Limit <= 1 || o.Limit >= 51 {
    v.Add("limit", "1")
  } else {
    v.Add("limit", strconv.Itoa(o.Limit))
  }
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  search.LinearComplexity(&o.Status, "",[]string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  })
  v.Add("status", o.Status)

  search.LinearComplexity(&o.Season, "", []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  v.Add("season", o.Season)

  search.LinearComplexity(&o.Mylist, "", []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresManga(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}
