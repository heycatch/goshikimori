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

  ignore, status, err := c.FastIdUser("morr").AddIgnoreUser()
  //status, ignore, err := c.FastIdUser("morr").RemoveIgnoreUser()
  if err != nil {
    fmt.Println(err)
    return
  }

  if status == 200 { fmt.Println(ignore.User_id, ignore.Is_ignored) }
}
