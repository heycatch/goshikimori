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

  fr, err := c.FastIdUser("morr").AddFriend()
  //fr, err := c.FastIdUser("morr").RemoveFriend()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(fr.Notice)
}
