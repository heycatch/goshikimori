package main

import (
  "fmt"
  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()
  s, err := c.SearchStudios()
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(s) == 0 {
    fmt.Println("not found studios")
    return
  }
  for _, v := range s {
    fmt.Println(v.Id, v.Name, v.Filtered_name, v.Real)
  }
}