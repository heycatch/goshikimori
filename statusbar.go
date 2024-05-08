package goshikimori

import "fmt"

// Status bar serves to slow down requests in tests.
type StatusBar struct {
  Percent, Cur, Total int
  Rate, Graph string
}

func (s *StatusBar) newOption(start, end int) {
  s.Cur = start
  s.Total = end
  s.Percent = s.getPercent()

  if s.Graph == "" { s.Graph = "#" }

  for i := 0; i < s.Percent; i += 1 { s.Rate += s.Graph }
}

func (s *StatusBar) getPercent() int {
  return int((float32(s.Cur) / float32(s.Total)) * 100)
}

func (s *StatusBar) play(cur int) {
  s.Cur = cur
  last := s.Percent
  s.Percent = s.getPercent()

  if s.Percent != last && s.Percent%2 == 0 { s.Rate += s.Graph }

  fmt.Printf("\r[%-5s]%3d%% %8d/%d", s.Rate, s.Percent, s.Cur, s.Total)
}
