### EN | [RU](https://github.com/heycatch/goshikimori/blob/master/graphql/README_ru.md)

## At the moment, the GraphQL API is stated as experimental.

Next, let's look at an examples:
```golang
package main

import (
  "fmt"

  "github.com/heycatch/goshikimori"
  "github.com/heycatch/goshikimori/graphql"
)

func conf() *goshikimori.Configuration {
  return goshikimori.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // The first parameter is the values of the anime; values: "id", "name", "score", "episodes", "airedOn{year month day date}".
  // The second parameter is the name of the anime; name: "initial d".
  // Now let's move on to the interface:
  //    1)  limit: 5;
  //    2)  score: 8;
  //    3)  order: ""; skipped;
  //    4)  kind: "tv";
  //    5)  status: "released";
  //    6)  season: ""; skipped;
  //    7)  duration: ""; skipped;
  //    8)  rating: "pg_13";
  //    9)  mylist: ""; skipped;
  //    10) censored: false;
  //
  // The available values can be found in the function description: graphql.Values();
  // The available interface parameters can be found in the function description: graphql.AnimeSchema();
  schema, err := graphql.AnimeSchema(
    graphql.Values("id", "name", "score", "episodes", "airedOn{year month day date}"),
    "initial d",
    5, 8, "", "tv", "released", "", "", "pg_13", "", false,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  a, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Here you can track errors received during server response.
  fmt.Println(a.Errors)
  // Standard output of our search, nothing new.
  for _, v := range a.Data.Animes {
    fmt.Println(
      v.Id, v.Name, v.Score, v.Episodes, v.AiredOn.Year,
      v.AiredOn.Month, v.AiredOn.Day, v.AiredOn.Date,
    )
  }
}
```
```golang
package main

import (
  "fmt"

  "github.com/heycatch/goshikimori"
  "github.com/heycatch/goshikimori/graphql"
)

func conf() *goshikimori.Configuration {
  return goshikimori.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // The first parameter is the values of the manga; values: "id", "name", "score", "volumes", "chapters", "releasedOn{year}".
  // The second parameter is the name of the manga; name: "initial d".
  // Now let's move on to the interface:
  //    1) limit: 1;
  //    2) score: 8;
  //    3) order: ""; skipped;
  //    4) kind: "manga";
  //    5) status: "released";
  //    6) season: ""; skipped;
  //    7) mylist: ""; skipped;
  //    8) censored: false;
  //
  // The available values can be found in the function description: graphql.Values();
  // The available interface parameters can be found in the function description: graphql.MangaSchema();
  schema, err := graphql.MangaSchema(
    graphql.Values("id", "name", "score", "volumes", "chapters", "releasedOn{year}"),
    "initial d",
    1, 8, "", "manga", "released", "", "", false,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  m, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Here you can track errors received during server response.
  fmt.Println(m.Errors)
  // Standard output of our search, nothing new.
  for _, v := range m.Data.Mangas {
    fmt.Println(v.Id, v.Name, v.Score, v.Volumes, v.Chapters, v.ReleasedOn.Year)
  }
}
```
```golang
package main

import (
  "fmt"

  "github.com/heycatch/goshikimori"
  "github.com/heycatch/goshikimori/graphql"
)

func conf() *goshikimori.Configuration {
  return goshikimori.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // The first parameter is the values of the character; values: "id", "name", "russian", "url", "description".
  // The second parameter is the name of the character; name: "onizuka".
  // Now let's move on to the interface:
  //    1) page: 1;
  //    2) limit: 2;
  //
  // The available values can be found in the function description: graphql.Values();
  // The available interface parameters can be found in the function description: graphql.CharacterSchema();
  schema, err := graphql.CharacterSchema(
    graphql.Values("id", "name", "russian", "url", "description"),
    "onizuka",
    1, 2,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  m, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Here you can track errors received during server response.
  fmt.Println(m.Errors)
  // Standard output of our search, nothing new.
  for _, v := range m.Data.Characters {
    fmt.Println(v.Id, v.Name, v.Russian, v.Url, v.Description)
  }
}
```
