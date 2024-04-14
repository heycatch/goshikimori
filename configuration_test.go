package goshikimori

import "testing"

func TestOptionsTopics(t *testing.T) {
  empty := Options{}
  if empty.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Empty OptionsTopics passed")
  } else {
    t.Error("Empty OptionsTopics failed")
  }

  big := Options{
    Page: 100002, Limit: 32, Forum: "1111111111",
    Linked_id: 222222222, Linked_type: "1111111111111111",
  }
  if big.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  zero := Options{
    Page: 0, Limit: 0, Forum: "0",
    Linked_id: 0, Linked_type: "0",
  }
  if zero.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  negative := Options{
    Page: -1, Limit: -1, Forum: "-1",
    Linked_id: -1, Linked_type: "-1",
  }
  if negative.OptionsTopics() == "forum=all&limit=1&page=1" {
    t.Log("Big OptionsTopics passed")
  } else {
    t.Error("Big OptionsTopics failed")
  }

  normal_one := Options{
    Page: 5, Limit: 10, Forum: "animanga",
    Linked_id: 342908, Linked_type: "Anime",
  }
  if normal_one.OptionsTopics() == "forum=animanga&limit=10&linked_id=342908&linked_type=Anime&page=5" {
    t.Log("Normal-one OptionsTopics passed")
  } else {
    t.Error("Normal-one OptionsTopics failed")
  }

  normal_two := Options{
    Page: 3, Limit: 8, Forum: "animanga",
    Linked_id: 2323, Linked_type: "Manga",
  }
  if normal_two.OptionsTopics() == "forum=animanga&limit=8&linked_id=2323&linked_type=Manga&page=3" {
    t.Log("Normal-two OptionsTopics passed")
  } else {
    t.Error("Normal-two OptionsTopics failed")
  }
}

func TestOptionsMessages(t *testing.T) {
  empty := Options{}
  if empty.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Empty OptionsMessages passed")
  } else {
    t.Error("Empty OptionsMessages failed")
  }

  big := Options{Type: "11111111111111", Page: 100002, Limit: 102}
  if big.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Big OptionsMessages passed")
  } else {
    t.Error("Big OptionsMessages failed")
  }

  zero := Options{Type: "0", Page: 0, Limit: 0}
  if zero.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Zero OptionsMessages passed")
  } else {
    t.Error("Zero OptionsMessages failed")
  }

  negative := Options{Type: "-1", Page: -1, Limit: -1}
  if negative.OptionsMessages() == "limit=1&page=1&type=news" {
    t.Log("Negative OptionsMessages passed")
  } else {
    t.Error("Negative OptionsMessages failed")
  }

  normal := Options{Type: "private", Page: 2, Limit: 10}
  if normal.OptionsMessages() == "limit=10&page=2&type=private" {
    t.Log("Normal OptionsMessages passed")
  } else {
    t.Error("Normal OptionsMessages failed")
  }
}

func TestOptionsUserHistory(t *testing.T) {
  empty := Options{}
  if empty.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Empty OptionsUserHistory passed")
  } else {
    t.Error("Empty OptionsUserHistory failed")
  }

  big := Options{Page: 100002, Limit: 102, Target_id: "", Target_type: "11111111111111"}
  if big.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Big OptionsUserHistory passed")
  } else {
    t.Error("Big OptionsUserHistory failed")
  }

  zero := Options{Page: 0, Limit: 0, Target_id: "", Target_type: "0"}
  if zero.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Zero OptionsUserHistory passed")
  } else {
    t.Error("Zero OptionsUserHistory failed")
  }

  negative := Options{Page: -1, Limit: -1, Target_id: "", Target_type: "-1"}
  if negative.OptionsUserHistory() == "limit=1&page=1&target_type=Anime" {
    t.Log("Negative OptionsUserHistory passed")
  } else {
    t.Error("Negative OptionsUserHistory failed")
  }

  normal := Options{Page: 3, Limit: 20, Target_id: "1337", Target_type: "Manga"}
  if normal.OptionsUserHistory() == "limit=20&page=3&target_id=1337&target_type=Manga" {
    t.Log("Zero OptionsUserHistory passed")
  } else {
    t.Error("Zero OptionsUserHistory failed")
  }
}

func TestOptionsUsers(t *testing.T) {
  empty := Options{}
  if empty.OptionsUsers() == "limit=1&page=1" {
    t.Log("Empty OptionsUsers passed")
  } else {
    t.Error("Empty OptionsUsers failed")
  }

  big := Options{Page: 100002, Limit: 102}
  if big.OptionsUsers() == "limit=1&page=1" {
    t.Log("Big OptionsUsers passed")
  } else {
    t.Error("Big OptionsUsers failed")
  }

  zero := Options{Page: 0, Limit: 0}
  if zero.OptionsUsers() == "limit=1&page=1" {
    t.Log("Zero OptionsUsers passed")
  } else {
    t.Error("Zero OptionsUsers failed")
  }

  negative := Options{Page: -1, Limit: -1}
  if negative.OptionsUsers() == "limit=1&page=1" {
    t.Log("Negative OptionsUsers passed")
  } else {
    t.Error("Negative OptionsUsers failed")
  }

  normal := Options{Page: 7, Limit: 37}
  if normal.OptionsUsers() == "limit=37&page=7" {
    t.Log("Normal OptionsUsers passed")
  } else {
    t.Error("Normal OptionsUsers failed")
  }
}

func TestOptionsAnime(t *testing.T) {
  empty := Options{}
  if empty.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&season=&status=" {
    t.Log("Empty OptionsAnime passed")
  } else {
    t.Error("Empty OptionsAnime failed")
  }

  big := Options{
    Page: 100002, Limit: 52, Order: "111111111", Kind: "11111111", Status: "111111111",
    Season: "1111111111", Score: 111111111111, Rating: "10",
    Duration: "11111111111111", Censored: false, Mylist: "1111111111111111", Genre_v2: []int{111111111},
  }
  if big.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&season=&status=" {
    t.Log("Big OptionsAnime passed")
  } else {
    t.Error("Big OptionsAnime failed")
  }

  zero := Options{
    Page: 0, Limit: 0, Order: "0", Kind: "0", Status: "0",
    Season: "0", Score: 0, Rating: "0", Duration: "0",
    Censored: false, Mylist: "0", Genre_v2: []int{0},
  }
  if zero.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&season=&status=" {
    t.Log("Zero OptionsAnime passed")
  } else {
    t.Error("Zero OptionsAnime failed")
  }

  negative := Options{
    Page: -1, Limit: -1, Order: "-1", Kind: "-1", Status: "-1",
    Season: "-1", Score: -1, Rating: "-1", Duration: "-1",
    Censored: false, Mylist: "-1", Genre_v2: []int{-1},
  }
  if negative.OptionsAnime() == "censored=false&duration=&kind=&limit=1&mylist=&order=&page=1&rating=&season=&status=" {
    t.Log("Negative OptionsAnime passed")
  } else {
    t.Error("Negative OptionsAnime failed")
  }

  normal := Options{
    Page: 2, Limit: 12, Order: "id", Kind: "tv", Status: "released",
    Season: "199x", Score: 8, Rating: "r", Duration: "D",
    Censored: true, Mylist: "watching", Genre_v2: []int{539},
  }
  if normal.OptionsAnime() == "censored=true&duration=D&genre_v2=539-Erotica&kind=tv&limit=12&mylist=watching&order=id&page=2&rating=r&score=8&season=199x&status=released" {
    t.Log("Normal OptionsAnime passed")
  } else {
    t.Error("Normal OptionsAnime failed")
  }
}

func TestOptionsManga(t *testing.T) {
  empty := Options{}
  if empty.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&season=&status=" {
    t.Log("Empty OptionsManga passed")
  } else {
    t.Error("Empty OptionsManga failed")
  }

  big := Options{
    Page: 100002, Limit: 52, Order: "111111111", Kind: "11111111", Status: "111111111",
    Season: "1111111111", Score: 1111111111,
    Censored: false, Mylist: "11111111111111", Genre_v2: []int{111111111111},
  }
  if big.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&season=&status=" {
    t.Log("Big OptionsManga passed")
  } else {
    t.Error("Big OptionsManga failed")
  }

  zero := Options{
    Page: 0, Limit: 0, Order: "0", Kind: "0", Status: "0",
    Season: "0", Score: 0, Censored: false, Mylist: "0", Genre_v2: []int{0},
  }
  if zero.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&season=&status=" {
    t.Log("Zero OptionsManga passed")
  } else {
    t.Error("Zero OptionsManga failed")
  }

  negative := Options{
    Page: -1, Limit: -1, Order: "-1", Kind: "-1", Status: "-1",
    Season: "-1", Score: -1, Censored: false, Mylist: "-1", Genre_v2: []int{-1},
  }
  if negative.OptionsManga() == "censored=false&kind=&limit=1&mylist=&order=&page=1&season=&status=" {
    t.Log("Negative OptionsManga passed")
  } else {
    t.Error("Negative OptionsManga failed")
  }

  normal := Options{
    Page: 4, Limit: 5, Order: "id", Kind: "manga", Status: "anons",
    Season: "198x", Score: 7, Censored: true, Mylist: "planned", Genre_v2: []int{540},
  }
  if normal.OptionsManga() == "censored=true&genre_v2=540-Erotica&kind=manga&limit=5&mylist=planned&order=id&page=4&score=7&season=198x&status=anons" {
    t.Log("Normal OptionsManga passed")
  } else {
    t.Error("Normal OptionsManga failed")
  }
}

func TestOptionsClub(t *testing.T) {
  empty := Options{}
  if empty.OptionsClub() == "limit=1&page=1" {
    t.Log("Empty OptionsClub passed")
  } else {
    t.Error("Empty OptionsClub failed")
  }

  big := Options{Page: 100002, Limit: 32}
  if big.OptionsClub() == "limit=1&page=1" {
    t.Log("Big OptionsClub passed")
  } else {
    t.Error("Big OptionsClub failed")
  }

  zero := Options{Page: 0, Limit: 0}
  if zero.OptionsClub() == "limit=1&page=1" {
    t.Log("Zero OptionsClub passed")
  } else {
    t.Error("Zero OptionsClub failed")
  }

  negative := Options{Page: -1, Limit: -1}
  if negative.OptionsClub() == "limit=1&page=1" {
    t.Log("Negative OptionsClub passed")
  } else {
    t.Error("Negative OptionsClub failed")
  }

  normal := Options{Page: 2, Limit: 22}
  if normal.OptionsClub() == "limit=22&page=2" {
    t.Log("Normal OptionsClub passed")
  } else {
    t.Error("Normal OptionsClub failed")
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
  if empty.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Empty OptionsAnimeRates passed")
  } else {
    t.Error("Empty OptionsAnimeRates failed")
  }

  big := Options{Page: 100002, Limit: 5002, Status: "11111111", Censored: false}
  if big.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Big OptionsAnimeRates passed")
  } else {
    t.Error("Big OptionsAnimeRates failed")
  }

  zero := Options{Page: 0, Limit: 0, Status: "0", Censored: false}
  if zero.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Zero OptionsAnimeRates passed")
  } else {
    t.Error("Zero OptionsAnimeRates failed")
  }

  negative := Options{Page: -1, Limit: -1, Status: "-1", Censored: false}
  if negative.OptionsAnimeRates() == "censored=false&limit=1&page=1&status=watching" {
    t.Log("Negative OptionsAnimeRates passed")
  } else {
    t.Error("Negative OptionsAnimeRates failed")
  }

  normal := Options{Page: 15, Limit: 405, Status: "dropped", Censored: true}
  if normal.OptionsAnimeRates() == "censored=true&limit=405&page=15&status=dropped" {
    t.Log("Normal OptionsAnimeRates passed")
  } else {
    t.Error("Normal OptionsAnimeRates failed")
  }
}

func TestOptionsMangaRates(t *testing.T) {
  empty := Options{}
  if empty.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Empty OptionsMangaRates passed")
  } else {
    t.Error("Empty OptionsMangaRates failed")
  }

  big := Options{Page: 100002, Limit: 5002, Censored: false}
  if big.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Big OptionsMangaRates passed")
  } else {
    t.Error("Big OptionsMangaRates failed")
  }

  zero := Options{Page: 0, Limit: 0, Censored: false}
  if zero.OptionsMangaRates() == "censored=false&limit=1&page=1" {
    t.Log("Zero OptionsMangaRates passed")
  } else {
    t.Error("Zero OptionsMangaRates failed")
  }

  negative := Options{Page: -1, Limit: -1, Censored: false}
  if negative.OptionsMangaRates() == "censored=false&limit=1&page=1" {
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
  if empty.OptionsPeople() == "kind=seyu" {
    t.Log("Empty OptionsPeople passed")
  } else {
    t.Error("Empty OptionsPeople failed")
  }

  big := Options{Kind: "111111111"}
  if big.OptionsPeople() == "kind=seyu" {
    t.Log("Big OptionsPeople passed")
  } else {
    t.Error("Big OptionsPeople failed")
  }

  zero := Options{Kind: "0"}
  if zero.OptionsPeople() == "kind=seyu" {
    t.Log("Zero OptionsPeople passed")
  } else {
    t.Error("Zero OptionsPeople failed")
  }

  negative := Options{Kind: "-1"}
  if negative.OptionsPeople() == "kind=seyu" {
    t.Log("Negative OptionsPeople passed")
  } else {
    t.Error("Negative OptionsPeople failed")
  }

  normal := Options{Kind: "mangaka"}
  if normal.OptionsPeople() == "kind=mangaka" {
    t.Log("Normal OptionsPeople passed")
  } else {
    t.Error("Normal OptionsPeople failed")
  }
}

func TestOptionsClubInformation(t *testing.T) {
  empty := Options{}
  if empty.OptionsClubAnimeManga() == "limit=1&page=1" {
    t.Log("Empty OptionsClubInformation passed")
  } else {
    t.Error("Empty OptionsClubInformation failed")
  }

  big := Options{Page: 100002, Limit: 22}
  if big.OptionsClubAnimeManga() == "limit=1&page=1" {
    t.Log("Big OptionsClubInformation passed")
  } else {
    t.Error("Big OptionsClubInformation failed")
  }

  zero := Options{Page: 0, Limit: 0}
  if zero.OptionsClubAnimeManga() == "limit=1&page=1" {
    t.Log("Zero OptionsClubInformation passed")
  } else {
    t.Error("Zero OptionsClubInformation failed")
  }

  negative := Options{Page: -1, Limit: -1}
  if negative.OptionsClubAnimeManga() == "limit=1&page=1" {
    t.Log("Negative OptionsClubInformation passed")
  } else {
    t.Error("Negative OptionsClubInformation failed")
  }

  normal := Options{Page: 1337, Limit: 5}
  if normal.OptionsClubAnimeManga() == "limit=5&page=1337" {
    t.Log("Normal OptionsClubInformation passed")
  } else {
    t.Error("Normal OptionsClubInformation failed")
  }
}
