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

func TestMapGenresAnime(t *testing.T) {
  if MapGenresAnime([]int{2, 14, 10, 88, 31, 12, 2, 539, 10, 31, 29}) ==
      "2-Adventure,14-Horror,10-Fantasy,31-Super Power,12-Hentai,539-Erotica,29-Space" {
    t.Log("MapGenresAnime passed")
  } else {
    t.Error("MapGenresAnime failed")
  }
}

func TestMapGenresManga(t *testing.T) {
  if MapGenresManga([]int{49, 58, 66, 45, 49, 540, 78, 78, 85, 88, 63}) ==
      "49-Comedy,58-Magic,66-Martial Arts,540-Erotica,78-Music,85-Space,88-Samurai,63-Shoujo" {
    t.Log("MapGenresManga passed")
  } else {
    t.Error("MapGenresManga failed")
  }
}

func BenchmarkIdsToString(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = IdsToString([]int{1336, 1337, 1338, 1339, 1400})
  }
  b.StopTimer()
}

func BenchmarkUrl(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = Url(46, []string{"https://shikimori.one/api/", "users/", "search=arctica"})
  }
  b.StopTimer()
}

func BenchmarkGenres(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = MapGenresAnime([]int{2, 14, 10, 10, 12})
    //_ = MapGenresManga([]int{49, 59, 51, 51, 73})
  }
  b.StopTimer()
}
