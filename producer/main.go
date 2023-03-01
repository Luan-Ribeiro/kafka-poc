package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic string
var kafkaConfig kafka.ConfigMap

func init() {
	kafkaServer = "localhost:9092"
	kafkaTopic = "purchase"
	kafkaConfig = kafka.ConfigMap{
		"bootstrap.servers": kafkaServer}

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)
}

func main() {
	p, err := kafka.NewProducer(&kafkaConfig)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for _, item := range items {
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
			Value:          []byte(item),
		}, nil)
		if err != nil {
			fmt.Println("unable to enqueue message ")
		}

		event := <-p.Events()

		message := event.(*kafka.Message)

		if message.TopicPartition.Error != nil {
			fmt.Println("Delivery failed due to error ", message.TopicPartition.Error)
		} else {
			fmt.Println("Delivered message to offset " + message.TopicPartition.Offset.String() + " in partition " + message.TopicPartition.String())
		}
	}

	p.Close()
}
