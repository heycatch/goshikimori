package graphql

import "testing"

/*
BenchmarkSchemaAnimeV1-4   447230   4336 ns/op   1008 B/op   19 allocs/op
BenchmarkSchemaAnimeV2-4   662517   3582 ns/op   1200 B/op   10 allocs/op

BenchmarkSchemaMangaV1-4   424665   3645 ns/op    960 B/op   19 allocs/op
BenchmarkSchemaMangaV2-4   678516   4145 ns/op   1200 B/op   10 allocs/op

BenchmarkSchemaCharacterV1-4   1000000    1675 ns/op   288 B/op   11 allocs/op
BenchmarkSchemaCharacterV2-4   1991398   718.3 ns/op   272 B/op    4 allocs/op

BenchmarkSchemaPeopleV1-4    868294   2282 ns/op   608 B/op   15 allocs/op
BenchmarkSchemaPeopleV2-4   1000000   1251 ns/op   656 B/op    7 allocs/op
*/
func BenchmarkSchema(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    /*
    _, _ = AnimeSchema(
      Values("id", "name", "score", "episodes", "airedOn{year month day date}"),
      "initial d",
      1, 5, 8, "", "tv", "released", "", "", "pg_13", "", false, []int{3, 4, 4, 6, 44},
    )
    _, _ = MangaSchema(
      Values("id", "name", "score", "volumes", "chapters", "releasedOn{year}"),
      "initial d",
      1, 5, 8, "", "manga", "released", "", "", false, []int{49, 56, 56, 88, 72},
    )
    _, _ = CharacterSchema(
      Values("id", "name", "russian", "url", "description"),
      "onizuka",
      1, 5,
    )
    _, _ = PeopleSchema(
      Values("id", "name", "russian", "url", "website", "birthOn{year month day date}"),
      "satsuki",
      1, 1, true, false, false,
    )
    */
  }
  b.StopTimer()
}

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
