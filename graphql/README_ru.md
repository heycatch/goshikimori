### [EN](https://github.com/heycatch/goshikimori/blob/master/graphql/README.md) | RU

## На данный момент API GraphQL заявлен как экспериментальный.

## Как использовать:
В стандартных примерах можно заметить, что реализация для кастомизации
поиска идет через структуры, но в данном случае мы используем *вариативные функции*.
  - первым параметром идет название аниме, манги и т.д.
  - вторым параметром следует интерфейс, со строгой последовательностью,
    принимает int - string - bool и который уже нужен как раз таки для кастомизации поиска.

Далее рассмотрим на примере:
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

  // Первым параметром идет название аниме; name: "initial d".
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
  // Про доступные параметры интерфейса можно почитать к описанию функции: SearchAnimeGraphql();
  a, status, err := c.SearchAnimeGraphql(
    "initial d", 5, 8, "", "tv", "released", "", "", "pg_13", "", false,
  )
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(a.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range a.Data.Animes {
    fmt.Println(v.Id, v.Name, v.Score, v.Episodes, v.ReleasedOn.Year)
  }
}
```
