package concat

import (
  "fmt"
  "bytes"
  "unsafe"
  "strconv"
  "strings"

  "github.com/heycatch/goshikimori/api"
)

// TODO: remove duplicates.
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

// TODO: remove duplicates.
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

/* Below are examples of the old implementation.

   Needed ONLY for benchmarks.
   You dont even have to look.
*/
func idsToStringOld(slice []int) string {
  var res string
  for i := 0; i < len(slice); i++ {
    if slice[i] != 0 { res += strconv.Itoa(slice[i]) + "," }
  }
  return strings.TrimSuffix(res, ",")
}

func urlOld(site, search string) string {
  return fmt.Sprintf("%s%s", site, search)
}

func dataReorderOld(postion int) []byte {
  return []byte(fmt.Sprintf("{\"new_index\": \"%d\"}", postion))
}

func dataMarkReadOld(ids string, is_read int) []byte {
  return []byte(fmt.Sprintf(`{"ids": "%s", "is_read": "%d"}`, ids, is_read))
}

func dataReadDeleteOld(name string) []byte {
  return []byte(fmt.Sprintf(`{"frontend": "false", "type": "%s"}`, name))
}

func dataSendMessageOld(body string, from_id, to_id int) []byte {
  return []byte(fmt.Sprintf(
    `{"frontend": "false", "message": {"body": "%s",
    "from_id": "%d", "kind": "Private", "to_id": "%d"}}`,
    body, from_id, to_id,
  ))
}

func dataChangeMessageOld(body string) []byte {
  return []byte(fmt.Sprintf(`{"frontend": "false", "message": {"body": "%s"}}`, body))
}
