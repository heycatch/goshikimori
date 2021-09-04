package req

type Config struct {
  Application string
  SecretKey   string
}

const (
  Get    = "GET"
  Post   = "POST"
  Put    = "PUT"
  Patch  = "PATCH"
  Delete = "DELETE"
)
