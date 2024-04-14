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
  fr, err := fast.AddFriend()
  //fr, err := fast.RemoveFriend()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(fr.Notice)
}
