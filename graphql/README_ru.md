### [EN](https://github.com/heycatch/goshikimori/blob/master/graphql/README.md) | RU

## На данный момент API GraphQL заявлен как экспериментальный.

Далее рассмотрим на примерах:
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

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "score", "episodes", "airedOn{year month day date}".
  // Вторым параметром идет название аниме; name: "initial d".
  // Теперь переходим к интерфейсу:
  //    1)  limit: 5;
  //    2)  score: 8;
  //    3)  order: ""; пропустил;
  //    4)  kind: "tv";
  //    5)  status: "released";
  //    6)  season: ""; пропустил;
  //    7)  duration: ""; пропустил;
  //    8)  rating: "pg_13";
  //    9)  mylist: ""; пропустил;
  //    10) censored: false;
  //
  // Про доступные значения можно почитать в описании функции: graphql.Values();
  // Про доступные параметры интерфейса можно почитать в описании функции: graphql.AnimeSchema();
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

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(a.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
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

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "score", "volumes", "chapters", "releasedOn{year}".
  // Вторым параметром идет название манги; name: "initial d".
  // Теперь переходим к интерфейсу:
  //    1) limit: 1;
  //    2) score: 8;
  //    3) order: ""; пропустил;
  //    4) kind: "manga";
  //    5) status: "released";
  //    6) season: ""; пропустил;
  //    7) mylist: ""; пропустил;
  //    8) censored: false;
  //
  // Про доступные значения можно почитать в описании функции: graphql.Values();
  // Про доступные параметры интерфейса можно почитать в описании функции: graphql.MangaSchema();
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

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(m.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
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

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "russian", "url", "description"".
  // Вторым параметром идет название персонажа; name: "onizuka".
  // Теперь переходим к интерфейсу:
  //    1) page: 1;
  //    2) limit: 2;
  //
  // Про доступные значения можно почитать в описании функции: graphql.Values();
  // Про доступные параметры интерфейса можно почитать в описании функции: graphql.CharacterSchema();
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

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(m.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range m.Data.Characters {
    fmt.Println(v.Id, v.Name, v.Russian, v.Url, v.Description)
  }
}
```
