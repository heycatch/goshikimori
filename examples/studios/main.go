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
  s, status, err := c.SearchStudios()
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    if len(s) == 0 {
      fmt.Println("not found studios")
      return
    }
    for _, v := range s {
      fmt.Println(v.Id, v.Name, v.Filtered_name, v.Real)
    }
  } else {
    fmt.Println(status)
  }
}
