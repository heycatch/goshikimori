package graphql

import "testing"

func TestAnimeSchema(t *testing.T) {
  pass_normal := `graphql?query={animes(search: "initial d", limit: 1, score: 8, order: id, kind: "tv", status: "!anons", season: "199x", duration: "F", rating: "!rx", mylist: "completed", censored: false){id name russian english japanese score airedOn{year month day date} }}`
  normal, _ := AnimeSchema(
    Values("id", "name", "russian", "english", "japanese", "score", "airedOn{year month day date}"),
    "initial d", 1, 8, "id", "tv", "!anons", "199x", "F", "!rx", "completed", false,
  )
  if normal == pass_normal {
    t.Log("Normal AnimeSchema passed")
  } else {
    t.Error("Normal AnimeSchema failed")
  }

  pass_empty := `graphql?query={animes(search: "initial d", limit: 1, score: 1, censored: false){id}}`
  empty, _ := AnimeSchema(
    Values(""),
    "initial d", 1, 1, "", "", "", "", "", "", "", false,
  )
  if empty == pass_empty {
    t.Log("Empty AnimeSchema passed")
  } else {
    t.Error("Empty AnimeSchema failed")
  }
}

func TestMangaSchema(t *testing.T) {
  pass_normal := `graphql?query={mangas(search: "initial d", limit: 3, score: 8, order: ranked, kind: "manga", status: "released", mylist: "planned", censored: false){id name russian volumes chapters releasedOn{year month day date} url }}`
  normal, _ := MangaSchema(
    Values("id", "name", "russian", "volumes", "chapters", "releasedOn{year month day date}", "url"),
    "initial d", 3, 8, "ranked", "manga", "released", "", "planned", false,
  )
  if normal == pass_normal {
    t.Log("Normal MangaSchema passed")
  } else {
    t.Error("Normal MangaSchema faile")
  }

  pass_empty := `graphql?query={mangas(search: "initial d", limit: 1, score: 1, censored: false){id}}`
  empty, _ := MangaSchema(
    Values(""),
    "initial d", 1, 1, "", "", "", "", "", false,
  )
  if empty == pass_empty {
    t.Log("Empty MangaSchema passed")
  } else {
    t.Error("Empty MangaSchema failed")
  }
}
