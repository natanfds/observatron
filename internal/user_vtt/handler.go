package user_vtt

import (
	"encoding/json"
	"net/http"

	"github.com/natanfds/observatron/dtos"
	"github.com/natanfds/observatron/utils"
)

type Handler struct {
	userVttService Service
}

func NewHandler(userVttService Service) *Handler {
	return &Handler{userVttService: userVttService}
}

func (l *Handler) UserHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	var userVttData dtos.UserVtt
	err := json.NewDecoder(r.Body).Decode(&userVttData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}
	err = utils.Validate.Struct(userVttData)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	err = l.userVttService.Create(userVttData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create user VTT"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User VTT created successfully"))

}
