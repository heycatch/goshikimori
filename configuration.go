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

var (
  topic_forum = []string{
    "cosplay", "animanga", "site", "games", "vn",
    "contests", "offtopic", "clubs", "my_clubs",
    "critiques", "news", "collections", "articles",
  }

  topic_linkedType = []string{
    "Anime", "Manga", "Ranobe", "Character", "Person",
    "Club", "ClubPage", "Critique", "Review",
    "Contest", "CosplayGallery", "Collection", "Article",
  }

  message_type = []string{
    "inbox", "private", "sent", "news", "notifications",
  }

  userHistory_targetType = []string{"Anime", "Manga"}

  anime_order = []string {
    "id", "ranked", "kind", "popularity",
    "name", "aired_on", "episodes", "status",
  }

  anime_kind = []string{
    "tv", "movie", "ova", "ona", "special",
    "music", "tv_13", "tv_24",
    "tv_48", "!tv", "!movie", "!ova",
    "!ona", "!special", "!music",
    "!tv_13", "!tv_24", "!tv_48",
  }

  anime_status = []string{
    "anons", "ongoing", "released",
    "!anons", "!ongoing", "!released",
  }

  anime_season = []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  }

  anime_rating = []string{
    "none", "g", "pg", "pg_13",
    "r", "r_plus", "rx", "!g", "!pg",
    "!pg_13", "!r", "!r_plus", "!rx",
  }

  anime_duration = []string{
    "S", "D", "F", "!S", "!D", "!F",
  }

  anime_mylist = []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  }


  manga_order = []string{
    "id", "ranked", "kind", "popularity", "name",
    "aired_on", "volumes", "chapters", "status",
  }

  manga_kind = []string{
    "manga", "manhwa", "manhua", "light_novel", "novel",
    "one_shot", "doujin", "!manga", "!manhwa", "!manhua",
    "!light_novel", "!novel", "!one_shot", "!doujin",
  }

  manga_status = []string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  }

  people_kind = &[]string{"seyu", "mangaka", "producer"}
)

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

  search.LinearComplexity(&o.Forum, "all", topic_forum)
  v.Add("forum", o.Forum)
  search.LinearComplexity(&o.Linked_type, "", topic_linkedType)
  // linked_id and linked_type are only used together.
  if o.Linked_id >= 1 && o.Linked_type != "" {
    v.Add("linked_id", strconv.Itoa(o.Linked_id))
    v.Add("linked_type", o.Linked_type)
  }

  return v.Encode()
}

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

  search.LinearComplexity(&o.Type, "news", message_type)
  v.Add("type", o.Type)

  return v.Encode()
}

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

  search.LinearComplexity(&o.Target_type, "Anime", userHistory_targetType)
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

  search.LinearComplexity(&o.Order, "", anime_order)
  v.Add("order", o.Order)
  search.LinearComplexity(&o.Kind, "", anime_kind)
  v.Add("kind", o.Kind)
  search.LinearComplexity(&o.Status, "", anime_status)
  v.Add("status", o.Status)
  search.LinearComplexity(&o.Season, "", anime_season)
  v.Add("season", o.Season)
  search.LinearComplexity(&o.Rating, "", anime_rating)
  v.Add("rating", o.Rating)
  search.LinearComplexity(&o.Duration, "", anime_duration)
  v.Add("duration", o.Duration)
  search.LinearComplexity(&o.Mylist, "", anime_mylist)
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresAnime(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

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

  search.LinearComplexity(&o.Order, "", manga_order)
  v.Add("order", o.Order)
  search.LinearComplexity(&o.Kind, "", manga_kind)
  v.Add("kind", o.Kind)
  search.LinearComplexity(&o.Status, "", manga_status)
  v.Add("status", o.Status)
  search.LinearComplexity(&o.Season, "", anime_season)
  v.Add("season", o.Season)
  search.LinearComplexity(&o.Mylist, "", anime_mylist)
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresManga(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}

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

  search.LinearComplexity(&o.Order, "", manga_order)
  v.Add("order", o.Order)
  search.LinearComplexity(&o.Status, "", manga_status)
  v.Add("status", o.Status)
  search.LinearComplexity(&o.Season, "", anime_season)
  v.Add("season", o.Season)
  search.LinearComplexity(&o.Mylist, "", anime_mylist)
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
    &o.Status, "watching", anime_mylist,
  )
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

func (o *Options) OptionsPeople() string {
  v := url.Values{}

  search.LinearComplexity(
    &o.Kind, "seyu", *people_kind,
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

  search.LinearComplexity(&o.Kind, "", anime_kind)
  v.Add("kind", o.Kind)
  search.LinearComplexity(&o.Status, "", anime_status)
  v.Add("status", o.Status)
  search.LinearComplexity(&o.Season, "", anime_season)
  v.Add("season", o.Season)
  search.LinearComplexity(&o.Rating, "", anime_rating)
  v.Add("rating", o.Rating)
  search.LinearComplexity(&o.Duration, "", anime_duration)
  v.Add("duration", o.Duration)
  search.LinearComplexity(&o.Mylist, "", anime_mylist)
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

  search.LinearComplexity(&o.Kind, "", manga_kind)
  v.Add("kind", o.Kind)
  search.LinearComplexity(&o.Status, "", manga_status)
  v.Add("status", o.Status)
  search.LinearComplexity(&o.Season, "", anime_season)
  v.Add("season", o.Season)
  search.LinearComplexity(&o.Mylist, "", anime_mylist)
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

  search.LinearComplexity(&o.Status, "", manga_status)
  v.Add("status", o.Status)
  search.LinearComplexity(&o.Season, "", anime_status)
  v.Add("season", o.Season)
  search.LinearComplexity(&o.Mylist, "", anime_mylist)
  v.Add("mylist", o.Mylist)

  genre := concat.MapGenresManga(o.Genre_v2)
  if genre != "" { v.Add("genre_v2", genre) }

  v.Add("censored", strconv.FormatBool(o.Censored))

  return v.Encode()
}
