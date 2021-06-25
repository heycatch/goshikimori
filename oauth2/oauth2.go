package oauth2

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"

  "goshiki/api"
)

const (
  oauth_app    = ""
  access_token = ""
  bearer       = "Bearer "
  urlOrig      = "%s://%s"
  protocol     = "https"
  urlShiki     = "shikimori.one/api/"
)

func NewRequest(method string, input api.Animes) ([]byte, error) {
  req, err := http.NewRequest(method, fmt.Sprintf(urlOrig,
    protocol, urlShiki)+input.Animes+"/"+input.Id+"/"+input.
    Other.Roles+"/"+input.Other.Similar+"/"+input.
    Other.Related+"/"+input.Other.Screenshots+"/"+input.
    Other.Franchise+"/"+input.Other.External_links+"/"+input.
    Other.Topics, nil)
  req.Header.Add("User-Agent", oauth_app)
  req.Header.Add("Authorization", bearer+access_token)
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
