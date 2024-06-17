package goshikimori

import "testing"

func TestOptionsTopics(t *testing.T) {
  empty := Options{}
  if empty.OptionsTopics() == "" {
    t.Log("Empty OptionsTopics passed")
  } else {
    t.Error("Empty OptionsTopics failed")
  }

  big := Options{
    Page: 100002, Limit: 32,
    Linked_id: 222222222, Linked_type: TOPIC_LINKED_TYPE_ANIME,
  }
  if big.OptionsTopics() == "linked_id=222222222&linked_type=Anime" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  zero := Options{
    Page: 0, Limit: 0,
    Linked_id: 0, Linked_type: TOPIC_LINKED_TYPE_ANIME,
  }
  if zero.OptionsTopics() == "" {
    t.Log("Zero OptionsTopics passed")
  } else {
    t.Error("Zero OptionsTopics failed")
  }

  negative := Options{
    Page: -1, Limit: -1,
    Linked_id: -1, Linked_type: TOPIC_LINKED_TYPE_ANIME,
  }
  if negative.OptionsTopics() == "" {
    t.Log("Negative OptionsTopics passed")
  } else {
    t.Error("Negative OptionsTopics failed")
  }

  normal_one := Options{
    Page: 5, Limit: 10, Forum: TOPIC_FORUM_ANIMANGA,
    Linked_id: 342908, Linked_type: TOPIC_LINKED_TYPE_ANIME,
  }
  if normal_one.OptionsTopics() == "forum=animanga&limit=10&linked_id=342908&linked_type=Anime&page=5" {
    t.Log("Normal-one OptionsTopics passed")
  } else {
    t.Error("Normal-one OptionsTopics failed")
  }

  normal_two := Options{
    Page: 3, Limit: 8, Forum: TOPIC_FORUM_CLUBS,
    Linked_id: 2323, Linked_type: TOPIC_LINKED_TYPE_MANGA,
  }
  if normal_two.OptionsTopics() == "forum=clubs&limit=8&linked_id=2323&linked_type=Manga&page=3" {
    t.Log("Normal-two OptionsTopics passed")
  } else {
    t.Error("Normal-two OptionsTopics failed")
  }
}

func TestOptionsMessages(t *testing.T) {
  empty := Options{}
  if empty.OptionsMessages() == "type=news" {
    t.Log("Empty OptionsMessages passed")
  } else {
    t.Error("Empty OptionsMessages failed")
  }

  big := Options{Type: MESSAGE_TYPE_NEWS, Page: 100002, Limit: 102}
  if big.OptionsMessages() == "type=news" {
    t.Log("Big OptionsMessages passed")
  } else {
    t.Error("Big OptionsMessages failed")
  }

  zero := Options{Type: MESSAGE_TYPE_NEWS, Page: 0, Limit: 0}
  if zero.OptionsMessages() == "type=news" {
    t.Log("Zero OptionsMessages passed")
  } else {
    t.Error("Zero OptionsMessages failed")
  }

  negative := Options{Type: MESSAGE_TYPE_NEWS, Page: -1, Limit: -1}
  if negative.OptionsMessages() == "type=news" {
    t.Log("Negative OptionsMessages passed")
  } else {
    t.Error("Negative OptionsMessages failed")
  }

  normal := Options{Type: MESSAGE_TYPE_PRIVATE, Page: 2, Limit: 10}
  if normal.OptionsMessages() == "limit=10&page=2&type=private" {
    t.Log("Normal OptionsMessages passed")
  } else {
    t.Error("Normal OptionsMessages failed")
  }
}

func TestOptionsUserHistory(t *testing.T) {
  empty := Options{}
  if empty.OptionsUserHistory() == "" {
    t.Log("Empty OptionsUserHistory passed")
  } else {
    t.Error("Empty OptionsUserHistory failed")
  }

  big := Options{Page: 100002, Limit: 102}
  if big.OptionsUserHistory() == "" {
    t.Log("Big OptionsUserHistory passed")
  } else {
    t.Error("Big OptionsUserHistory failed")
  }

  zero := Options{Page: 0, Limit: 0, Target_id: 0}
  if zero.OptionsUserHistory() == "" {
    t.Log("Zero OptionsUserHistory passed")
  } else {
    t.Error("Zero OptionsUserHistory failed")
  }

  negative := Options{Page: -1, Limit: -1, Target_id: -1}
  if negative.OptionsUserHistory() == "" {
    t.Log("Negative OptionsUserHistory passed")
  } else {
    t.Error("Negative OptionsUserHistory failed")
  }

  normal := Options{Page: 3, Limit: 20, Target_id: 1337, Target_type: TARGET_TYPE_MANGA}
  if normal.OptionsUserHistory() == "limit=20&page=3&target_id=1337&target_type=Manga" {
    t.Log("Zero OptionsUserHistory passed")
  } else {
    t.Error("Zero OptionsUserHistory failed")
  }
}

func TestOptionsOnlyPageLimit(t *testing.T) {
  empty := Options{}
  if empty.OptionsOnlyPageLimit(100000, 100) == "" {
    t.Log("Empty OptionsOnlyPageLimit passed")
  } else {
    t.Error("Empty OptionsOnlyPageLimit failed")
  }

  big := Options{Page: 100002, Limit: 102}
  if big.OptionsOnlyPageLimit(100000, 100) == "" {
    t.Log("Big OptionsOnlyPageLimit passed")
  } else {
    t.Error("Big OptionsOnlyPageLimit failed")
  }

  zero := Options{Page: 0, Limit: 0}
  if zero.OptionsOnlyPageLimit(100000, 100) == "" {
    t.Log("Zero OptionsOnlyPageLimit passed")
  } else {
    t.Error("Zero OptionsOnlyPageLimit failed")
  }

  negative := Options{Page: -1, Limit: -1}
  if negative.OptionsOnlyPageLimit(100000, 100) == "" {
    t.Log("Negative OptionsOnlyPageLimit passed")
  } else {
    t.Error("Negative OptionsOnlyPageLimit failed")
  }

  normal := Options{Page: 7, Limit: 37}
  if normal.OptionsOnlyPageLimit(100000, 100) == "limit=37&page=7" {
    t.Log("Normal OptionsOnlyPageLimit passed")
  } else {
    t.Error("Normal OptionsOnlyPageLimit failed")
  }
}

func TestOptionsAnime(t *testing.T) {
  empty := Options{}
  if empty.OptionsAnime() == "censored=false" {
    t.Log("Empty OptionsAnime passed")
  } else {
    t.Error("Empty OptionsAnime failed")
  }

  big := Options{
    Page: 100002, Limit: 52, Score: 12, Censored: false, Genre_v2: []int{111111111},
  }
  if big.OptionsAnime() == "censored=false" {
    t.Log("Big OptionsAnime passed")
  } else {
    t.Error("Big OptionsAnime failed")
  }

  zero := Options{
    Page: 0, Limit: 0, Score: 0,
    Censored: false, Genre_v2: []int{0},
  }
  if zero.OptionsAnime() == "censored=false" {
    t.Log("Zero OptionsAnime passed")
  } else {
    t.Error("Zero OptionsAnime failed")
  }

  negative := Options{
    Page: -1, Limit: -1, Score: -1,
    Censored: false, Genre_v2: []int{-1},
  }
  if negative.OptionsAnime() == "censored=false" {
    t.Log("Negative OptionsAnime passed")
  } else {
    t.Error("Negative OptionsAnime failed")
  }

  normal := Options{
    Page: 2, Limit: 12, Order: ANIME_ORDER_ID, Kind: ANIME_KIND_TV,
    Status: ANIME_STATUS_RELEASED, Season: SEASON_199x,
    Score: 8, Rating: ANIME_RATING_R, Duration: ANIME_DURATION_D,
    Censored: true, Mylist: MY_LIST_WATCHING, Genre_v2: []int{539, 539},
  }
  if normal.OptionsAnime() == "censored=true&duration=D&genre_v2=539-Erotica&kind=tv&limit=12&mylist=watching&order=id&page=2&rating=r&score=8&season=199x&status=released" {
    t.Log("Normal OptionsAnime passed")
  } else {
    t.Error("Normal OptionsAnime failed")
  }
}

func TestOptionsManga(t *testing.T) {
  empty := Options{}
  if empty.OptionsManga() == "censored=false" {
    t.Log("Empty OptionsManga passed")
  } else {
    t.Error("Empty OptionsManga failed")
  }

  big := Options{
    Page: 100002, Limit: 52, Score: 1111111111,
    Censored: false, Genre_v2: []int{111111111111},
  }
  if big.OptionsManga() == "censored=false" {
    t.Log("Big OptionsManga passed")
  } else {
    t.Error("Big OptionsManga failed")
  }

  zero := Options{
    Page: 0, Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
  }
  if zero.OptionsManga() == "censored=false" {
    t.Log("Zero OptionsManga passed")
  } else {
    t.Error("Zero OptionsManga failed")
  }

  negative := Options{
    Page: -1, Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
  }
  if negative.OptionsManga() == "censored=false" {
    t.Log("Negative OptionsManga passed")
  } else {
    t.Error("Negative OptionsManga failed")
  }

  normal := Options{
    Page: 4, Limit: 5, Order: MANGA_ORDER_ID, Kind: MANGA_KIND_MANGA,
    Status: MANGA_STATUS_ANONS, Season: SEASON_198x,
    Score: 7, Censored: true, Mylist: MY_LIST_PLANNED, Genre_v2: []int{540, 540},
  }
  if normal.OptionsManga() == "censored=true&genre_v2=540-Erotica&kind=manga&limit=5&mylist=planned&order=id&page=4&score=7&season=198x&status=anons" {
    t.Log("Normal OptionsManga passed")
  } else {
    t.Error("Normal OptionsManga failed")
  }
}

func TestOptionsRanobe(t *testing.T) {
  empty := Options{}
  if empty.OptionsRanobe() == "censored=false" {
    t.Log("Empty OptionsRanobe passed")
  } else {
    t.Error("Empty OptionsRanobe failed")
  }

  big := Options{
    Page: 100002, Limit: 52, Score: 1111111111,
    Censored: false, Genre_v2: []int{111111111111},
  }
  if big.OptionsRanobe() == "censored=false" {
    t.Log("Big OptionsRanobe passed")
  } else {
    t.Error("Big OptionsRanobe failed")
  }

  zero := Options{
    Page: 0, Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
  }
  if zero.OptionsRanobe() == "censored=false" {
    t.Log("Zero OptionsRanobe passed")
  } else {
    t.Error("Zero OptionsRanobe failed")
  }

  negative := Options{
    Page: -1, Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
  }
  if negative.OptionsRanobe() == "censored=false" {
    t.Log("Negative OptionsRanobe passed")
  } else {
    t.Error("Negative OptionsRanobe failed")
  }

  normal := Options{
    Page: 4, Limit: 5, Order: MANGA_ORDER_ID, Status: MANGA_STATUS_ANONS,
    Season: SEASON_198x, Score: 7, Censored: true, Mylist: MY_LIST_PLANNED,
    Genre_v2: []int{540, 540},
  }
  if normal.OptionsRanobe() == "censored=true&genre_v2=540-Erotica&limit=5&mylist=planned&order=id&page=4&score=7&season=198x&status=anons" {
    t.Log("Normal OptionsRanobe passed")
  } else {
    t.Error("Normal OptionsRanobe failed")
  }
}

func TestOptionsCalendar(t *testing.T) {
  empty := Options{}
  if empty.OptionsCalendar() == "censored=false" {
    t.Log("Empty OptionsCalendar passed")
  } else {
    t.Error("Empty OptionsCalendar failed")
  }

  normal := Options{Censored: true}
  if normal.OptionsCalendar() == "censored=true" {
    t.Log("Normal OptionsCalendar passed")
  } else {
    t.Error("Normal OptionsCalendar failed")
  }
}

func TestOptionsAnimeRates(t *testing.T) {
  empty := Options{}
  if empty.OptionsAnimeRates() == "censored=false" {
    t.Log("Empty OptionsAnimeRates passed")
  } else {
    t.Error("Empty OptionsAnimeRates failed")
  }

  big := Options{Page: 100002, Limit: 5002}
  if big.OptionsAnimeRates() == "censored=false" {
    t.Log("Big OptionsAnimeRates passed")
  } else {
    t.Error("Big OptionsAnimeRates failed")
  }

  zero := Options{Page: 0, Limit: 0}
  if zero.OptionsAnimeRates() == "censored=false" {
    t.Log("Zero OptionsAnimeRates passed")
  } else {
    t.Error("Zero OptionsAnimeRates failed")
  }

  negative := Options{Page: -1, Limit: -1}
  if negative.OptionsAnimeRates() == "censored=false" {
    t.Log("Negative OptionsAnimeRates passed")
  } else {
    t.Error("Negative OptionsAnimeRates failed")
  }

  normal := Options{Page: 15, Limit: 405, Status: MY_LIST_DROPPED, Censored: true}
  if normal.OptionsAnimeRates() == "censored=true&limit=405&page=15&status=dropped" {
    t.Log("Normal OptionsAnimeRates passed")
  } else {
    t.Error("Normal OptionsAnimeRates failed")
  }
}

func TestOptionsMangaRates(t *testing.T) {
  empty := Options{}
  if empty.OptionsMangaRates() == "censored=false" {
    t.Log("Empty OptionsMangaRates passed")
  } else {
    t.Error("Empty OptionsMangaRates failed")
  }

  big := Options{Page: 100002, Limit: 5002}
  if big.OptionsMangaRates() == "censored=false" {
    t.Log("Big OptionsMangaRates passed")
  } else {
    t.Error("Big OptionsMangaRates failed")
  }

  zero := Options{Page: 0, Limit: 0}
  if zero.OptionsMangaRates() == "censored=false" {
    t.Log("Zero OptionsMangaRates passed")
  } else {
    t.Error("Zero OptionsMangaRates failed")
  }

  negative := Options{Page: -1, Limit: -1}
  if negative.OptionsMangaRates() == "censored=false" {
    t.Log("Negative OptionsMangaRates passed")
  } else {
    t.Error("Negative OptionsMangaRates failed")
  }

  normal := Options{Page: 33, Limit: 25, Censored: true}
  if normal.OptionsMangaRates() == "censored=true&limit=25&page=33" {
    t.Log("Normal OptionsMangaRates passed")
  } else {
    t.Error("Normal OptionsMangaRates failed")
  }
}

func TestOptionsPeople(t *testing.T) {
  empty := Options{}
  if empty.OptionsPeople() == "" {
    t.Log("Empty OptionsPeople passed")
  } else {
    t.Error("Empty OptionsPeople failed")
  }

  normal := Options{Kind: PEOPLE_KIND_MANGAKA}
  if normal.OptionsPeople() == "kind=mangaka" {
    t.Log("Normal OptionsPeople passed")
  } else {
    t.Error("Normal OptionsPeople failed")
  }
}

func TestOptionsRandomAnime(t *testing.T) {
  empty := Options{}
  if empty.OptionsRandomAnime() == "censored=false" {
    t.Log("Empty OptionsRandomAnime passed")
  } else {
    t.Error("Empty OptionsRandomAnime failed")
  }

  big := Options{
    Limit: 52, Score: 111111111111,
    Censored: false, Genre_v2: []int{111111111},
  }
  if big.OptionsRandomAnime() == "censored=false" {
    t.Log("Big OptionsRandomAnime passed")
  } else {
    t.Error("Big OptionsRandomAnime failed")
  }

  zero := Options{
    Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
  }
  if zero.OptionsRandomAnime() == "censored=false" {
    t.Log("Zero OptionsRandomAnime passed")
  } else {
    t.Error("Zero OptionsRandomAnime failed")
  }

  negative := Options{
    Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
  }
  if negative.OptionsRandomAnime() == "censored=false" {
    t.Log("Negative OptionsRandomAnime passed")
  } else {
    t.Error("Negative OptionsRandomAnime failed")
  }

  normal := Options{
    Limit: 12, Kind: ANIME_KIND_TV, Status: ANIME_STATUS_RELEASED,
    Mylist: MY_LIST_ON_HOLD, Season: SEASON_199x,
    Score: 8, Rating: ANIME_RATING_R, Duration: ANIME_DURATION_D,
    Censored: true, Genre_v2: []int{539, 539},
  }
  if normal.OptionsRandomAnime() == "censored=true&duration=D&genre_v2=539-Erotica&kind=tv&limit=12&mylist=on_hold&rating=r&score=8&season=199x&status=released" {
    t.Log("Normal OptionsRandomAnime passed")
  } else {
    t.Error("Normal OptionsRandomAnime failed")
  }
}

func TestOptionsRandomManga(t *testing.T) {
  empty := Options{}
  if empty.OptionsRandomManga() == "censored=false" {
    t.Log("Empty OptionsRandomManga passed")
  } else {
    t.Error("Empty OptionsRandomManga failed")
  }

  big := Options{
    Limit: 52, Score: 1111111111,
    Censored: false, Genre_v2: []int{111111111111},
  }
  if big.OptionsRandomManga() == "censored=false" {
    t.Log("Big OptionsRandomManga passed")
  } else {
    t.Error("Big OptionsRandomManga failed")
  }

  zero := Options{
    Limit: 0,Score: 0, Censored: false, Genre_v2: []int{0},
  }
  if zero.OptionsRandomManga() == "censored=false" {
    t.Log("Zero OptionsRandomManga passed")
  } else {
    t.Error("Zero OptionsRandomManga failed")
  }

  negative := Options{
    Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
  }
  if negative.OptionsRandomManga() == "censored=false" {
    t.Log("Negative OptionsRandomManga passed")
  } else {
    t.Error("Negative OptionsRandomManga failed")
  }

  normal := Options{
    Limit: 5, Kind: MANGA_KIND_MANGA, Status: MANGA_STATUS_ANONS,
    Season: SEASON_198x, Score: 7, Censored: true,
    Mylist: MY_LIST_PLANNED, Genre_v2: []int{540, 540},
  }
  if normal.OptionsRandomManga() == "censored=true&genre_v2=540-Erotica&kind=manga&limit=5&mylist=planned&score=7&season=198x&status=anons" {
    t.Log("Normal OptionsRandomManga passed")
  } else {
    t.Error("Normal OptionsRandomManga failed")
  }
}

func TestOptionsRandomRanobe(t *testing.T) {
  empty := Options{}
  if empty.OptionsRandomRanobe() == "censored=false" {
    t.Log("Empty OptionsRandomRanobe passed")
  } else {
    t.Error("Empty OptionsRandomRanobe failed")
  }

  big := Options{
    Limit: 52, Score: 1111111111,
    Censored: false, Genre_v2: []int{111111111111},
  }
  if big.OptionsRandomRanobe() == "censored=false" {
    t.Log("Big OptionsRandomRanobe passed")
  } else {
    t.Error("Big OptionsRandomRanobe failed")
  }

  zero := Options{
    Limit: 0,  Score: 0, Censored: false, Genre_v2: []int{0},
  }
  if zero.OptionsRandomRanobe() == "censored=false" {
    t.Log("Zero OptionsRandomRanobe passed")
  } else {
    t.Error("Zero OptionsRandomRanobe failed")
  }

  negative := Options{
    Limit: -1,  Score: -1, Censored: false, Genre_v2: []int{-1},
  }
  if negative.OptionsRandomRanobe() == "censored=false" {
    t.Log("Negative OptionsRandomRanobe passed")
  } else {
    t.Error("Negative OptionsRandomRanobe failed")
  }

  normal := Options{
    Limit: 5, Status: MANGA_STATUS_ANONS,
    Season: SEASON_198x, Score: 7, Censored: true,
    Mylist: MY_LIST_PLANNED, Genre_v2: []int{540, 540},
  }
  if normal.OptionsRandomRanobe() == "censored=true&genre_v2=540-Erotica&limit=5&mylist=planned&score=7&season=198x&status=anons" {
    t.Log("Normal OptionsRandomRanobe passed")
  } else {
    t.Error("Normal OptionsRandomRanobe failed")
  }
}
