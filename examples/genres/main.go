package main

import (
	"fmt"

	g "github.com/heycatch/goshikimori"
	"github.com/heycatch/goshikimori/concat"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()
	genres, status, err := c.SearchGenres(g.GENRES_ANIME)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(genres) == 0 {
		fmt.Println("not found genres")
		return
	}
	for _, v := range genres {
		fmt.Println(v.Id, v.Name, v.Russian, v.Kind, v.Entry_type)
	}
	// A small map helper.
	m := concat.GenerateGenres(g.GENERATE_GENRES_ANIME, genres)
	fmt.Println(m)
}
