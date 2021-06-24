package main

import (
  "fmt"
  "log"

  "goshiki/api"
  "goshiki/oauth2"
)

func goshiki() {
  q := api.Animes{Animes:"animes", Id:"22"}
  q.Other.Topics = "topics"
  result, err := oauth2.NewRequest(q)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(result))
}

func main() {
  goshiki()
}
