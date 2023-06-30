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

  status, ignore, err := c.FastIdUser("morr").AddIgnoreUser()
  //status, ignore, err := c.FastIdUser("morr").RemoveIgnoreUser()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(status, ignore.User_id, ignore.Is_ignored)
}
