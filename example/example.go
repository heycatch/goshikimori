package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "encoding/json"

  "github.com/joho/godotenv"
)

var (
  bearer    = "Bearer " + parseToken()
  url_shiki = "https://shikimori.one/api/users/whoami"
)

func parseToken() string {
  err := godotenv.Load("config.env")
  if err != nil {
    log.Fatal("Not found config.env")
  }

  result := os.Getenv("ACCESS_TOKEN")
  if result == "" {
    log.Fatal("Not found ACCESS_TOKEN")
  }

  return result
}

func parseApplication() string {
  err := godotenv.Load("config.env")
  if err != nil {
    log.Fatal("Not found config.env")
  }

  result := os.Getenv("APP_NAME")
  if result == "" {
    log.Fatal("Not found APP_NAME")
  }

  return result
}

func exampleRequest() map[string]interface{} {
  var result map[string]interface{}

  client := http.Client{}

  req, err := http.NewRequest("GET", url_shiki, nil)
  req.Header.Add("User-Agent", parseApplication())
  req.Header.Add("Authorization", bearer)
  if err != nil {
    log.Fatal(err)
  }

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  json.NewDecoder(resp.Body).Decode(&result)
  return result
}

func main() {
  fmt.Println(exampleRequest())
}
