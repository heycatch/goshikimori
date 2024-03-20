package req

import (
  "fmt"
  "net/http"
  "context"
  "time"
  "bytes"
)

const site string = "shikimori.one/api"

func NewGetRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodGet, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  return req, cancel
}

// To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewPostRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Reorder: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewReorderPostRequestWithCancel(application, accessToken, search string,
    position int, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(`{"new_index": "%d"}`, position))
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, bytes.NewBuffer(data))
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Mark order messages: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewMarkReadPostRequestWithCancel(application, accessToken, search, ids string,
    is_read int, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(`{"ids": "%s", "is_read": "%d"}`, ids, is_read))
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, bytes.NewBuffer(data))
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Read/Delete all messages: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewReadDeleteAllPostRequestWithCancel(application, accessToken, search, name string,
    number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(`{"frontend": "false", "type": "%s"}`, name))
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, bytes.NewBuffer(data))
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Send message: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewSendMessagePostRequestWithCancel(application, accessToken, search, body string,
    from_id, to_id int, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(
    `{"frontend": "false",
    "message": {"body": "%s", "from_id": "%d", "kind": "Private", "to_id": "%d"}}`,
    body, from_id, to_id),
  )
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, bytes.NewBuffer(data))
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Change message. To work correctly with the PUT method,
// make sure that your application has all the necessary permissions.
func NewChangeMessagePutRequestWithCancel(application, accessToken, search, body string,
    number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(`{"frontend": "false", "message": {"body": "%s"}}`, body))
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPut, custom_url, bytes.NewBuffer(data))
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Delete message. To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteMessageDeleteRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodDelete, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodDelete, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  return req, cancel
}
