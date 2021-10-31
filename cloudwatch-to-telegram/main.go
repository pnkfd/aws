package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var chat_id = os.Getenv("chat_id") //variáveis
var bot_id = os.Getenv("bot_id")

func HandleRequest(logsEvent events.CloudwatchLogsEvent) (string, error) {

	data, _ := logsEvent.AWSLogs.Parse()
	for _, logEvent := range data.LogEvents {
		resp, err := http.Get("https://api.telegram.org/" + bot_id + "/sendMessage?chat_id=" + chat_id + "&text=" + logEvent.Message)

		if err != nil { //se der erro, "loga"
			log.Print(err)
		}
		defer resp.Body.Close() //fecha a requisição
	}
	return "sucesso", nil

}

func main() {

	lambda.Start(HandleRequest)
}
