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
  //    1)  page: 1;
  //    2)  limit: 5;
  //    3)  score: 8;
  //    4)  order: ""; skipped;
  //    5)  kind: "tv";
  //    6)  status: "released";
  //    7)  season: ""; skipped;
  //    8)  duration: ""; skipped;
  //    9)  rating: "pg_13";
  //    10) mylist: ""; skipped;
  //    11) censored: false;
  //    12) genre: nil; skipped
  //
  // The available values can be found in the function description: graphql.Values();
  // The available interface parameters can be found in the function description: graphql.AnimeSchema();
  schema, err := graphql.AnimeSchema(
    graphql.Values("id", "name", "score", "episodes", "airedOn{year month day date}"),
    "initial d",
    1, 5, 8, "", "tv", "released", "", "", "pg_13", "", false, nil,
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
  //    1) page: 1;
  //    2) limit: 1;
  //    3) score: 8;
  //    4) order: ""; skipped;
  //    5) kind: "manga";
  //    6) status: "released";
  //    7) season: ""; skipped;
  //    8) mylist: ""; skipped;
  //    9) censored: false;
  //    10) genre: nil; skipped;
  //
  // The available values can be found in the function description: graphql.Values();
  // The available interface parameters can be found in the function description: graphql.MangaSchema();
  schema, err := graphql.MangaSchema(
    graphql.Values("id", "name", "score", "volumes", "chapters", "releasedOn{year}"),
    "initial d",
    1, 1, 8, "", "manga", "released", "", "", false, nil,
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

  ch, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Here you can track errors received during server response.
  fmt.Println(ch.Errors)
  // Standard output of our search, nothing new.
  for _, v := range ch.Data.Characters {
    fmt.Println(v.Id, v.Name, v.Russian, v.Url, v.Description)
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

  // The first parameter is the values of the people; values: "id", "name", "russian", "url",
  // "website", "birthOn{year month day date}".
  // The second parameter is the name of the people; name: "satsuki".
  // Now let's move on to the interface:
  //    1) page: 1;
  //    2) limit: 2;
  //    3) isSeyu: true;
  //    4) isMangaka: false;
  //    5) isProducer: false;
  //
  // The available values can be found in the function description: graphql.Values();
  // The available interface parameters can be found in the function description: graphql.PeopleSchema();
  schema, err := graphql.PeopleSchema(
    graphql.Values("id", "name", "russian", "url", "website", "birthOn{year month day date}"),
    "satsuki",
    1, 1, true, false, false,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  p, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Here you can track errors received during server response.
  fmt.Println(p.Errors)
  // Standard output of our search, nothing new.
  for _, v := range p.Data.People {
    fmt.Println(
      v.Id, v.Name, v.Russian, v.Url, v.Website,
      v.BirthOn.Year, v.BirthOn.Month, v.BirthOn.Day, v.BirthOn.Date,
    )
  }
}
```
