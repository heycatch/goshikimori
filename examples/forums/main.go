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
  f, status, err := c.SearchForums()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  if len(f) == 0 {
    fmt.Println("not found forums")
    return
  }
  for _, v := range f {
    fmt.Println(v.Id, v.Position, v.Name, v.Permalink, v.Url)
  }
}
