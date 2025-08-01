package main

import (
	"fmt"

	"github.com/natanfds/observatron/services"
	"github.com/natanfds/observatron/utils"
)

func main() {
	err := services.StartAPI()
	if err != nil {
		send_err := utils.SendToWebhook("Error at startup **" + utils.ENV.AppName + "**\n" + err.Error())
		fmt.Println(send_err)
		if send_err != nil {
			panic(send_err)
		}
	}
}
