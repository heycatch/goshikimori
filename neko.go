package goshikimori

import "strings"

// String formatting for achievements search. Check [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/achievements/main.go
func NekoSearch(name string) string {
  return strings.Replace(strings.ToLower(name), " ", "_", -1)
}
