// It's a testing place, there's nothing interesting here.
package main

import (
  "fmt"
  "text/tabwriter"
  "os"
  "net/http/httputil"

  g "github.com/heycatch/goshikimori"
)

func config() *g.Configuration {
  return g.SetConfiguration(
    "",
    "",
  )
}

func main() {
  c := config()
  w, _, err := c.WhoAmi()
  if err != nil {
    fmt.Println(err)
    return
  }

  req, cancel, err := g.NewGetRequestWithCancel(c.Application, c.AccessToken, "whoami", 10)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer cancel()
  dump, err := httputil.DumpRequestOut(req, true)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(dump))

  t := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
  fmt.Fprintf(t, "%d\t", w.Id)
  fmt.Fprintf(t, "%s\t", w.Nickname)
  fmt.Fprintf(t, "%s\t", w.Locale)
  fmt.Fprintf(t, "%s\t\n", w.Last_online_at)
  t.Flush()
}
