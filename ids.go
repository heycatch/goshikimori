package goshikimori

import (
  "io"
  "strconv"
  "net/http"
  "encoding/json"
  "errors"

  "github.com/heycatch/goshikimori/api"
  "github.com/heycatch/goshikimori/str"
  "github.com/heycatch/goshikimori/req"
)

// Converting an array with an ids to a string.
func IdsToStirng(target []int) string {
  var res string
  for i := 0; i < len(target); i++ {
    if target[i] != 0 { res += strconv.Itoa(target[i]) + "," }
  }
  if res == "" { return "" } // check for zeros or panic error.
  return res[:len(res)-1]
}

// Name: unread message type.
//
// 'Name' settings:
//  - messages
//  - news
//  - notifications
//
// Empty array to be filled with ids for messages.
func (f *FastId) UnreadMessagesIds(name string) ([]int, error) {
  var res []int
  var um api.UnreadMessages
  var client = &http.Client{}

  get, cancel := req.NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    str.ConvertUser(f.Id, "unread_messages"), 10,
  )
  defer cancel()

  resp, err := client.Do(get)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  if err := json.Unmarshal(data, &um); err != nil {
    return nil, err
  }

  switch name {
  case "messages":
    res = make([]int, um.Messages)
  case "news":
    res = make([]int, um.News)
  case "notifications":
    res = make([]int, um.Notifications)
  default:
    return nil, errors.New("wrong name... try messages, news or notifications")
  }

  return res, nil
}
