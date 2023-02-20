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
  u, err := c.SearchUser("incarnati0n")
  if err != nil {
    fmt.Println(err)
    return
  }
  if u.Id == 0 {
    fmt.Println("user not found")
    return
  }

  um, err := c.UserUnreadMessages(u.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(um.Messages, um.News, um.Notifications)
}
