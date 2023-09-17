package req

import (
  "fmt"
  "net/http"
  "context"
  "time"
  "bytes"
)

const site string = "shikimori.me/api"

func NewGetRequestWithCancel(application, accessToken, search string, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodGet, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  return req, cancel
}

// To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewPostRequestWithCancel(application, accessToken, search string, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// Custom POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewCustomPostRequestWithCancel(application, accessToken, search string, position int, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  data := []byte(fmt.Sprintf(`"new_index": "%d"`, position))
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, custom_url, bytes.NewBuffer(data))
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  req.Header.Set("Content-Type", "application/json")
  return req, cancel
}

// To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteRequestWithCancel(application, accessToken, search string, number time.Duration) (*http.Request, context.CancelFunc) {
  custom_url := fmt.Sprintf("https://%s/%s", site, search)
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second) // number->10seconds
  req, _ := http.NewRequestWithContext(ctx, http.MethodDelete, custom_url, nil)
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", "Bearer " + accessToken)
  return req, cancel
}
