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
  o := &g.Options{Page: "1", Limit: "10"}
  u, err := c.SearchUsers("angel", o)
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(u) == 0 {
    fmt.Println("users not found")
    return
  }

  var ignore_id int
  var access bool
  for _, v := range u {
    if !access {
      // Вытащить первый Id. Ведь при лимите в 1 на выходе
      // имеем всегда 2 результата(вопросы почему к шикимори).
      ignore_id = v.Id
      access = true
    }
  }
  status, ignore, err := c.AddIgnoreUser(ignore_id)
  //status, ignore, err := c.RemoveIgnoreUser(ignore_id)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(status, ignore.User_id, ignore.Is_ignored)
}
