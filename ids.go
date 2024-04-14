package goshikimori

import (
  "io"
  "strconv"
  "net/http"
  "encoding/json"
  "errors"
  "strings"

  "github.com/heycatch/goshikimori/api"
)

// Converting an array with an ids to a string.
func IdsToString(target []int) string {
  var res string
  for i := 0; i < len(target); i++ {
    if target[i] != 0 { res += strconv.Itoa(target[i]) + "," }
  }
  return strings.TrimSuffix(res, ",")
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
  var um api.UnreadMessages
  var client = &http.Client{}

  get, cancel := NewGetRequestWithCancel(
    f.Conf.Application, f.Conf.AccessToken,
    "users/" + strconv.Itoa(f.Id) + "/unread_messages", 10,
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
  case "messages": return make([]int, um.Messages), nil
  case "news": return make([]int, um.News), nil
  case "notifications": return make([]int, um.Notifications), nil
  default: return nil, errors.New("wrong name... try messages, news or notifications")
  }
}
