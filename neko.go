package goshikimori

import (
  "errors"
  "strings"

  "github.com/heycatch/goshikimori/concat"
)

// String formatting for achievements search. Check [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/achievements/main.go
func NekoSearch(name string) (string, error) {
  // Extra spaces removed.
  words := strings.Fields(name)
  if len(words) == 0 { return "", errors.New("too short string") }
  return strings.ToLower(concat.NekoSliceToString(words)), nil
}
