package stockproductservice

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marc/go-clean-example/adapter/rabbitMQ/publish"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

func SendMessageStockProduct(response http.ResponseWriter, request *http.Request) {

	configRabbitMQServiceApp := publish.NewRabbitMQApp(&publish.ConfigRabbitMQService{})

	conn, channel, queue, err := configRabbitMQServiceApp.ConfigRabbitMQ("queueStockProduct")

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}
	text := string(bytes)
	message := fmt.Sprint(text)

	err = configRabbitMQServiceApp.PublishMessage(conn, channel, queue, message)
	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusOK, nil)

}
