package search

import "testing"

func TestIndexSlice(t *testing.T) {
  if IndexInSlice("", []string{"one", "two", "three"}) == -1 {
    t.Log("Empty target IndexSlice passed")
  } else {
    t.Error("Empty target IndexSlice failed")
  }

  if IndexInSlice("bob", []string{""}) == -1 {
    t.Log("Empty list IndexSlice passed")
  } else {
    t.Error("Empty list IndexSlice failed")
  }

  if IndexInSlice("bob", []string{"-1", "-1", "-1", "-1"}) == -1 {
    t.Log("Broken list IndexSlice passed")
  } else {
    t.Error("Broken list IndexSlcie failed")
  }

  if IndexInSlice("bob", []string{"one", "two", "bob", "three"}) == 2 {
    t.Log("Normal IndexSlice passed")
  } else {
    t.Error("Normal IndexSlice failed")
  }
}

func TestLinearComplexity(t *testing.T) {
  type LinearTemp struct {
    Empty_string, Empty_list, Negative_numbers, Normal string
  }
  var l LinearTemp

  LinearComplexity(&l.Empty_string, "", []string{"one", "two", "three"})
  if l.Empty_string == "" {
    t.Log("Empty string LinearComplexity passed")
  } else {
    t.Error("Empty string LinearComplexity passed")
  }

  l.Empty_list = "bob"
  LinearComplexity(&l.Empty_list, "", []string{""})
  if l.Empty_list == "" {
    t.Log("Empty list LinearComplexity passed")
  } else {
    t.Error("Empty list LinearComplexity failed")
  }

  l.Negative_numbers = "bob"
  LinearComplexity(&l.Negative_numbers, "", []string{"-1", "-1", "-1"})
  if l.Negative_numbers == "" {
    t.Log("Broken list LinearComplexity passed")
  } else {
    t.Error("Broken list LinearComplexity failed")
  }

  l.Normal = "bob"
  LinearComplexity(&l.Normal, "", []string{"one", "two", "bob", "three"})
  if l.Normal == "bob" {
    t.Log("Normal LinearComplexity passed")
  } else {
    t.Error("Normal LinearComplexity failed")
  }
}
