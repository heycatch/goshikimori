package api

import (
  "fmt"
  "strings"
)

func Search(s string) string {
  convert := strings.Replace(s, " ", "_", -1)
  return fmt.Sprintf("?search=%s", convert)
}
