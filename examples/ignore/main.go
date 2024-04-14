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
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  ignore, err := fast.AddIgnoreUser()
  //ignore, err := fast.RemoveIgnoreUser()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(ignore.User_id, ignore.Is_ignored)
}
