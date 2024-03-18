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
  // show user messages
  fast, status, err := c.FastIdUser("arctica")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  m, err := fast.UserMessages(&g.Options{Type: "news", Page: "1", Limit: "10"})
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(m) == 0 {
    fmt.Println("messages not found")
    return
  }
  for _, v := range m {
    fmt.Println(v.Id, v.Kind, v.HTMLBody, v.Created_at)
  }
  // show user unread messages counts
  um, err := fast.UserUnreadMessages()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(um.Messages, um.News, um.Notifications)
}
