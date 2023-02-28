package publish

import (
	"log"
	"time"

	"github.com/avast/retry-go"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type IConfigRabbitMQService interface {
	ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error)
	PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string) error
}

const (
	MaxRetriesPublishMessage = 3
	RetryDelayPublishMessage = 1 * time.Second
)

type ConfigRabbitMQService struct{}

func (config *ConfigRabbitMQService) ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error) {

	connectionAMQPURL := viper.GetString("rabbitMQ.connectionAMQP")
	conn, err := amqp.Dial(connectionAMQPURL)
	if err != nil {
		log.Printf("Error Connection RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error Channel RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		queue, //name string,
		true,  // durable bool,
		false, // autodelete
		false, // exclusive
		false, // nowait
		nil)   // args
	if err != nil {
		log.Printf("Error Queue Declare  RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}

	ch.QueueBind(
		q.Name,       //name string,
		"",           //key string,
		"amq.fanout", //exchange string
		false,        //noWait bool,
		nil)          //args amqp.Table

	return conn, ch, q, nil

}

func (config *ConfigRabbitMQService) PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string) error {

	msg := amqp.Publishing{
		Headers:         map[string]interface{}{},
		ContentType:     "",
		ContentEncoding: "",
		DeliveryMode:    2,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "messageStockProduct",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "go-clean-example",
		Body:            []byte(message),
	}

	errRetryFail := retry.Do(
		func() error {

			// _, erroTest := strconv.ParseUint("ABC", 10, 64)
			// if erroTest != nil {
			// return erroTest
			// }

			err1 := channel.Publish("", queue.Name, false, false, msg)
			if err1 != nil {
				return err1
			}

			return nil

		}, retry.Attempts(MaxRetriesPublishMessage),
		retry.Delay(RetryDelayPublishMessage),
	)

	if errRetryFail != nil {
		log.Printf("Exceeded number of message publishing attempts")
		return errRetryFail
	}

	return nil

}
