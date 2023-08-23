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
  fast, status, err := c.FastIdUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    um, err := fast.UserUnreadMessages()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(um.Messages, um.News, um.Notifications)
  } else {
    fmt.Println(status)
  }
}
