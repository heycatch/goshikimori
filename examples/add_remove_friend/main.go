package main

import (
  "fmt"

  g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := config()
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
