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

	err, args := configDeadLetterStrategy(ch)
	if err != nil {
		return nil, nil, amqp.Queue{}, err
	}

	err, q := configQueueStrategy(ch, queue, args)
	if err != nil {
		return nil, nil, amqp.Queue{}, err
	}

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

func configQueueStrategy(ch *amqp.Channel, queue string, args amqp.Table) (error, amqp.Queue) {

	err := ch.ExchangeDeclare("ExchangeStockProduct", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Printf("Error ExchangeStockProduct Declare RabbitMQ: %s", err.Error())
		return err, amqp.Queue{}
	}
	q, err := ch.QueueDeclare(
		queue, //name string,
		true,  // durable bool,
		false, // autodelete
		false, // exclusive
		false, // nowait
		args)  // args
	if err != nil {
		log.Printf("Error Queue Declare  RabbitMQ: %s", err.Error())
		return err, amqp.Queue{}
	}

	err = ch.QueueBind(
		q.Name,                 //name string,
		"KeyStockProduct",      //key string,
		"ExchangeStockProduct", //exchange string
		false,                  //noWait bool,
		args)                   //args amqp.Table
	if err != nil {
		log.Printf("Error Queue Bind StockProduct Declare RabbitMQ: %s", err.Error())
		return err, amqp.Queue{}
	}

	return nil, q

}

func configDeadLetterStrategy(ch *amqp.Channel) (error, amqp.Table) {

	err := ch.ExchangeDeclare("DeadLetterExchangeStockProduct", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Printf("Error Exchange Declare RabbitMQ: %s", err.Error())
		return err, nil
	}

	_, err = ch.QueueDeclare("DeadLetterQueueStockProduct", true, false, false, false, nil)
	if err != nil {
		log.Printf("Error Queue Dead Letter StockProduct Declare RabbitMQ: %s", err.Error())
		return err, nil
	}

	err = ch.QueueBind("DeadLetterQueueStockProduct", "DeadLetterKeyStockProduct",
		"DeadLetterExchangeStockProduct", false, nil)
	if err != nil {
		log.Printf("Error Queue Bind Dead Letter StockProduct Declare RabbitMQ: %s", err.Error())
		return err, nil
	}
	args := make(amqp.Table)

	args["x-dead-letter-exchange"] = "DeadLetterExchangeStockProduct"
	args["x-dead-letter-routing-key"] = "DeadLetterKeyStockProduct"

	return nil, args
}
