package bench

import (
	"testing"

	"github.com/heycatch/goshikimori/concat"
)

func BenchmarkIdsToString(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = concat.IdsToString([]int{1336, 1337, 1338, 1339, 1400})
	}
	b.StopTimer()
}

func BenchmarkUrl(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = concat.Url(46, []string{"https://shikimori.one/api/", "users/", "search=arctica"})
	}
	b.StopTimer()
}

func BenchmarkGenresAnime(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = concat.MapGenresAnime([]int{2, 14, 10, 10, 12})
	}
	b.StopTimer()
}

func BenchmarkGenresManga(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = concat.MapGenresManga([]int{49, 59, 51, 51, 73})
	}
	b.StopTimer()
}
