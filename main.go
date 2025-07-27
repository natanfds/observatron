package main

import (
	"fmt"
	"net/http"

	"github.com/natanfds/observatron/utils"
)

func main() {
	err := utils.ENV.Load()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

}
