package api

import (
  "strconv"
  "strings"
)

// Auxiliary function to get the correct list of genres.
//
// name: Anime or Manga;
//
// genres: []api.Genres;
func GenerateGenres(name string, genres []Genres) map[int]string {
  data := make(map[int]string)
  for _, v := range genres {
    if v.Entry_type == name {
      data[v.Id] = toString(strconv.Itoa(v.Id), v.Name)
    }
  }
  return data
}

func toString(id, name string) string {
  var res strings.Builder
  res.Grow(len(id) + 1 + len(name))
  res.WriteString(id)
  res.WriteString("-")
  res.WriteString(name)
  return res.String()
}
