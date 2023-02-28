package publish

import "github.com/streadway/amqp"

type RabbitMQApp struct {
	configRabbitMQService IConfigRabbitMQService
}

func NewRabbitMQApp(configRabbitMQService IConfigRabbitMQService) *RabbitMQApp {

	return &RabbitMQApp{
		configRabbitMQService: configRabbitMQService,
	}

}

func (app *RabbitMQApp) ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error) {
	return app.configRabbitMQService.ConfigRabbitMQ(queue)
}

func (app *RabbitMQApp) PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string) error {
	return app.configRabbitMQService.PublishMessage(conn, channel, queue, message)
}
