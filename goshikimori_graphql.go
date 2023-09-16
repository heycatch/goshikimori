package goshikimori

import (
	"encoding/json"
	"io"
	"net/http"

  "github.com/heycatch/goshikimori/api"
  "github.com/heycatch/goshikimori/req"
  "github.com/heycatch/goshikimori/graphql"
)

// Name: anime name.
//
// Exclamation mark(!) indicates ignore.
//
// Order: it's not working at the moment.
//
// If 'Options' empty fields
// 	- Limit: 1;
//	- Score: 1;
//  - Order: empty field;
//	- Kind: empty field;
//	- Status: empty field;
//	- Season: empty field;
//	- Duration: empty field;
//	- Rating: empty field;
//	- Mylist: empty field;
//	- Censored: false;
//
// 'Options' settings
//	- Limit: 50 maximum;
//	- Order: id, ranked, kind, popularity, name, aired_on, episodes, statust; random has been moved to a separate function, check [RandomAnime];
//	- Kind: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48, !tv, !movie, !ova, !ona, !special, !music, !tv_13, !tv_24, !tv_48;
//	- Status: anons, ongoing, released, !anons, !ongoing, !released;
//	- Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//	- Score: 1-9 maximum;
//  - Duration: S, D, F, !S, !D, !F;
//	- Rating: none, g, pg, pg_13, r, r_plus, rx, !g, !pg, !pg_13, !r, !r_plus, !rx;
//	- Censored: true, false;
//	- Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func (c *Configuration) SearchAnimeGraphql(name string, options ...interface{}) (api.GraphQL, int, error) {
	var client = &http.Client{}
	var g api.GraphQL

	schema, err := graphql.AnimeSchema(name, options...)
	if err != nil {
		return g, 0, err
	}

	post, cancel := req.NewPostRequestWithCancel(
		c.Application, c.AccessToken, schema, 10,
	)
	defer cancel()

	resp, err := client.Do(post)
	if err != nil {
		return g, resp.StatusCode, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return g, resp.StatusCode, err
	}

	if err := json.Unmarshal(data, &g); err != nil {
		return g, resp.StatusCode, err
	}

	return g, resp.StatusCode, nil
}

// Name: manga name.
//
// Order: it's not working at the moment.
//
// Exclamation mark(!) indicates ignore.
//
// 'Season' is not working at the moment. To skip this parameter for now.
//
// If 'Options' empty fields
// 	- Limit: 1;
//  - Order: empty field;
//	- Kind: empty field;
//	- Status: empty field;
//	- Season: empty field;
//	- Score: empty field;
//  - Censored: false;
//  - Mylist: empty field;
//
// 'Options' settings
//	- Limit: 50 maximum;
//	- Order: id, ranked, kind, popularity, name, aired_on, volumes, chapters, status; random has been moved to a separate function, check [RandomManga];
//	- Kind: manga, manhwa, manhua, light_novel, novel, one_shot, doujin, !manga, !manhwa, !manhua, !light_novel, !novel, !one_shot, !doujin;
//	- Status: anons, ongoing, released, paused, discontinued, !anons, !ongoing, !released, !paused, !discontinued;
//	- Season: 198x, 199x, 2000_2010, 2010_2014, 2015_2019, 2020_2021, 2022, 2023, !198x, !199x, !2000_2010, !2010_2014, !2015_2019, !2020_2021, !2022, !2023;
//	- Score: 1-9 maximum;
//	- Censored: true, false;
//	- Mylist: planned, watching, rewatching, completed, on_hold, dropped;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/heycatch/goshikimori/blob/master/graphql/README.md
func (c *Configuration) SearchMangaGraphql(name string, options ...interface{}) (api.GraphQL, int, error) {
	var client = &http.Client{}
	var g api.GraphQL

	schema, err := graphql.MangaSchema(name, options...)
	if err != nil {
		return g, 0, err
	}

	post, cancel := req.NewPostRequestWithCancel(
		c.Application, c.AccessToken, schema, 10,
	)
	defer cancel()

	resp, err := client.Do(post)
	if err != nil {
		return g, resp.StatusCode, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return g, resp.StatusCode, err
	}

	if err := json.Unmarshal(data, &g); err != nil {
		return g, resp.StatusCode, err
	}

	return g, resp.StatusCode, nil
}
