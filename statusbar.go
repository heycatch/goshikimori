package goshikimori

import (
  "fmt"
  "time"
)

// Status bar serves to slow down requests in tests.
type conv float32

type StatusBar struct {
  Percent, Current, Total int
  Rate, Graph string
  Wait time.Duration
}

func (s *StatusBar) settings(length int, symbol string, wait time.Duration) {
  s.Total = length
  s.Graph = symbol
  s.Wait = wait
}

func (s *StatusBar) run() {
  for i := 0; i <= s.Total; i++ {
    s.Current = i
    last := s.Percent
    s.Percent = int((conv(s.Current)/conv(s.Total)) * 100)

    if s.Percent != last && s.Percent%2 == 0 { s.Rate += s.Graph }

    fmt.Printf("\r[%-5s] %5d%% %5d/%d", s.Rate, s.Percent, s.Current, s.Total)

    time.Sleep(s.Wait * time.Second)
  }
  fmt.Println()
}
