package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	go consumeMessages("goapp-consumer1")
	go consumeMessages("goapp-consumer2")
	go consumeMessages("goapp-consumer3")

	// Mantém o programa em execução
	select {}
}

func consumeMessages(clientID string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-go_kafka_1:9092",
		"client.id":         clientID,
		"group.id":          "goapp-group1",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error creating consumer", err.Error())
		return
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[%s] %s\n", clientID, string(msg.Value))
		}
	}
}
