package graphql

import "testing"

// Constants are not available in tests, so they are implemented manually.

func TestAnimeSchema(t *testing.T) {
  pass_normal := `graphql?query={animes(search: "initial d", page: 1, limit: 1, score: 8, order: id, kind: "tv", status: "!anons", season: "199x", duration: "F", rating: "!rx", mylist: "completed", censored: false, genre: "3-Cars,4-Comedy"){id name russian english japanese score airedOn{year month day date} }}`
  normal, _ := AnimeSchema(
    Values("id", "name", "russian", "english", "japanese", "score", "airedOn{year month day date}"),
    "initial d",
    1, 1, 8, "id", "tv", "!anons", "199x", "F", "!rx", "completed", false, []int{3, 4, 4, 3},
  )
  if normal == pass_normal {
    t.Log("Normal AnimeSchema passed")
  } else {
    t.Error("Normal AnimeSchema failed")
  }

  pass_empty := `graphql?query={animes(search: "initial d", page: 1, limit: 1, score: 1, censored: false){id}}`
  empty, _ := AnimeSchema(
    Values(""),
    "initial d",
    1, 1, 1, "", "", "", "", "", "", "", false, nil,
  )
  if empty == pass_empty {
    t.Log("Empty AnimeSchema passed")
  } else {
    t.Error("Empty AnimeSchema failed")
  }
}

func TestMangaSchema(t *testing.T) {
  pass_normal := `graphql?query={mangas(search: "angel", page: 1, limit: 3, score: 8, order: ranked, kind: "manga", status: "released", mylist: "planned", censored: false, genre: "50-Drama,64-Vampire"){id name russian volumes chapters releasedOn{year month day date} url }}`
  normal, _ := MangaSchema(
    Values("id", "name", "russian", "volumes", "chapters", "releasedOn{year month day date}", "url"),
    "angel",
    1, 3, 8, "ranked", "manga", "released", "", "planned", false, []int{50, 64, 64, 50},
  )
  if normal == pass_normal {
    t.Log("Normal MangaSchema passed")
  } else {
    t.Error("Normal MangaSchema failed")
  }

  pass_empty := `graphql?query={mangas(search: "initial d", page: 1, limit: 1, score: 1, censored: false){id}}`
  empty, _ := MangaSchema(
    Values(""),
    "initial d",
    1, 1, 1, "", "", "", "", "", false, nil,
  )
  if empty == pass_empty {
    t.Log("Empty MangaSchema passed")
  } else {
    t.Error("Empty MangaSchema failed")
  }
}

func TestCharacterSchema(t *testing.T) {
  pass := `graphql?query={characters(search: "onizuka", page: 1, limit: 1){id name russian poster{originalUrl} description }}`
  normal, _ := CharacterSchema(
    Values("id", "name", "russian", "poster{originalUrl}", "description"),
    "onizuka",
    1, 1,
  )
  if normal == pass {
    t.Log("Normal CharacterSchema passed")
  } else {
    t.Error("Normal CharacterSchema failed")
  }
}

func TestPeopleSchema(t *testing.T) {
  pass := `graphql?query={people(search: "satsuki", page: 1, limit: 1, isSeyu: true, isMangaka: false, isProducer: false){id name russian url website birthOn{year month day date} }}`
  normal, _ := PeopleSchema(
    Values("id", "name", "russian", "url", "website", "birthOn{year month day date}"),
    "satsuki",
    1, 1, true, false, false,
  )
  if normal == pass {
    t.Log("Normal PeopleSchema passed")
  } else {
    t.Error("Normal PeopleSchema failed")
  }
}
