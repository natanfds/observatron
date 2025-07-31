package main

import (
	"fmt"
	"net/http"

	"github.com/natanfds/observatron/internal/user_vtt"
	"github.com/natanfds/observatron/services"
	"github.com/natanfds/observatron/utils"
)

func startAPI() error {
	err := utils.ENV.Load()
	if err != nil {
		return err
	}

	db, err := services.NewDatabase([]interface{}{user_vtt.UserVttModel{}})
	if err != nil {
		return err
	}

	//declarações de services
	userVttService := user_vtt.NewUserVttService(
		user_vtt.NewUserVttRepo(db),
	)

	//declarações de handlers
	userVttHandler := user_vtt.NewHandler(*userVttService)

	httpServer := http.NewServeMux()
	server := &http.Server{
		Addr:    utils.ENV.ApiPort,
		Handler: httpServer,
	}

	httpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})

	httpServer.HandleFunc("/log/vtt/user", userVttHandler.UserHandle)

	fmt.Println("Server running on port" + utils.ENV.ApiPort)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := startAPI()
	if err != nil {
		send_err := utils.SendToWebhook("Error at startup **" + utils.ENV.AppName + "**\n" + err.Error())
		fmt.Println(send_err)
		if send_err != nil {
			panic(send_err)
		}
	}
}
