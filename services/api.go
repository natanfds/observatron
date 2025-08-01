package services

import (
	"fmt"
	"net/http"

	"github.com/natanfds/observatron/internal/task_queue"
	"github.com/natanfds/observatron/internal/user_vtt"
	"github.com/natanfds/observatron/utils"
)

func StartAPI() error {
	err := utils.ENV.Load()
	if err != nil {
		return err
	}

	db, err := NewDatabase([]interface{}{user_vtt.UserVttModel{}})
	if err != nil {
		return err
	}

	dbTaskQueue := task_queue.NewTaskQueue(1000)
	taskQueues := []task_queue.TaskQueue{
		*dbTaskQueue,
	}

	for _, queue := range taskQueues {
		queue.Start()
	}

	//declarações de services
	userVttService := user_vtt.NewUserVttService(
		user_vtt.NewUserVttRepo(db),
		dbTaskQueue,
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
