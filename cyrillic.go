package goshikimori

import (
  "net/url"
  "unicode"
)

// url.QueryEscape() breaks Cyrillic with spaces, and without it,
// Latin with spaces breaks, so there is this intermediate function
// that solves this problem and there are no more errors.
func languageCheck(target string) string {
  for i := 0; i < len(target); i++ {
    if target[i] > unicode.MaxASCII {
      return target
    }
  }
  return url.QueryEscape(target)
}
