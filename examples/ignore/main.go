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
  fast, status, err := c.FastIdUser("morr")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    ignore, err := fast.AddIgnoreUser()
    //status, err := fast.RemoveIgnoreUser()
    if err != nil {
      fmt.Println(err)
      return
    }
    if status == 200 { fmt.Println(ignore.User_id, ignore.Is_ignored) }
  } else {
    fmt.Println(status)
  }
}
