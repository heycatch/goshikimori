package main

import (
  "fmt"
  g "github.com/heycatch/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()
  fast_anime, status, err := c.FastIdAnime("Naruto")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  // add/remove favorites anime.
  fa, err := fast_anime.FavoritesCreate("Anime", "")
  //fa, err := fast_anime.FavoritesDelete("Anime")
  if err != nil {
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
  fm, err := fast_manga.FavoritesCreate("Manga", "")
  //fm, err := fast_manga.FavoritesDelete("Manga")
  if err != nil {
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
  fr, err := fast_ranobe.FavoritesCreate("Ranobe", "")
  //fr, err := fast_ranobe.FavoritesDelete("Ranobe")
  if err != nil {
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
  fp, err := fast_person.FavoritesCreate("Person", "seyu")
  //fp, err := fast_person.FavoritesDelete("Person")
  if err != nil {
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
  fc, err := fast_character.FavoritesCreate("Character", "")
  //fc, err := fast_character.FavoritesDelete("Character")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(fc.Success, fc.Notice)
}
