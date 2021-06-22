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
  applic = "Api Test"
  bearer = "Bearer " + parseToken()
)

func parseToken() string {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Not found .env file")
  }

  result := os.Getenv("ACCESS_TOKEN")
  if result == "" {
    log.Fatal("ACCESS TOKEN not found")
  }

  return result
}

func main() {
  var result map[string]interface{}
  var URL_SHIKIMORI = "https://shikimori.one/api/users/whoami"

  client := http.Client{}

  req, err := http.NewRequest("GET", URL_SHIKIMORI, nil)
  req.Header.Add("User-Agent", applic)
  req.Header.Add("Authorization", bearer)
  if err != nil {
    log.Fatal(err)
  }

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  json.NewDecoder(resp.Body).Decode(&result)
  fmt.Println(result)
}
