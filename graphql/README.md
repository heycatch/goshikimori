## [EN](https://github.com/heycatch/goshikimori/blob/master/graphql/README_en.md) | RU

### На данный момент API GraphQL заявлен как экспериментальный.

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
  //    1)  page: 1;
  //    2)  limit: 5;
  //    3)  score: 8;
  //    4)  order: ""; пропустил;
  //    5)  kind: "tv";
  //    6)  status: "released";
  //    7)  season: ""; пропустил;
  //    8)  duration: ""; пропустил;
  //    9)  rating: "pg_13";
  //    10) mylist: ""; пропустил;
  //    11) censored: false;
  //    12) genre: nil; пропустил;
  //
  // Про доступные значения можно почитать в описании функции: graphql.Values();
  // Про доступные параметры интерфейса можно почитать в описании функции: graphql.AnimeSchema();
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
  //    1) page: 1;
  //    2) limit: 1;
  //    3) score: 8;
  //    4) order: ""; пропустил;
  //    5) kind: "manga";
  //    6) status: "released";
  //    7) season: ""; пропустил;
  //    8) mylist: ""; пропустил;
  //    9) censored: false;
  //    10) genre: nil; пропустил;
  //
  // Про доступные значения можно почитать в описании функции: graphql.Values();
  // Про доступные параметры интерфейса можно почитать в описании функции: graphql.MangaSchema();
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

  ch, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(ch.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
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

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "russian", "url", "website", "birthOn{year month day date}".
  // Вторым параметром идет имя человека; name: "satsuki".
  // Теперь переходим к интерфейсу:
  //    1) page: 1;
  //    2) limit: 2;
  //    3) isSeyu: true;
  //    4) isMangaka: false;
  //    5) isProducer: false;
  //
  // Про доступные значения можно почитать в описании функции: graphql.Values();
  // Про доступные параметры интерфейса можно почитать в описании функции: graphql.PeopleSchema();
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

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(p.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range p.Data.People {
    fmt.Println(
      v.Id, v.Name, v.Russian, v.Url, v.Website,
      v.BirthOn.Year, v.BirthOn.Month, v.BirthOn.Day, v.BirthOn.Date,
    )
  }
}
```
