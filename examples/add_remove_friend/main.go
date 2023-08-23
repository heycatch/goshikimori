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
    fr, err := fast.AddFriend()
    //fr, err := fast.RemoveFriend()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(fr.Notice)
  } else {
    fmt.Println(status)
  }
}
