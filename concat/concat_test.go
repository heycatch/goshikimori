package concat

import "testing"

// BenchmarkIdsToStringOld-4   1760985   688.7 ns/op   120 B/op   10 allocs/op   1.906s
func BenchmarkIdsToStringOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = idsToStringOld([]int{1336, 1337, 1338, 1339, 1400})
  }
  b.StopTimer()
}

// BenchmarkIdsToString-4   2245165   491.6 ns/op   116 B/op   7 allocs/op   1.648s
func BenchmarkIdsToString(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = IdsToString([]int{1336, 1337, 1338, 1339, 1400})
  }
  b.StopTimer()
}

// BenchmarkUrlOld-4   4291356   313.9 ns/op   80 B/op   3 allocs/op   1.636s
func BenchmarkUrlOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = urlOld("https://shikimori.one/api/", "animes?search=initial+d&page=1&limit=50")
  }
  b.StopTimer()
}

// BenchmarkUrl-4   15332547	  82.20 ns/op   80 B/op	  1 allocs/op   1.345s
func BenchmarkUrl(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = Url("https://shikimori.one/api/", "animes?search=initial+d&page=1&limit=50")
  }
  b.StopTimer()
}

// BenchmarkDataReorderOld-4   4893867   253.0 ns/op   32 B/op   2 allocs/op   1.492s
// BenchmarkDataBuffer-4       6330697   204.8 ns/op   72 B/op   2 allocs/op   1.489s
func BenchmarkDataReorderOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = dataReorderOld(1337)
  }
  b.StopTimer()
}

// BenchmarkDataMarkReadOld-4   3413079   411.1 ns/op   80 B/op   2 allocs/op   1.765s
// BenchmarkDataBuffer-4        7376164   189.7 ns/op   64 B/op   1 allocs/op   1.572s
func BenchmarkDataMarkReadOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = dataMarkReadOld("1336,1337,1338,1339,1340", 0)
  }
  b.StopTimer()
}

// BenchmarkDataReadDeleteOld-4   4201455   262.7 ns/op   64 B/op   2 allocs/op   1.400s
// BenchmarkDataCopy-4            9057472   123.5 ns/op   48 B/op   1 allocs/op   1.259s
func BenchmarkDataReadDeleteOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = dataReadDeleteOld("news")
  }
  b.StopTimer()
}

// BenchmarkDataSendMessageOld-4   1574340   721.4 ns/op   160 B/op   4 allocs/op   1.910s
// BenchmarkDataBuffer-4   	       3007321   524.3 ns/op   200 B/op   4 allocs/op   1.985s
func BenchmarkDataSendMessageOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = dataSendMessageOld("test message", 1337, 1338)
  }
  b.StopTimer()
}

// BenchmarkDataChangeMessageOld-4   4111257   310.7 ns/op   96 B/op   2 allocs/op   1.579s
// BenchmarkDataCopy-4               14183776  88.33 ns/op   80 B/op   1 allocs/op   1.343s
func BenchmarkDataChangeMessageOld(b *testing.B) {
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    _ = dataChangeMessageOld("updated test message")
  }
  b.StopTimer()
}
