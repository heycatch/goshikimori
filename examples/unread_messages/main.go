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

  um, err := c.FastIdUser("incarnati0n").UserUnreadMessages()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(um.Messages, um.News, um.Notifications)
}
