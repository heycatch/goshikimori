package goshikimori

import "testing"

func TestIdsToString(t *testing.T) {
  if IdsToStirng([]int{}) == "" {
    t.Log("IdsTostring passed")
  } else {
    t.Error("IdsToString failed")
  }

  if IdsToStirng([]int{0, 0}) == "" {
    t.Log("IdsTostring passed")
  } else {
    t.Error("IdsToString failed")
  }

  if IdsToStirng([]int{1, 2, 3}) == "1,2,3" {
    t.Log("IdsTostring passed")
  } else {
    t.Error("IdsToString failed")
  }
}
