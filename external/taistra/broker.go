package taistra

import "github.com/streadway/amqp"

type Broker struct {
  conn *amqp.Connection
  channel *amqp.Channel
}

func NewBroker(cfg *Config) (*Broker, error) {
    conn, err := amqp.Dial(cfg.AmqpURL)
    if err != nil {
      return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
      return nil, err
    }

    return &Broker{conn: conn, channel: channel}, nil
}

func (b *Broker) Close() error {
  err := b.channel.Close()

  if err != nil {
    return err
  }

  err = b.conn.Close()

  if err != nil {
    return err
  }

  return nil
}

