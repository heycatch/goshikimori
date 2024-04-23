package search

import (
  "testing"
  "sync"
)

func TestIntegerWithoutChan(t *testing.T) {
  if IntegerWithoutChan("", []string{"one", "two", "three"}) == -1 {
    t.Log("Empty target IntegerWithoutChan passed")
  } else {
    t.Error("Empty target IntegerWithoutChan failed")
  }

  if IntegerWithoutChan("bob", []string{""}) == -1 {
    t.Log("Empty list IntegerWithoutChan passed")
  } else {
    t.Error("Empty list IntegerWithoutChan failed")
  }

  if IntegerWithoutChan("bob", []string{"-1", "-1", "-1", "-1"}) == -1 {
    t.Log("Broken list IntegerWithoutChan passed")
  } else {
    t.Error("Broken list IntegerWithoutChan failed")
  }

  if IntegerWithoutChan("bob", []string{"one", "two", "bob", "three"}) == 2 {
    t.Log("Normal IntegerWithoutChan passed")
  } else {
    t.Error("Normal IntegerWithoutChan failed")
  }
}

func TestIntegerWithChan(t *testing.T) {
  var wg sync.WaitGroup
  wg.Add(4)

  ch := make(chan int)

  go IntegerWithChan(&wg, ch, "", []string{"one", "two", "three"})
  if <-ch == -1 {
    t.Log("Empty target IntegerWithChan passed")
  } else {
    t.Error("Empty target IntegerWithChan failed")
  }

  go IntegerWithChan(&wg, ch, "bob", []string{""})
  if <-ch == -1 {
    t.Log("Empty list IntegerWithChan passed")
  } else {
    t.Error("Empty list IntegerWithChan failed")
  }

  go IntegerWithChan(&wg, ch, "bob", []string{"-1", "-1", "-1", "-1"})
  if <-ch == -1 {
    t.Log("Broken list IntegerWithChan passed")
  } else {
    t.Error("Broken list IntegerWithChan failed")
  }

  go IntegerWithChan(&wg, ch, "bob", []string{"one", "two", "bob", "three"})
  if <-ch == 2 {
    t.Log("Normal IntegerWithChan passed")
  } else {
    t.Error("Normal IntegerWithChan failed")
  }

  wg.Wait()
}

func TestStringWithChan(t *testing.T) {
  var wg sync.WaitGroup
  wg.Add(4)

  ch := make(chan string)

  go StringWithChan(&wg, ch, "", []string{"one", "two", "three"})
  if <-ch == "" {
    t.Log("Empty target StringWithChan passed")
  } else {
    t.Error("Empty target StringWithChan failed")
  }

  go StringWithChan(&wg, ch, "bob", []string{""})
  if <-ch == "" {
    t.Log("Empty list StringWithChan passed")
  } else {
    t.Error("Empty list StringWithChan failed")
  }

  go StringWithChan(&wg, ch, "bob", []string{"-1", "-1", "-1", "-1"})
  if <-ch == "" {
    t.Log("Broken list StringWithChan passed")
  } else {
    t.Error("Broken list StringWithChan failed")
  }

  go StringWithChan(&wg, ch, "bob", []string{"one", "two", "bob", "three"})
  if <-ch == "bob" {
    t.Log("Normal StringWithChan passed")
  } else {
    t.Error("Normal StringWithChan failed")
  }

  wg.Wait()
}
