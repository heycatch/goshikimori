package concat

import (
  "bytes"
  "unsafe"
  "strconv"
  "strings"

  "github.com/heycatch/goshikimori/api"
)

// TODO: duplicate check.
//
// Anime value map search.
func MapGenresAnime(slice []int) string {
  var res bytes.Buffer
  for i := 0; i < len(slice); i++ {
    _, ok := api.GenreAnime[slice[i]]; if ok {
      res.WriteString(api.GenreAnime[slice[i]])
      res.WriteString(",")
    }
  }
  return strings.TrimSuffix(res.String(), ",")
}

// TODO: duplicate check.
//
// Manga value map search.
func MapGenresManga(slice []int) string {
  var res bytes.Buffer
  for i := 0; i < len(slice); i++ {
    _, ok := api.GenreManga[slice[i]]; if ok {
      res.WriteString(api.GenreManga[slice[i]])
      res.WriteString(",")
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
    offset += copy(res[offset:], []byte(slice[i]))
  }
  return *(*string)(unsafe.Pointer(&res))
}

// Quick creation of a bearer token.
func Bearer(token string) string {
  var offset int
  res := make([]byte, 7 + len(token))
  array := [2]string{"Bearer ", token}
  for i := range array {
    offset += copy(res[offset:], []byte(array[i]))
  }
  return *(*string)(unsafe.Pointer(&res))
}

/* Converting a slice to a []byte using a bytes.Buffer.

BenchmarkDataSendMessageV1-4   1574340   721.4 ns/op   160 B/op   4 allocs/op   1.910s
BenchmarkDataBufferV2-4   	   3007321   524.3 ns/op   200 B/op   4 allocs/op   1.985s

BenchmarkDataReorderV1-4   4893867   253.0 ns/op   32 B/op   2 allocs/op   1.492s
BenchmarkDataBufferV2-4    6330697   204.8 ns/op   72 B/op   2 allocs/op   1.489s

BenchmarkDataMarkReadV1-4   3413079   411.1 ns/op   80 B/op   2 allocs/op   1.765s
BenchmarkDataBufferV2-4     7376164   189.7 ns/op   64 B/op   1 allocs/op   1.572s
*/
func DataBuffer(slice []string) []byte {
  var res bytes.Buffer
  for i := range slice { res.WriteString(slice[i]) }
  return res.Bytes()
}

/* Converting a slice to a []byte using a copy.

BenchmarkDataChangeMessageV1-4   4111257   310.7 ns/op   96 B/op   2 allocs/op   1.579s
BenchmarkDataCopyV2-4            14183776  88.33 ns/op   80 B/op   1 allocs/op   1.343s

BenchmarkDataReadDeleteV1-4   4201455   262.7 ns/op   64 B/op   2 allocs/op   1.400s
BenchmarkDataCopyV2-4         9057472   123.5 ns/op   48 B/op   1 allocs/op   1.259s
*/
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
// Name(!) must be capitalized.
//
// name: Anime or Manga;
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
