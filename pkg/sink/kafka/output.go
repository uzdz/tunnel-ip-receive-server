package kafka

import (
	"encoding/json"
	"ip-receive-server/internal/app/sink"
	"log"

	"github.com/Shopify/sarama"
)

var TopicKey = "topic"

type service struct {
	producer *sarama.AsyncProducer
}

func (s *service) Send(event interface{}) {
	eventStr, err := json.Marshal(event)

	if err != nil {
		log.Println("[处理失败] JSON解析失败：", event)
	}

	var topic = ""
	var key = ""

	(*s.producer).Input() <- &sarama.ProducerMessage{
		Topic: topic, Key: sarama.StringEncoder(key),
		Value: sarama.StringEncoder(eventStr)}
}

func (s *service) Close() {
	(*s.producer).Close()
}

func NewProduce(producer *sarama.AsyncProducer) sink.Sink {
	return &service{
		producer: producer,
	}
}
