package kafka

import (
	"fmt"
	"ip-receive-server/internal/pkg/core"
	"strings"

	"github.com/Shopify/sarama"
)

func Init() {
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = false
	configs.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	client, err := sarama.NewClient(strings.Split(core.ApplicationContent.Sink.Kafka.BrokerAddrs, ","), configs)
	if err != nil {
		panic(fmt.Sprintf("unable to create kafka client:  #%v ", err))
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(fmt.Sprintf("unable to create kafka producer:  #%v ", err))
	}

	core.SinkContent = NewProduce(&producer)
}
