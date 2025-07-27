package main

import (
	"fmt"
	"net/http"
)

func main() {
	httpServer := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: httpServer,
	}

	httpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})

	fmt.Println("Server running on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
