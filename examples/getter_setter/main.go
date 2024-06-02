package main

import (
  "fmt"

  g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := config()
  
  fast, status, err := c.FastIdUser("arctica")
  if status != 200 || err != nil {
    fmt.Println(err)
    return
  }

  // Getting an id.
  fmt.Println(fast.Id)
  // Additional option for receiving id.
  fmt.Println(fast.GetFastId())
  // Quick id change.
  new_id := c.SetFastId(1337)
  fmt.Println(new_id.Id)

  // Getting configuration.
  fmt.Println(c.GetConfiguration())
  // Quick configuration change.
  new_config := g.SetConfiguration("Bob", "XXX-XXX-XXX")
  fmt.Println(new_config.Application, new_config.AccessToken)
}
