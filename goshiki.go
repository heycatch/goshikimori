package main

import (
  "fmt"
  "log"

  "goshiki/api"
  "goshiki/oauth2"
)

func main() {
  found := api.Animes{Animes:"animes", Id:"22"}
  found.Other.Topics = "topics"
  result, err := oauth2.NewRequest("GET", found)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(result))
}
