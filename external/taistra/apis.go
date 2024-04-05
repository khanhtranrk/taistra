package taistra

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Apis struct {
  baseURL string
  apiKey string
}

func NewApi(cfg *Config, apiKey string) *Apis {
  return &Apis{cfg.TaistraURL, apiKey}
}

func (a *Apis) ValidateMessage(msg []byte) ([]byte, error) {
  request, err := http.NewRequest("POST", a.baseURL + "/api/v1/messages/validate", bytes.NewBuffer(msg))
  if err != nil {
    return nil, err
  }

  request.Header.Set("Content-Type", "application/octet-stream; charset=UTF-8")
  request.Header.Add("Authorization", a.apiKey)

  client := &http.Client{}
  resp, err := client.Do(request)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return nil, fmt.Errorf("ValidateMessage: status code %d", resp.StatusCode)
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return body, nil
}
