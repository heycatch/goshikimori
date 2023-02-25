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

  ca, err := c.SearchConstantsAnime()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(ca.Kind, ca.Status)

  cm, err := c.SearchConstantsManga()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cm.Kind, cm.Status)

  ur, err := c.SearchConstantsUserRate()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(ur.Status)

  cc, err := c.SearchConstantsClub()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cc.Join_policy, cc.Comment_policy, cc.Image_upload_policy)

  cs, err := c.SearchConstantsSmileys()
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range cs {
    fmt.Println(v.Bbcode, v.Path)
  }
}
