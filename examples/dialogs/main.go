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
  // search all dialogs
  d, status, err := c.Dialogs()
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    if len(d) == 0 {
      fmt.Println("dialogs not found")
      return
    }
    for _, v := range d {
      fmt.Println(
        v.Target_user.Id, v.Target_user.Nickname,
        v.Target_user.Image.X160, v.Target_user.Last_online_at,
        v.Message.Id, v.Message.Kind, v.Message.Body, v.Message.Created_at,
      )
    }
  } else {
    fmt.Println(status)
  }
  // search a special dialogs
  fast, status, err := c.FastIdUser("morr")
  if err != nil {
    fmt.Println(err)
    return
  }
  if status == 200 {
    sd, err := fast.SearchDialogs()
    if err != nil {
      fmt.Println(err)
      return
    }
    if len(sd) == 0 {
      fmt.Println("dialog not found")
    }
    for _, v := range sd {
      fmt.Println(
        v.Id, v.Read, v.Body, v.Created_at,
        v.From.Id, v.From.Nickname, v.From.Image.X160,
        v.To.Id, v.To.Nickname, v.To.Image.X160,
      )
    }
    // delete a special dialogs
    dd, err := fast.DeleteDialogs()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(dd.Notice)
  } else {
    fmt.Println(status)
  }
}
