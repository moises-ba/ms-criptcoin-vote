package infrastructure

type TopicProducerIf interface {
	WriteMessage(message, topic string) error
}

type TopicConsumerIf interface {
	Consume(topic string) (<-chan string, error)
	Stop() error
}

type CoinVoteTopicMessage struct {
	CoinId                string
	TotalApprovedVotes    int
	TotalDisapprovedVotes int
}
