package concat

import (
  "bytes"
  "strconv"
  "strings"

  "github.com/heycatch/goshikimori/api"
)

// Write key to slice and check for duplicates.
func checkForDuplicates(target int, slice []int) bool {
  for i := 0; i < 50; i++ {
    if slice[i] == target { return false }
  }
  return true
}

// Anime value map search.
func MapGenresAnime(slice []int) string {
  var res bytes.Buffer
  var count int
  tempSlice := make([]int, 50)

  for i := 0; i < len(slice); i++ {
    _, ok := api.GenreAnime[slice[i]]
    if ok && checkForDuplicates(slice[i], tempSlice) {
      res.WriteString(api.GenreAnime[slice[i]])
      res.WriteString(",")
      tempSlice[count] = slice[i]
      count++
    }
  }

  return strings.TrimSuffix(res.String(), ",")
}

// Manga value map search.
func MapGenresManga(slice []int) string {
  var res bytes.Buffer
  var count int
  tempSlice := make([]int, 50)

  for i := 0; i < len(slice); i++ {
    _, ok := api.GenreManga[slice[i]]
    if ok && checkForDuplicates(slice[i], tempSlice) {
      res.WriteString(api.GenreManga[slice[i]])
      res.WriteString(",")
      tempSlice[count] = slice[i]
      count++
    }
  }

  return strings.TrimSuffix(res.String(), ",")
}

// Convert a slice with an ids into a string.
func IdsToString(slice []int) string {
  var res bytes.Buffer
  for i := range slice {
    res.WriteString(strconv.Itoa(slice[i]))
    res.WriteString(",")
  }
  return strings.TrimSuffix(res.String(), ",")
}

// Convert a slice with an words into a string.
func NekoSliceToString(slice []string) string {
  var res bytes.Buffer
  for i := range slice {
    res.WriteString(slice[i])
    res.WriteString("_")
  }
  return strings.TrimSuffix(res.String(), "_")
}

// Quick creation of a url.
func Url(max_len int, slice []string) string {
  var offset int
  res := make([]byte, max_len)
  for i := range slice {
    offset += copy(res[offset:], slice[i])
  }
  return string(res[:offset])
}

// Quick creation of a bearer token.
func Bearer(token string) string {
  res := make([]byte, 7 + len(token))
  copy(res, "Bearer ")
  copy(res[7:], token)
  return string(res)
}

// Converting a slice to a []byte using a bytes.Buffer.
func DataBuffer(slice []string) []byte {
  var res bytes.Buffer
  for i := range slice { res.WriteString(slice[i]) }
  return res.Bytes()
}

// Converting a slice to a []byte using a copy.
func DataCopy(max_len int, slice []string) []byte {
  var offset int
  res := make([]byte, max_len)
  for i := range slice {
    offset += copy(res[offset:], []byte(slice[i]))
  }
  return res
}

// Auxiliary function to get the correct list of genres.
//
// name:
//
// > GENERATE_GENRES_ANIME, GENERATE_GENRES_MANGA;
//
// genres: []api.Genres;
func GenerateGenres(name string, genres []api.Genres) map[int]string {
  data := make(map[int]string)
  for _, v := range genres {
    if v.Entry_type == name {
      data[v.Id] = string(DataBuffer(
        []string{strconv.Itoa(v.Id), "-", name},
      ))
    }
  }
  return data
}
