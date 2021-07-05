package infrastructure

import (
	"context"
	"moises-ba/ms-criptcoin-vote/config"
	"moises-ba/ms-criptcoin-vote/log"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
)

func NewKafkaProducer() TopicProducerIf {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(config.GetKafkaBrokerURL(), ","),
		Topic:   config.GetKafkaBrokerURL(),
		Logger:  log.Logger(),
	})

	return &kafkaProducer{
		writer: w,
	}
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func (p *kafkaProducer) WriteMessage(message, topic string) error {

	err := p.writer.WriteMessages(context.TODO(), kafka.Message{
		//Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		Value: []byte(message),
	})

	if err != nil {
		log.Logger().Error("Nao foi possivel enviar a mensagem para a fila.", err)
		return err
	}

	return nil
}

////////

type kafkaConsumer struct {
}

func (c *kafkaConsumer) Consume(topic string) (<-chan string, error) {

	newUUID, err := uuid.NewV4()
	if err != nil {
		log.Logger().Error("Nao foi possivel gerar o UUID", err)
		return nil, err
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(config.GetKafkaBrokerURL(), ","),
		Topic:   config.GetKafkaBrokerURL(),
		Logger:  log.Logger(),
		GroupID: "group_" + newUUID,
	})

	messageChan := make(chan string)

	go func() {
		defer close(messageChan)

		for {
			msg, err := r.ReadMessage(context.TODO())
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
