package main

import (
	"fmt"
	"net/http"

	"github.com/natanfds/observatron/app/services"
	"github.com/natanfds/observatron/utils"
)

func startMain() error {
	err := utils.ENV.Load()
	if err != nil {
		return err
	}

	err = services.StartDatabase()
	if err != nil {
		return err
	}

	httpServer := http.NewServeMux()
	server := &http.Server{
		Addr:    utils.ENV.ApiPort,
		Handler: httpServer,
	}

	httpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})

	fmt.Println("Server running on port" + utils.ENV.ApiPort)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := startMain()
	if err != nil {
		send_err := utils.SendToWebhook("Error at startup **" + utils.ENV.AppName + "**\n" + err.Error())
		fmt.Println(send_err)
		if send_err != nil {
			panic(send_err)
		}
	}
}
