### EN | [RU](https://github.com/heycatch/goshikimori/blob/master/graphql/README_ru.md)

## At the moment, the GraphQL API is stated as experimental.

## How to use:
In the standard examples, you may notice that the implementation for customizing the
search goes through structures, but in this case we use *variant functions*.
- First parameters what we are looking for (anime title, manga title, etc.)
- The second parameter is the interface, with a strict sequence,
  accepts int - string - bool and which is already needed just for search customization.

Next, let's look at an examples:
```golang
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

  // The first parameter is the name of the anime; name: "initial d".
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
  // The available interface parameters can be found in the function description: SearchAnimeGraphql();
  a, status, err := c.SearchAnimesGraphql(
    "initial d", 5, 8, "", "tv", "released", "", "", "pg_13", "", false,
  )
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Here you can track errors received during server response.
  fmt.Println(a.Errors)
  // Standard output of our search, nothing new.
  for _, v := range a.Data.Animes {
    fmt.Println(v.Id, v.Name, v.Score, v.Episodes, v.ReleasedOn.Year)
  }
}
```
```golang
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

  // The first parameter is the name of the manga; name: "initial d".
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
  // The available interface parameters can be found in the function description: SearchMangaGraphql();
  m, status, err := c.SearchMangasGraphql(
    "initial d", 1, 8, "", "manga", "released", "", "", false,
  )
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
