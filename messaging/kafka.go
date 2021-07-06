package messaging

import (
	"context"
	"encoding/json"
	"moises-ba/ms-criptcoin-vote/config"
	"moises-ba/ms-criptcoin-vote/log"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
)

func init() {
	createTopic() //cria o topico no inicio
}

func NewKafkaProducer() TopicProducerIf {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(config.GetKafkaBrokerURL(), ","),
		Topic:    config.GetVoteTopic(),
		Balancer: &kafka.LeastBytes{},
		Logger:   log.Logger(),
	})

	return &kafkaProducer{
		writer: w,
	}
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func (p *kafkaProducer) WriteMessage(objMessage interface{}, topic string) error {

	message, err := json.Marshal(objMessage)
	if err != nil {
		log.Logger().Error("Falha ao transformar struct em json.", err)
		return err
	}

	err = p.writer.WriteMessages(context.Background(), kafka.Message{
		//	Key:   []byte(strconv.Itoa(rand.Intn(100))), //melhora a key
		//Partition: 0,
		Value: message,
	})

	if err != nil {
		log.Logger().Error("Nao foi possivel enviar a mensagem para a fila.", err)
		return err
	}

	return nil
}

func NewKafkaConsumer() TopicConsumerIf {
	return &kafkaConsumer{}
}

type kafkaConsumer struct {
	reader *kafka.Reader
}

func (c *kafkaConsumer) Consume(topic string) (<-chan string, error) {

	newUUID := uuid.NewV4().String()

	c.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     strings.Split(config.GetKafkaBrokerURL(), ","),
		Topic:       config.GetVoteTopic(),
		Logger:      log.Logger(),
		GroupID:     "group_" + newUUID,
		StartOffset: kafka.LastOffset,
	})

	messageChan := make(chan string)

	go func() {
		defer close(messageChan)

		for {
			msg, err := c.reader.ReadMessage(context.TODO())
			if err != nil {
				log.Logger().Error("Nao foi possivel ler mensagem do topico ", err)
				return
			}

			log.Logger().Debug("Mensagem recebida: " + string(msg.Value))

			messageChan <- string(msg.Value)

		}

	}()

	return messageChan, nil
}

func (c *kafkaConsumer) Stop() error {
	if c.reader != nil {
		return c.reader.Close()
	}

	return nil
}

func createTopic() {

	conn, err := kafka.Dial("tcp", config.GetKafkaBrokerURL())
	if err != nil {
		panic(err.Error())
	}

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             config.GetVoteTopic(),
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

}
