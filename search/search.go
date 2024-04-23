package search

import (
  "slices"
  "sync"
)

// Return index. If not found then -1.
func IntegerWithoutChan(target string, list []string) int {
  return slices.IndexFunc(list, func(temp string) bool { return temp == target })
}

// Return chan int(index). If not found then chan int(-1).
func IntegerWithChan(wg *sync.WaitGroup, ch chan int, target string, list []string) {
  defer wg.Done()
  ch <- slices.IndexFunc(list, func(s string) bool { return s == target })
}

// Return chan string(target). If not found then chan string(empty).
func StringWithChan(wg *sync.WaitGroup, ch chan string, target string, list []string) {
  defer wg.Done()
  i := slices.IndexFunc(list, func(s string) bool { return s == target })
  if i == -1 { ch <- "" } else { ch <- list[i] }
}
