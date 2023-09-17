package graphql

import "testing"

func TestAnimeSchema(t *testing.T) {
  pass_normal := `graphql?query={animes(search: "initial d", limit: 1, score: 8, order: id, kind: "tv", status: "!anons", season: "199x", duration: "F", rating: "!rx", mylist: "completed", censored: false){id malId name russian licenseNameRu english japanese synonyms kind rating score status episodes episodesAired duration airedOn{year month day date} releasedOn{year month day date} url poster{id originalUrl mainUrl} fansubbers fandubbers licensors createdAt updatedAt isCensored genres{id name russian kind} studios{id name imageUrl} personRoles{id rolesRu rolesEn person{id name poster{id}}} characterRoles{id rolesRu rolesEn character{id name poster{id}}} related{id anime{id name} manga{id name} relationRu relationEn} videos{id url name kind} screenshots{id originalUrl x166Url x332Url} scoresStats{score count} statusesStats{status count} description descriptionHtml descriptionSource}}`
  normal, _ := AnimeSchema("initial d", 1, 8, "id", "tv", "!anons", "199x", "F", "!rx", "completed", false)
  if normal == pass_normal {
    t.Log("Normal AnimeSchema passed")
  } else {
    t.Error("Normal AnimeSchema failed")
  }

  pass_empty := `graphql?query={animes(search: "initial d", limit: 1, score: 1, censored: false){id malId name russian licenseNameRu english japanese synonyms kind rating score status episodes episodesAired duration airedOn{year month day date} releasedOn{year month day date} url poster{id originalUrl mainUrl} fansubbers fandubbers licensors createdAt updatedAt isCensored genres{id name russian kind} studios{id name imageUrl} personRoles{id rolesRu rolesEn person{id name poster{id}}} characterRoles{id rolesRu rolesEn character{id name poster{id}}} related{id anime{id name} manga{id name} relationRu relationEn} videos{id url name kind} screenshots{id originalUrl x166Url x332Url} scoresStats{score count} statusesStats{status count} description descriptionHtml descriptionSource}}`
  empty, _ := AnimeSchema("initial d", 1, 1, "", "", "", "", "", "", "", false)
  if empty == pass_empty {
    t.Log("Empty AnimeSchema passed")
  } else {
    t.Error("Empty AnimeSchema failed")
  }
}

func TestMangaSchema(t *testing.T) {
  pass_normal := `graphql?query={mangas(search: "initial d", limit: 3, score: 8, order: ranked, kind: "manga", status: "released", mylist: "planned", censored: false){id malId name russian licenseNameRu english japanese synonyms kind score status volumes chapters airedOn{year month day date} releasedOn{year month day date} url poster{id originalUrl mainUrl} licensors createdAt updatedAt isCensored genres{id name russian kind} publishers{id name} personRoles{id rolesRu rolesEn person{id name poster{id}}} characterRoles{id rolesRu rolesEn character{id name poster{id}}} related{id anime{id name} manga{id name} relationRu relationEn} scoresStats{score count} statusesStats{status count} description descriptionHtml descriptionSource}}`
  normal, _ := MangaSchema("initial d", 3, 8, "ranked", "manga", "released", "", "planned", false)
  if normal == pass_normal {
    t.Log("Normal MangaSchema passed")
  } else {
    t.Error("Normal MangaSchema faile")
  }

  pass_empty := `graphql?query={mangas(search: "initial d", limit: 1, score: 1, censored: false){id malId name russian licenseNameRu english japanese synonyms kind score status volumes chapters airedOn{year month day date} releasedOn{year month day date} url poster{id originalUrl mainUrl} licensors createdAt updatedAt isCensored genres{id name russian kind} publishers{id name} personRoles{id rolesRu rolesEn person{id name poster{id}}} characterRoles{id rolesRu rolesEn character{id name poster{id}}} related{id anime{id name} manga{id name} relationRu relationEn} scoresStats{score count} statusesStats{status count} description descriptionHtml descriptionSource}}`
  empty, _ := MangaSchema("initial d", 1, 1, "", "", "", "", "", false)
  if empty == pass_empty {
    t.Log("Empty MangaSchema passed")
  } else {
    t.Error("Empty MangaSchema failed")
  }
}
