package concat

import (
  "bytes"
  "testing"
)

func TestIdsToString(t *testing.T) {
  if IdsToString([]int{1336, 1337, 1338}) == "1336,1337,1338" {
    t.Log("IdsToString passed")
  } else {
    t.Error("IdsToString failed")
  }
}

func TestNekoSliceToString(t *testing.T) {
  if NekoSliceToString([]string{"initial", "d", "first", "stage"}) == "initial_d_first_stage" {
    t.Log("NekoSliceToString passed")
  } else {
    t.Error("NekoSliceToString failed")
  }
}

func TestUrl(t *testing.T) {
  if Url(46, []string{"https://shikimori.one/api/", "users/", "search=arctica"}) ==
      "https://shikimori.one/api/users/search=arctica" {
    t.Log("Url apssed")
  } else {
    t.Error("Url failed")
  }
}

func TestBearer(t *testing.T) {
  if Bearer("XXX_TOKEN_XXX") == "Bearer XXX_TOKEN_XXX" {
    t.Log("Bearer passed")
  } else {
    t.Error("Bearer failed")
  }
}

func TestDataBuffer(t *testing.T) {
  if bytes.Equal(DataBuffer([]string{"zero", "one", "1337"}), []byte("zeroone1337")) {
    t.Log("DataBuffer passed")
  } else {
    t.Error("DataBuffer failed")
  }
}

func TestDataCopy(t *testing.T) {
  if bytes.Equal(DataCopy(11, []string{"zero", "one", "1337"}), []byte("zeroone1337")) {
    t.Log("DataCopy passed")
  } else {
    t.Error("DataCopy failed")
  }
}

// BenchmarkIdsToStringOld-4   1760985   688.7 ns/op   120 B/op   10 allocs/op   1.906s
// BenchmarkIdsToString-4      2245165   491.6 ns/op   116 B/op   7 allocs/op    1.648s
func BenchmarkIdsToString(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = IdsToString([]int{1336, 1337, 1338, 1339, 1400})
  }
  b.StopTimer()
}

// BenchmarkUrlOld-4   4291356   313.9 ns/op   80 B/op   3 allocs/op   1.636s
// BenchmarkUrl-4     15332547	 82.20 ns/op   48 B/op	 1 allocs/op   1.345s
func BenchmarkUrl(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = Url(46, []string{"https://shikimori.one/api/", "users/", "search=arctica"})
  }
  b.StopTimer()
}

// BenchmarkGenresV1-4   1338148    1003 ns/op   200 B/op   5 allocs/op
// BenchmarkGenresV2-4   2062106   583.6 ns/op   128 B/op   2 allocs/op
func BenchmarkGenres(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = MapGenresAnime([]int{2, 14, 10, 31, 12})
    //_ = MapGenresManga([]int{49, 59, 51, 82, 73})
  }
  b.StopTimer()
}
