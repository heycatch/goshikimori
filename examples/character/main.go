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

	// search characters
	chs, status, err := c.SearchCharacters("D")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range chs {
		fmt.Println(v.Id, v.Name, v.Russian)
	}

	// search character anime
	anime, status, err := c.FastIdCharacter("D")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	ch_anime, err := anime.SearchCharacter()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ch_anime.Id, ch_anime.Name, ch_anime.Altname, ch_anime.Description)
	for _, v := range ch_anime.Animes {
		fmt.Println(v.Id, v.Name, v.Score)
	}

	// search character manga
	manga, status, err := c.FastIdCharacter("Shinichi Akiyama")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	ch_manga, err := manga.SearchCharacter()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ch_manga.Id, ch_manga.Name, ch_manga.Altname, ch_manga.Description)
	for _, v := range ch_manga.Mangas {
		fmt.Println(v.Id, v.Name, v.Score, v.Chapters, v.Volumes)
	}
}
