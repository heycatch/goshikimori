package goshikimori

import (
  "errors"
  "strings"
)

// String formatting for achievements search. Check [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/achievements/main.go
func NekoSearch(name string) (string, error) {
  var spaces string

  // extra spaces removed.
  words := strings.Fields(name)
  if len(words) == 0 { return "", errors.New("too short string") }
  spaces += words[0]
  for i := 1; i < len(words); i++ {
    spaces += " " + words[i]
  }

  return strings.Replace(strings.ToLower(spaces), " ", "_", -1), nil
}
