package goshikimori

import (
  "net/url"
  "unicode"
)

// Comparison of Cyrillic and Latin alphabet.
// Search sometimes breaks, needs checking.
func languageCheck(target string) string {
  for i := 0; i < len(target); i++ {
    if target[i] > unicode.MaxASCII {
      return target
    }
  }
  return url.QueryEscape(target)
}
