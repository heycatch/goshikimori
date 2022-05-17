```golang
package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  e := &g.Extra{
    Limit: "2", Kind: "", Status: "released",
    Season: "199x", Score: "", Rating: "",
  }
  a := c.ExtraSearchAnime("Initial D", e)
  for _, v := range a {
    fmt.Println(v.Name, v.Released_on, v.Score)
  }
}
```

``` golang
package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  s := c.SearchAnime("Initial D")
  fmt.Println(s.Name, s.Status, s.Score)
}
```
``` golang
package main

import (
  "os"
  "fmt"

  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  u := c.SearchUser("incarnati0n")
  r := c.SearchAchievement(u.Id)
  for _, v := range r {
    if v.Neko_id == g.NekoSearch("Initial D") {
      fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
      fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
    } else {
      fmt.Println("Achievement not found")
      os.Exit(1)
    }
  }
}
```
``` golang
package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  a := c.SearchAnime("Initial D")
  r := c.SearchRelatedAnime(a.Id)
  fmt.Println(r.Relation, r.Relation_Russian, r.Anime.Score)
}
```
