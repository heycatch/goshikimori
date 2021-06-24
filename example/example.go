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

type TestApi struct {
  Test_conn string `json:"users/whoami"`
}

type AnimesApi struct {
  Animes string `json:"animes/"`
  Id     string `json:":id/"`
  OtherAnimes struct {
    Roles          string `json:"roles"`
    Similar        string `json:"similar"`
    Related        string `json:"related"`
    Screenshots    string `json:"screenshots"`
    Franchise      string `json:"franchise"`
    External_links string `json:"external_link"`
    Topics         string `json:"topics"`
  }
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

func exampleRequest(input AnimesApi) ([]byte, error) {
  req, err := http.NewRequest("GET", orig_url+input.
    Animes+input.Id+input.OtherAnimes.Roles+input.
    OtherAnimes.Similar+input.OtherAnimes.Related+input.
    OtherAnimes.Screenshots+input.OtherAnimes.Franchise+input.
    OtherAnimes.External_links+input.OtherAnimes.Topics, nil)
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

func (a TestApi) StringTest() string {
  return fmt.Sprintf("%s", a.Test_conn)
}

func (n AnimesApi) StringAnimes() string {
  return fmt.Sprintf("%s %s %s %s %s %s %s %s %s", n.Animes, n.Id,
    n.OtherAnimes.Roles, n.OtherAnimes.Similar,
    n.OtherAnimes.Related, n.OtherAnimes.Screenshots,
    n.OtherAnimes.Franchise, n.OtherAnimes.External_links,
    n.OtherAnimes.Topics)
}

func main() {
  gg := AnimesApi{Animes:"animes/", Id:"24/"}
  gg.OtherAnimes.Topics = "topics"
  result, err := exampleRequest(gg)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(result))
}
