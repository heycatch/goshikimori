package main

import (
  "fmt"
  "text/tabwriter"
  "os"
  "net/http/httputil"

  g "github.com/vexilology/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "",
    "",
  )
}

func main() {
  c := conf()
  w, err := c.WhoAmi()
  if err != nil {
    fmt.Println(err)
    return
  }

  dump, err := httputil.DumpRequestOut(c.NewGetRequest("whoami"), true)
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
