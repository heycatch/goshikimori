package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "io/ioutil"

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

func exampleRequest() ([]byte, error) {
  req, err := http.NewRequest("GET", url_shiki, nil)
  req.Header.Add("User-Agent", parseApplication())
  req.Header.Add("Authorization", bearer)
  if err != nil {
    log.Fatal(err)
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  return ioutil.ReadAll(resp.Body)
}

func main() {
  result, err := exampleRequest()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(result))
}
