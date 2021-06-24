// Shikimori API 1.0
package api

import "fmt"

type Animes struct {
  Animes string `json:"animes"`
  Id     string `json:":id"`
  Other struct {
    Roles          string `json:"roles"`
    Similar        string `json:"similar"`
    Related        string `json:"related"`
    Screenshots    string `json:"screenshots"`
    Franchise      string `json:"franchise"`
    External_links string `json:"external_link"`
    Topics         string `json:"topics"`
  }
}

func (a Animes) StringAnimes() string {
  return fmt.Sprintf("%s %s %s %s %s %s %s %s %s", a.Animes, a.Id,
    a.Other.Roles, a.Other.Similar,
    a.Other.Related, a.Other.Screenshots,
    a.Other.Franchise, a.Other.External_links,
    a.Other.Topics)
}
