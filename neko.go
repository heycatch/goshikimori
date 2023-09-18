package goshikimori

import "strings"

// String formatting for achievements search. Check [example].
//
// [example]: https://github.com/heycatch/goshikimori/blob/master/examples/achievements/main.go
func NekoSearch(name string) string {
  var spaces string

  // extra spaces removed.
  words := strings.Fields(name)
  for i := 0; i < len(words); i++ {
    if i == 0 { spaces += words[i]; continue }
    spaces += " " + words[i]
  }

  return strings.Replace(strings.ToLower(spaces), " ", "_", -1)
}
