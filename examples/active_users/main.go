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
  result, status, err := c.ActiveUsers()
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 { fmt.Println(result) }
}
