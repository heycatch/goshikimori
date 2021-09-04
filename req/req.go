package req

type Config struct {
  Application string `json:"App"`
  SecretKey   string `json:"Key"`
}

const (
  Get    = "GET"
  Post   = "POST"
  Put    = "PUT"
  Patch  = "PATCH"
  Delete = "DELETE"
)
