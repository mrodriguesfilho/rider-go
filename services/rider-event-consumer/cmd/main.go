package main

import (
	"rider-go/internal/application/event"
	event_handlers "rider-go/internal/application/event/handlers"
	"rider-go/internal/domain/domainEvent"
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	event_kafka_adapter "rider-go/internal/infra/event/kafka_adapter"
	"rider-go/internal/infra/payment"

	"github.com/google/uuid"
)

func main() {

	kafkaAddrs := []string{"localhost:9092"}
	topics := []string{"ride-domain-event-topic"}

	// producer, err := eventAdapters.NewEventBrokerKafkaAdapter(kafkaAddrs)

	// if err != nil {
	// 	panic(err)
	// }

	// producer.Publish(domainEvent.NewRideAcceptedEvent(uuid.New()))
	accountInMemoryDB := make([]entity.Account, 0)
	accountRepository := inmemory.NewAccountRepository(accountInMemoryDB)
	rideRepository := inmemory.NewRideRepositoryInMemory(make(map[uuid.UUID]entity.Ride))
	paymentService := payment.NewPaymentServiceInMemory()
	rideFinishedHandler := event_handlers.NewRideFinishedEventHandler(rideRepository, accountRepository, paymentService)

	consumer, err := event_kafka_adapter.NewEventHandlerKafkaAdapter(kafkaAddrs, topics, map[string]event.EventHandler{domainEvent.RideFinishedEventName: rideFinishedHandler})

	if err != nil {
		panic(err)
	}

	// consumer.CreateTopics(topics)

	consumer.Listen()

	select {}
}
