package taistra

import "fmt"

type HttpGate struct {
  apis *Apis
}

func NewHttpGate(apis *Apis) *HttpGate {
  return &HttpGate{apis}
}

func (g *HttpGate) ValidateMessage(msg []byte) (byte, error) {
  res, err := g.apis.ValidateMessage(msg)

  if err != nil {
    return 0, err
  }

  if len(res) < 1 {
    return 0, fmt.Errorf("Response is wrong")
  }

  errCode := res[0]

  return errCode, nil
}
