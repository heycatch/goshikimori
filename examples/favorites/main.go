package main

import (
	"fmt"

	g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()
	fast_anime, status, err := c.FastIdAnime("Naruto")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	// add/remove favorites anime.
	fa, status, err := fast_anime.FavoritesCreate(g.FAVORITES_LINKED_TYPE_ANIME, "")
	//fa, status, err := fast_anime.FavoritesDelete("Anime")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fa.Success, fa.Notice)

	// add/remove favorites manga.
	fast_manga, status, err := c.FastIdManga("Naruto")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fm, status, err := fast_manga.FavoritesCreate(g.FAVORITES_LINKED_TYPE_MANGA, "")
	//fm, status, err := fast_manga.FavoritesDelete("Manga")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fm.Success, fm.Notice)

	// add/remove favorites ranobe.
	fast_ranobe, status, err := c.FastIdRanobe("Ookami to Koushinryou")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fr, status, err := fast_ranobe.FavoritesCreate(g.FAVORITES_LINKED_TYPE_RANOBE, "")
	//fr, status, err := fast_ranobe.FavoritesDelete("Ranobe")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fr.Success, fr.Notice)

	// add/remove favorites person.
	fast_person, status, err := c.FastIdPeople("Sumire Uesaka")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fp, status, err := fast_person.FavoritesCreate(
		g.FAVORITES_LINKED_TYPE_PERSON, g.FAVORITES_KIND_SEYU,
	)
	//fp, status, err := fast_person.FavoritesDelete("Person")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fp.Success, fp.Notice)

	// add/remove favorites character.
	fast_character, status, err := c.FastIdCharacter("Holo")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fc, status, err := fast_character.FavoritesCreate(g.FAVORITES_LINKED_TYPE_CHARACTER, "")
	//fc, status, err := fast_character.FavoritesDelete("Character")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fc.Success, fc.Notice)
}
