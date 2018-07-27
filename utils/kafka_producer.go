package utils

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaProducer(message string) {
	fmt.Println("Inside Producer")
	kConfig := AllConfig.KafkaConfig
	fmt.Println("Kakf Config", kConfig)
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kConfig.KafkaBootStrapServer})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := kConfig.KafkaTopic
	fmt.Println("Writing to topic", topic)

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message)}, nil)
	// Wait for message deliveries
	p.Flush(15 * 1000)
}
