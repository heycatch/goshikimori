package goshikimori

import (
  "net/url"
  "strconv"
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
  Page, Limit, Order, Kind, Status, Season, Score string
  Rating, Censored, Type, Target_id, Target_type string
  Duration, Mylist, Forum, Linked_id, Linked_type string
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
  OptionsClubInformation() string
  OptionsTopics()          string
  OptionsTopicsHot()       string
}

func (o *Options) OptionsTopics() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 31 { o.Limit = "1" }

  forum_map := map[string]int8{
    "cosplay": 1, "animanga": 2, "site": 3,
    "games": 4, "vn": 5, "contests": 6,
    "offtopic": 7, "clubs": 8, "my_clubs": 9,
    "critiques": 10, "news": 11,
    "collections": 12, "articles": 13,
  }
  _, ok_forum := forum_map[o.Forum]; if !ok_forum {
    o.Forum = "all"
  }

  li, _ := strconv.Atoi(o.Linked_id)
  if li <= 0 { o.Linked_id = "" }

  linked_type_map := map[string]int8{
    "Anime": 1, "Manga": 2, "Ranobe": 3,
    "Character": 4, "Person": 5, "Club": 6,
    "ClubPage": 7, "Critique": 8, "Review": 9,
    "Contest": 10, "CosplayGallery": 11,
    "Collection": 12, "Article": 13,
  }
  _, ok_linked_type := linked_type_map[o.Linked_type]; if !ok_linked_type {
    o.Linked_type = ""
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("forum", o.Forum)
  // linked_id and linked_type are only used together.
  if o.Linked_id != "" && o.Linked_type != "" {
    v.Add("linked_id", o.Linked_id)
    v.Add("linked_type", o.Linked_type)
  }

  return v.Encode()
}

func (o *Options) OptionsMessages() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 101 { o.Limit = "1" }

  type_map := map[string]int8{
    "inbox": 1, "private": 2, "sent": 3,
    "news": 4, "notifications": 5,
  }
  _, ok_type := type_map[o.Type]; if !ok_type {
    o.Type = "news"
  }

  v := url.Values{}
  v.Add("type", o.Type)
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)

  return v.Encode()
}

func (o *Options) OptionsUserHistory() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 101 { o.Limit = "1" }

  target_map := map[string]int8{"Anime": 1, "Manga": 2}
  _, ok := target_map[o.Target_type]; if !ok {
    o.Target_type = "Anime"
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  // We get an error if we do not process the request in this way.
  // json: cannot unmarshal string into Go value of type api.UserHistory
  if o.Target_id != "" { v.Add("target_id", o.Target_id) }
  v.Add("target_type", o.Target_type)

  return v.Encode()
}

func (o *Options) OptionsUsers() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 101 { o.Limit = "1" }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)

  return v.Encode()
}

func (o *Options) OptionsAnime() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 51 { o.Limit = "1" }

  order_map := map[string]int8{
    "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
    "name": 5, "aired_on": 6, "episodes": 7, "status": 8,
  }
  _, ok_order := order_map[o.Order]; if !ok_order {
    o.Order = ""
  }

  kind_map := map[string]int8{
    "tv": 1, "movie": 2, "ova": 3, "ona": 4, "special": 5, "music": 6,
    "tv_13": 7, "tv_24": 8, "tv_48": 9, "!tv": 10, "!movie": 11,
    "!ova": 12, "!ona": 13, "!special": 14, "!music": 15, "!tv_13": 16,
    "!tv_24": 17, "!tv_48": 18,
  }
  _, ok_kind := kind_map[o.Kind]; if !ok_kind {
    o.Kind = ""
  }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3,
    "!anons": 4, "!ongoing": 5, "!released": 6,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = ""
  }

  season_map := map[string]int8{
    "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
    "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
    "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
    "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
  }
  _, ok_season := season_map[o.Season]; if !ok_season {
    o.Season = ""
  }

  s, _ := strconv.Atoi(o.Score)
  if s <= 0 || s >= 10 { o.Score = "" }

  rating_map := map[string]int8{
    "none": 1, "g": 2, "pg": 3, "pg_13": 4,
    "r": 5, "r_plus": 6, "rx": 7, "!g": 8, "!pg": 9,
    "!pg_13": 10, "!r": 11, "!r_plus": 12, "!rx": 13,
  }
  _, ok_rating := rating_map[o.Rating]; if !ok_rating {
    o.Rating = ""
  }

  duration_map := map[string]int8{
    "S": 1, "D": 2, "F": 3, "!S": 4, "!D": 5, "!F": 6,
  }
  _, ok_duration := duration_map[o.Duration]; if !ok_duration {
    o.Duration = ""
  }

  mylist_map := map[string]int8{
    "planned": 1, "watching": 2, "rewatching": 3,
    "completed": 4, "on_hold": 5, "dropped": 6,
  }
  _, ok_mylist := mylist_map[o.Mylist]; if !ok_mylist {
    o.Mylist = ""
  }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok_censored := censored_map[o.Censored]; if !ok_censored {
    o.Censored = "false"
  }

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

  if len(genre_v2) != 0 { genre_v2 = genre_v2[:len(genre_v2)-1] }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("order", o.Order)
  v.Add("kind", o.Kind)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("score", o.Score)
  v.Add("rating", o.Rating)
  v.Add("duration", o.Duration)
  v.Add("censored", o.Censored)
  v.Add("mylist", o.Mylist)
  v.Add("genre_v2", genre_v2)

  return v.Encode()
}

func (o *Options) OptionsManga() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 51 { o.Limit = "1" }

  order_map := map[string]int8{
    "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
    "name": 5, "aired_on": 6, "volumes": 7,
    "chapters": 8, "status": 9,
  }
  _, ok_order := order_map[o.Order]; if !ok_order {
    o.Order = ""
  }

  kind_map := map[string]int8{
    "manga": 1, "manhwa": 2, "manhua": 3, "light_novel": 4, "novel": 5,
    "one_shot": 6, "doujin": 7, "!manga": 8, "!manhwa": 9, "!manhua": 10,
    "!light_novel": 11, "!novel": 12, "!one_shot": 13, "!doujin": 14,
  }
  _, ok_kind := kind_map[o.Kind]; if !ok_kind {
    o.Kind = ""
  }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3, "paused": 4, "discontinued": 5,
    "!anons": 6, "!ongoing": 7, "!released": 8, "!paused": 9, "!discontinued": 10,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = ""
  }

  season_map := map[string]int8{
    "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
    "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
    "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
    "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
  }
  _, ok_season := season_map[o.Season]; if !ok_season {
    o.Season = ""
  }

  s, _ := strconv.Atoi(o.Score)
  if s <= 0 || s >= 10 { o.Score = "" }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok_censored := censored_map[o.Censored]; if !ok_censored {
    o.Censored = "false"
  }

  mylist_map := map[string]int8{
    "planned": 1, "watching": 2, "rewatching": 3,
    "completed": 4, "on_hold": 5, "dropped": 6,
  }
  _, ok_mylist := mylist_map[o.Mylist]; if !ok_mylist {
    o.Mylist = ""
  }

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

  if len(genre_v2) != 0 { genre_v2 = genre_v2[:len(genre_v2)-1] }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("order", o.Order)
  v.Add("kind", o.Kind)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("score", o.Score)
  v.Add("censored", o.Censored)
  v.Add("mylist", o.Mylist)
  v.Add("genre_v2", genre_v2)

  return v.Encode()
}

func (o *Options) OptionsRanobe() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 51 { o.Limit = "1" }

  order_map := map[string]int8{
    "id": 1, "ranked": 2, "kind": 3, "popularity": 4,
    "name": 5, "aired_on": 6, "volumes": 7,
    "chapters": 8, "status": 9,
  }
  _, ok_order := order_map[o.Order]; if !ok_order {
    o.Order = ""
  }

  status_map := map[string]int8{
    "anons": 1, "ongoing": 2, "released": 3, "paused": 4, "discontinued": 5,
    "!anons": 6, "!ongoing": 7, "!released": 8, "!paused": 9, "!discontinued": 10,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = ""
  }

  season_map := map[string]int8{
    "2000_2010": 1, "2010_2014": 2, "2015_2019": 3, "199x": 4,
    "!2000_2010": 5, "!2010_2014": 6, "!2015_2019": 7, "!199x": 8,
    "198x": 9, "!198x": 10, "2020_2021": 11, "!2020_2021": 12,
    "2022": 13, "!2022": 14, "2023": 15, "!2023": 16,
  }
  _, ok_season := season_map[o.Season]; if !ok_season {
    o.Season = ""
  }

  s, _ := strconv.Atoi(o.Score)
  if s <= 0 || s >= 10 { o.Score = "" }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok_censored := censored_map[o.Censored]; if !ok_censored {
    o.Censored = "false"
  }

  mylist_map := map[string]int8{
    "planned": 1, "watching": 2, "rewatching": 3,
    "completed": 4, "on_hold": 5, "dropped": 6,
  }
  _, ok_mylist := mylist_map[o.Mylist]; if !ok_mylist {
    o.Mylist = ""
  }

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

  if len(genre_v2) != 0 { genre_v2 = genre_v2[:len(genre_v2)-1] }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("order", o.Order)
  v.Add("status", o.Status)
  v.Add("season", o.Season)
  v.Add("score", o.Score)
  v.Add("censored", o.Censored)
  v.Add("mylist", o.Mylist)
  v.Add("genre_v2", genre_v2)

  return v.Encode()
}

func (o *Options) OptionsClub() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 31 { o.Limit = "1" }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)

  return v.Encode()
}

func (o *Options) OptionsCalendar() string {
  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok := censored_map[o.Censored]; if !ok {
    o.Censored = "false"
  }

  v := url.Values{}
  v.Add("censored", o.Censored)

  return v.Encode()
}

func (o *Options) OptionsAnimeRates() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 5001 { o.Limit = "1" }

  status_map := map[string]int8{
    "planned": 1, "watching": 2,
    "rewatching": 3, "completed": 4,
    "on_hold": 5, "dropped": 6,
  }
  _, ok_status := status_map[o.Status]; if !ok_status {
    o.Status = "watching"
  }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok_censored := censored_map[o.Censored]; if !ok_censored {
    o.Censored = "false"
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("status", o.Status)
  v.Add("censored", o.Censored)

  return v.Encode()
}

func (o *Options) OptionsMangaRates() string {
  p, _ := strconv.Atoi(o.Page)
  l, _ := strconv.Atoi(o.Limit)

  if p <= 0 || p >= 100001 { o.Page = "1" }
  if l <= 0 || l >= 5001 { o.Limit = "1" }

  censored_map := map[string]int8{"true": 1, "false": 2}
  _, ok := censored_map[o.Censored]; if !ok {
    o.Censored = "false"
  }

  v := url.Values{}
  v.Add("page", o.Page)
  v.Add("limit", o.Limit)
  v.Add("censored", o.Censored)

  return v.Encode()
}

func (o *Options) OptionsPeople() string {
  kind_map := map[string]int8{
    "seyu": 1, "mangaka": 2, "producer": 3,
  }
  _, ok := kind_map[o.Kind]; if !ok {
    o.Kind = "seyu"
  }

  v := url.Values{}
  v.Add("kind", o.Kind)

  return v.Encode()
}

func (o *Options) OptionsClubInformation() string {
  p, _ := strconv.Atoi(o.Page)

  if p <= 0 || p >= 100001 { o.Page = "1" }

  v := url.Values{}
  v.Add("page", o.Page)

  return v.Encode()
}

func (o *Options) OptionsTopicsHot() string {
  l, _ := strconv.Atoi(o.Limit)

  if l <= 0 || l >= 11 { o.Limit = "1" }

  v := url.Values{}
  v.Add("limit", o.Limit)

  return v.Encode()
}
