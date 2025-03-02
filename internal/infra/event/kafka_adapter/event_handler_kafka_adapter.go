package event_kafka_adapter

import (
	"encoding/json"
	"fmt"
	"rider-go/internal/application/event"
	"rider-go/internal/domain/domainEvent"

	"github.com/IBM/sarama"
)

type EventHandlerKafkaAdapter struct {
	consumer sarama.Consumer
	handlers map[string]event.EventHandler
	admin    sarama.ClusterAdmin
	topics   []string
}

func NewEventHandlerKafkaAdapter(brokers []string, topics []string, handlers map[string]event.EventHandler) (*EventHandlerKafkaAdapter, error) {

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	admin, err := sarama.NewClusterAdmin(brokers, config)

	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		return nil, err
	}

	return &EventHandlerKafkaAdapter{
		consumer: consumer,
		handlers: handlers,
		admin:    admin,
		topics:   topics,
	}, nil
}

func (e *EventHandlerKafkaAdapter) Listen() {
	for _, topic := range e.topics {
		partitionConsumer, err := e.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)

		if err != nil {
			panic(err)
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()

			for msg := range partitionConsumer.Messages() {
				var event domainEvent.DomainEventType
				json.Unmarshal(msg.Value, &event)
				if handler, exists := e.handlers[event.GetEventName()]; exists {
					handler.Handle(&event)
				} else {
					fmt.Printf("No handler found for the given event")
				}
			}
		}(partitionConsumer)
	}
}

func (k *EventHandlerKafkaAdapter) CreateTopics(topics []string) {
	for _, topic := range topics {
		err := k.admin.CreateTopic(topic, &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		}, false)

		if err != nil {
			fmt.Println(err)
		}
	}
}
