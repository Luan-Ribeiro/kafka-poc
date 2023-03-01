package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic, kafkaGroupId string
var kafkaConfig kafka.ConfigMap

func init() {
	kafkaServer = "localhost:9092"
	kafkaTopic = "purchase"
	kafkaGroupId = "poc-group"

	kafkaConfig = kafka.ConfigMap{"bootstrap.servers": kafkaServer, "group.id": kafkaGroupId, "go.events.channel.enable": true}

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)
}

func main() {
	consumer, err := kafka.NewConsumer(&kafkaConfig)
	if err != nil {
		fmt.Println("consumer not created ", err.Error())
		os.Exit(1)
	}

	err = consumer.Subscribe(kafkaTopic, nil)

	if err != nil {
		fmt.Println("Unable to subscribe to topic " + kafkaTopic + " due to error - " + err.Error())
		os.Exit(1)
	}

	fmt.Println("subscribed to topic ", kafkaTopic)

	for {
		kafkaEvent := <-consumer.Events()
		switch event := kafkaEvent.(type) {
		case *kafka.Message:
			fmt.Println("Message " + string(event.Value))
		case kafka.Error:
			fmt.Println("Consumer error ", event.String())
		case kafka.PartitionEOF:
			fmt.Println(kafkaEvent)
		}
	}
}
