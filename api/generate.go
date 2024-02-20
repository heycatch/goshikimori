package api

import "strconv"

// Auxiliary function to get the correct list of genres.
//
// name: Anime or Manga;
//
// genres: []api.Genres;
func GenerateGenres(name string, genres []Genres) map[int]string {
  data := make(map[int]string)
  for _, v := range genres {
    if v.Entry_type == name {
      data[v.Id] = strconv.Itoa(v.Id) + "-" + v.Name
    }
  }
  return data
}
