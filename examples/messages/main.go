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

  o := &g.Options{Type: "news", Page: "1", Limit: "1"}
  m, err := c.FastIdUser("incarnati0n").UserMessages(o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(m) == 0 {
    fmt.Println("messages not found")
    return
  }
  for _, v := range m {
    fmt.Println(v.Id, v.Kind, v.HTMLBody, v.Created_at)
  }
}
