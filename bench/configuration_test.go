package bench

import (
	"testing"

	g "github.com/heycatch/goshikimori"
	"github.com/heycatch/goshikimori/consts"
)

func makeMinOptions() *g.Options {
	return &g.Options{
		Page: 1,
		Limit: 10,
	}
}

func makeMidOptions() *g.Options {
	return &g.Options{
		Page: 1,
		Limit: 10,
		Score: 7,
		Order: consts.ANIME_ORDER_AIRED_ON,
		Kind: consts.ANIME_KIND_TV,
		Season: consts.SEASON_199x,
		Censored: true,
	}
}

func makeFullOptions() *g.Options {
	return &g.Options{
		Page: 1,
		Limit: 10,
		Score: 7,
		Order: consts.ANIME_ORDER_AIRED_ON,
		Kind: consts.ANIME_KIND_TV,
		Status: consts.ANIME_STATUS_RELEASED,
		Season: consts.SEASON_199x,
		Rating: consts.ANIME_RATING_PG_13,
		Duration: consts.ANIME_DURATION_F,
		Mylist: consts.MY_LIST_PLANNED,
		Censored: true,
		Genre_v2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
	}
}

func BenchmarkOptionsAnimeV1Min(b *testing.B) {
	o := makeMinOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsAnime()
	}
}

func BenchmarkOptionsAnimeV2Min(b *testing.B) {
	o := makeMinOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsAnimeV2()
	}
}

func BenchmarkOptionsAnimeV1Mid(b *testing.B) {
	o := makeMidOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsAnime()
	}
}

func BenchmarkOptionsAnimeV2Mid(b *testing.B) {
	o := makeMidOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsAnimeV2()
	}
}

func BenchmarkOptionsAnimeV1Full(b *testing.B) {
	o := makeFullOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsAnime()
	}
}

func BenchmarkOptionsAnimeV2Full(b *testing.B) {
	o := makeFullOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsAnimeV2()
	}
}

func BenchmarkOptionsOnlyPageLimit(b *testing.B) {
	o := makeMinOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsOnlyPageLimit(100000, 30)
	}
}

func BenchmarkOptionsOnlyPageLimitV2(b *testing.B) {
	o := makeMinOptions()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = o.OptionsOnlyPageLimitV2()
	}
}
