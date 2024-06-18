package goshikimori

import (
  "bytes"
  "context"
  "strconv"
  "net/http"
  "time"

  "github.com/heycatch/goshikimori/concat"
)

// Normal GET request with User-Agent only.
func NewGetRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(ctx, http.MethodGet, search, nil)
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  return req, cancel, nil
}

// For certain GET requests where a Bearer is needed.
func NewGetRequestWithCancelAndBearer(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(ctx, http.MethodGet, search, nil)
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  return req, cancel, nil
}

// To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewPostRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, nil)
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// Reorder: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewReorderPostRequestWithCancel(application, accessToken, search string,
    position int, number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(
    ctx, http.MethodPost, search,
    bytes.NewBuffer(concat.DataBuffer(
      []string{"{\"new_index\": ", "\"", strconv.Itoa(position), "\"", "}"},
    )),
  )
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// Mark order messages: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewMarkReadPostRequestWithCancel(application, accessToken, search, ids string,
    is_read int, number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(
    ctx, http.MethodPost, search,
    bytes.NewBuffer(concat.DataBuffer([]string{
      "{\"ids\": ", "\"", ids, "\"", ", ", "\"is_read\": ",
      "\"", strconv.Itoa(is_read), "\"", "}",
    })),
  )
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// Read/Delete all messages: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewReadDeleteAllPostRequestWithCancel(application, accessToken, search, name string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(
    ctx, http.MethodPost, search,
    bytes.NewBuffer(concat.DataCopy(
      33 + len(name),
      []string{"{\"frontend\": ", "\"false\", ", "\"type\": ", "\"", name, "\"", "}"},
    )),
  )
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// Send message: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewSendMessagePostRequestWithCancel(application, accessToken, search, body string,
    from_id, to_id int, number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(
    ctx, http.MethodPost, search,
    bytes.NewBuffer(concat.DataBuffer([]string{
      "{\"frontend\": \"false\", \"message\": {\"body\": \"", body,
      "\", \"from_id\": \"", strconv.Itoa(from_id),
      "\", \"kind\": \"Private\", \"to_id\": \"", strconv.Itoa(to_id), "\"}}",
    })),
  )
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// Change message. To work correctly with the PUT method,
// make sure that your application has all the necessary permissions.
func NewChangeMessagePutRequestWithCancel(application, accessToken, search, body string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(
    ctx, http.MethodPut, search,
    bytes.NewBuffer(concat.DataCopy(
      46 + len(body),
      []string{"{\"frontend\": \"false\", \"message\": {\"body\": \"", body, "\"}}"},
    )),
  )
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// Delete message. To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteMessageDeleteRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(ctx, http.MethodDelete, search, nil)
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  req.Header.Set("Content-Type", "application/json")
  return req, cancel, nil
}

// To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteRequestWithCancel(application, accessToken, search string,
    number time.Duration) (*http.Request, context.CancelFunc, error) {
  // In requests I set the time to 10 seconds.
  ctx, cancel := context.WithTimeout(context.Background(), number * time.Second)
  req, err := http.NewRequestWithContext(ctx, http.MethodDelete, search, nil)
  if err != nil {
    return req, cancel, err
  }
  req.Header.Add("User-Agent", application)
  req.Header.Add("Authorization", concat.Bearer(accessToken))
  return req, cancel, nil
}
