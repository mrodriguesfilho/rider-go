package event_kafka_adapter

import (
	"encoding/json"
	"fmt"
	"rider-go/internal/domain/domainEvent"

	"github.com/IBM/sarama"
)

type EventBrokerKafkaAdapter struct {
	producer sarama.SyncProducer
}

func NewEventBrokerKafkaAdapter(brokers []string) (*EventBrokerKafkaAdapter, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		return nil, err
	}

	return &EventBrokerKafkaAdapter{
		producer: producer,
	}, nil
}

func (e *EventBrokerKafkaAdapter) Publish(event domainEvent.DomainEventInterface) {

	topic, err := e.getTopic(event.GetEventName())

	if err != nil {
		return
	}

	messageValue, _ := json.Marshal(event)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageValue),
	}

	_, _, err = e.producer.SendMessage(msg)

	if err != nil {
		panic(err)
	}
}

func (e *EventBrokerKafkaAdapter) getTopic(eventName string) (string, error) {
	switch eventName {
	case domainEvent.RideAccepetedEventName, domainEvent.RideFinishedEventName:
		return "ride-domain-event-topic", nil
	}

	return "", fmt.Errorf("no topic found for the given event")
}
