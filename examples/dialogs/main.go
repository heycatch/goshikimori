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

  // search all dialogs
  d, err := c.Dialogs()
  if err != nil {
    fmt.Println(err)
    return
  }
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
  // search a special dialogs
  sd, err := c.FastIdUser("morr").SearchDialogs()
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
  status, dd, err := c.FastIdUser("morr").DeleteDialogs()
  if err != nil {
    fmt.Println(status, err)
    return
  }
  fmt.Println(status, dd.Notice)
}
