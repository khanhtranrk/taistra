package taistra

import "github.com/streadway/amqp"

type AmqpGate struct {
  bk *Broker
}

func (g *AmqpGate) SendMessage(msg []byte) error {
  typ := []byte{1}
  tmsg := append(typ, msg...)

  return g.bk.channel.Publish("", "taistra", false, false, amqp.Publishing{
      ContentType: "application/octet-stream",
      Body: tmsg,
  })
}
