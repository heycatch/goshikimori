## Example 1
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  r, err := c.SearchUser("incarnati0n")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(r.Id, r.Last_Online, r.Sex)
}
```
## Example 2
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  r, err := c.SearchAnime("Initial D")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    fmt.Println(values.Name, values.Status, values.Score)
  }
}
```

## Example 3
``` golang
package main

import (
  "fmt"
  "log"

  g "github.com/vexilology/goshikimori/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION NAME",
    "PERSONAL KEY",
  )
}

func main() {
  c := conf()
  r, err := c.SearchManga("Initial D")
  if err != nil {
    log.Fatal(err)
  }
  for _, values := range r {
    fmt.Println(values.Name, values.Volumes, values.Chapters)
  }
}
```
