package api

import "strconv"

// Auxiliary function to get the correct list of genres
func GenerateGenres(genres []Genres) map[int]string {
  data := make(map[int]string)
  for _, v := range genres {
    data[v.Id] = strconv.Itoa(v.Id) + "-" + v.Name
  }
  return data
}
