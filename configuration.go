package goshikimori

import (
  "net/url"
  "strconv"
  "sync"
  "strings"

  "github.com/heycatch/goshikimori/search"
)

type Configuration struct {
  Application, AccessToken string
}

type FastId struct {
  Id   int
  Conf Configuration
  Err  error
}

// You need to enter the application and the private key.
//
// To register the application, follow the link from [OAuth].
//
// [OAuth]: https://github.com/heycatch/goshikimori#shikimori-documentation
func Add(appname, token string) *Configuration {
  return &Configuration{Application: appname, AccessToken: token}
}

type Options struct {
  Order, Kind, Status, Season, Rating, Type, Target_id, Target_type, Duration, Mylist, Forum, Linked_type string
  Page, Limit, Score, Linked_id int
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
}

func (o *Options) OptionsTopics() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 31 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))

  var wg sync.WaitGroup
  wg.Add(2)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Forum, []string{
    "cosplay", "animanga", "site", "games", "vn",
    "contests", "offtopic", "clubs", "my_clubs",
    "critiques", "news", "collections", "articles",
  })
  o.Forum = <-ch

  go search.StringWithChan(&wg, ch, o.Linked_type, []string{
    "Anime", "Manga", "Ranobe", "Character", "Person",
    "Club", "ClubPage", "Critique", "Review",
    "Contest", "CosplayGallery", "Collection", "Article",
  })
  o.Linked_type = <-ch

  wg.Wait()

  if o.Forum == "" { v.Add("forum", "all") } else { v.Add("forum", o.Forum) }
  // linked_id and linked_type are only used together.
  if o.Linked_id >= 1 && o.Linked_type != "" {
    v.Add("linked_id", strconv.Itoa(o.Linked_id))
    v.Add("linked_type", o.Linked_type)
  }
  return v.Encode()
}

func (o *Options) OptionsMessages() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 101 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))

  var wg sync.WaitGroup
  wg.Add(1)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Type, []string{
    "inbox", "private", "sent", "news", "notifications",
  })
  o.Type = <-ch

  wg.Wait()

  if o.Type == "" { v.Add("type", "news") } else { v.Add("type", o.Type) }
  return v.Encode()
}

func (o *Options) OptionsUserHistory() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 101 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))

  var wg sync.WaitGroup
  wg.Add(1)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Target_type, []string{"Anime", "Manga"})
  o.Target_type = <-ch

  wg.Wait()

  // We get an error if we do not process the request in this way.
  // json: cannot unmarshal string into Go value of type api.UserHistory
  if o.Target_id != "" { v.Add("target_id", o.Target_id) }
  if o.Target_type == "" { v.Add("target_type", "Anime") } else { v.Add("target_type", o.Target_type) }
  return v.Encode()
}

func (o *Options) OptionsUsers() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 101 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  return v.Encode()
}

func (o *Options) OptionsAnime() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 51 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  var wg sync.WaitGroup
  wg.Add(7)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Order, []string{
    "id", "ranked", "kind", "popularity",
    "name", "aired_on", "episodes", "status",
  })
  o.Order = <-ch

  go search.StringWithChan(&wg, ch, o.Kind, []string{
    "tv", "movie", "ova", "ona", "special", "music",
    "tv_13", "tv_24", "tv_48", "!tv", "!movie", "!ova",
    "!ona", "!special", "!music", "!tv_13", "!tv_24", "!tv_48",
  })
  o.Kind = <-ch

  go search.StringWithChan(&wg, ch, o.Status, []string{
    "anons", "ongoing", "released", "!anons", "!ongoing", "!released",
  })
  o.Status = <-ch

  go search.StringWithChan(&wg, ch, o.Season, []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  o.Season = <-ch

  go search.StringWithChan(&wg, ch, o.Rating, []string{
    "none", "g", "pg", "pg_13",
    "r", "r_plus", "rx", "!g", "!pg",
    "!pg_13", "!r", "!r_plus", "!rx",
  })
  o.Rating = <-ch

  go search.StringWithChan(&wg, ch, o.Duration, []string{
    "S", "D", "F", "!S", "!D", "!F",
  })
  o.Duration = <-ch

  go search.StringWithChan(&wg, ch, o.Mylist, []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  o.Mylist = <-ch

  wg.Wait()

  var genre_v2 string
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
  for i := 0; i < len(o.Genre_v2); i++ {
    _, ok := genres[o.Genre_v2[i]]; if ok {
      genre_v2 += genres[o.Genre_v2[i]] + ","
    }
  }

  v.Add("order", o.Order)
  v.Add("kind", o.Kind)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("rating", o.Rating)
  v.Add("duration", o.Duration)
  v.Add("mylist", o.Mylist)
  v.Add("censored", strconv.FormatBool(o.Censored))
  if len(genre_v2) > 0 { v.Add("genre_v2", strings.TrimSuffix(genre_v2, ",")) }
  return v.Encode()
}

func (o *Options) OptionsManga() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 51 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  var wg sync.WaitGroup
  wg.Add(5)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Order, []string{
    "id", "ranked", "kind", "popularity", "name",
    "aired_on", "volumes", "chapters", "status",
  })
  o.Order = <-ch

  go search.StringWithChan(&wg, ch, o.Kind, []string{
    "manga", "manhwa", "manhua", "light_novel", "novel",
    "one_shot", "doujin", "!manga", "!manhwa", "!manhua",
    "!light_novel", "!novel", "!one_shot", "!doujin",
  })
  o.Kind = <-ch

  go search.StringWithChan(&wg, ch, o.Status, []string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  })
  o.Status = <-ch

  go search.StringWithChan(&wg, ch, o.Season, []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  o.Season = <-ch

  go search.StringWithChan(&wg, ch, o.Mylist, []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  o.Mylist = <-ch

  wg.Wait()

  var genre_v2 string
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
  for i := 0; i < len(o.Genre_v2); i++ {
    _, ok := genres[o.Genre_v2[i]]; if ok {
      genre_v2 += genres[o.Genre_v2[i]] + ","
    }
  }

  v.Add("order", o.Order)
  v.Add("kind", o.Kind)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("mylist", o.Mylist)
  v.Add("censored", strconv.FormatBool(o.Censored))
  if len(genre_v2) > 0 { v.Add("genre_v2", strings.TrimSuffix(genre_v2, ",")) }
  return v.Encode()
}

func (o *Options) OptionsRanobe() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 51 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  if o.Score >= 1 && o.Score <= 9 { v.Add("score", strconv.Itoa(o.Score)) }

  var wg sync.WaitGroup
  wg.Add(4)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Order, []string{
    "id", "ranked", "kind", "popularity", "name",
    "aired_on", "volumes", "chapters", "status",
  })
  o.Order = <-ch

  go search.StringWithChan(&wg, ch, o.Status, []string{
    "anons", "ongoing", "released", "paused", "discontinued",
    "!anons", "!ongoing", "!released", "!paused", "!discontinued",
  })
  o.Status = <-ch

  go search.StringWithChan(&wg, ch, o.Season, []string{
    "2000_2010", "2010_2014", "2015_2019", "199x",
    "!2000_2010", "!2010_2014", "!2015_2019", "!199x",
    "198x", "!198x", "2020_2021", "!2020_2021",
    "2022", "!2022", "2023", "!2023",
  })
  o.Season = <-ch

  go search.StringWithChan(&wg, ch, o.Mylist, []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  o.Mylist = <-ch

  wg.Wait()

  var genre_v2 string
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
  for i := 0; i < len(o.Genre_v2); i++ {
    _, ok := genres[o.Genre_v2[i]]; if ok {
      genre_v2 += genres[o.Genre_v2[i]] + ","
    }
  }

  v.Add("order", o.Order)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("mylist", o.Mylist)
  v.Add("censored", strconv.FormatBool(o.Censored))
  if len(genre_v2) > 0 { v.Add("genre_v2", strings.TrimSuffix(genre_v2, ",")) }
  return v.Encode()
}

func (o *Options) OptionsClub() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 31 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  return v.Encode()
}

func (o *Options) OptionsCalendar() string {
  v := url.Values{}

  v.Add("censored", strconv.FormatBool(o.Censored))
  return v.Encode()
}

func (o *Options) OptionsAnimeRates() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 5001 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))

  var wg sync.WaitGroup
  wg.Add(1)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Status, []string{
    "planned", "watching", "rewatching",
    "completed", "on_hold", "dropped",
  })
  o.Status = <-ch

  wg.Wait()

  if o.Status == "" { v.Add("status", "watching") } else { v.Add("status", o.Status) }
  v.Add("censored", strconv.FormatBool(o.Censored))
  return v.Encode()
}

func (o *Options) OptionsMangaRates() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 5001 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  v.Add("censored", strconv.FormatBool(o.Censored))
  return v.Encode()
}

func (o *Options) OptionsPeople() string {
  v := url.Values{}

  var wg sync.WaitGroup
  wg.Add(1)

  ch := make(chan string)

  go search.StringWithChan(&wg, ch, o.Kind, []string{"seyu", "mangaka", "producer"})
  o.Kind = <-ch

  wg.Wait()

  if o.Kind == "" { v.Add("kind", "seyu") } else { v.Add("kind", o.Kind) }
  return v.Encode()
}

func (o *Options) OptionsClubAnimeManga() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 21 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  return v.Encode()
}

func (o *Options) OptionsClubCollections() string {
  v := url.Values{}

  if o.Page <= 0 || o.Page >= 100001 { o.Page = 1 }
  v.Add("page", strconv.Itoa(o.Page))
  if o.Limit <= 0 || o.Limit >= 5 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  return v.Encode()
}

func (o *Options) OptionsTopicsHot() string {
  v := url.Values{}

  if o.Limit <= 0 || o.Limit >= 11 { o.Limit = 1 }
  v.Add("limit", strconv.Itoa(o.Limit))
  return v.Encode()
}
