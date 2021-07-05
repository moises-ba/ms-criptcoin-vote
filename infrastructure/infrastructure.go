package infrastructure

type TopicProducerIf interface {
	WriteMessage(message, topic string) error
}

type TopicConsumerIf interface {
	Consume(topic string) (<-chan string, error)
}
