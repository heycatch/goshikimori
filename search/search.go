package search

import "slices"

// Return index. If not found then -1.
func IndexInSlice(target string, slice []string) int {
  return slices.IndexFunc(slice, func(temp string) bool { return temp == target })
}

// Target: checking the validity of the parameter.
//
// SetDefault: if we didn't find anything.
//
// Slice: available variables.
func LinearComplexity(target *string, setDefault string, slice []string) {
  var valid bool
  for i := 0; i < len(slice); i++ {
    if slice[i] == *target {
      *target = slice[i]
      valid = true
      break
    }
  }
  if !valid { *target = setDefault }
}
