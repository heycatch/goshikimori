package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "io/ioutil"

  "github.com/joho/godotenv"
)

const (
  bearer   = "Bearer "
  orig_url = "https://shikimori.one/api/"
)

type TestApiVerOne struct {
  Test_conn string `json:"users/whoami"`
}

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

func exampleRequest(input TestApiVerOne) ([]byte, error) {
  req, err := http.NewRequest("GET", orig_url+input.Test_conn, nil)
  req.Header.Add("User-Agent", parseApplication())
  req.Header.Add("Authorization", bearer+parseToken())
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

func (a TestApiVerOne) String() string {
  return fmt.Sprintf("%s", a.Test_conn)
}

func main() {
  tc := TestApiVerOne{Test_conn:"users/whoami"}

  result, err := exampleRequest(tc)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(result))
}
