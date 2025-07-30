package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/natanfds/observatron/app/dtos"
	"github.com/natanfds/observatron/app/interfaces"
)

type DiscordWebhookSender struct {
	Url string
}

/*
Função para enviar mensagens para o webhook do Discord
*/
func (d *DiscordWebhookSender) Send(message string) error {
	body := &dtos.DiscordWebhookRequest{
		Username: ENV.AppName,
		Content:  message,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		"POST",
		ENV.WebhookURL,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
Função para encaminhar mensagens para algum webhook disponível caso ocorra um erro ao subir a API
*/
func SendToWebhook(message string) error {
	var webhookSender interfaces.Webhook
	switch true {
	case strings.Contains(ENV.WebhookURL, "discord"):
		webhookSender = &DiscordWebhookSender{
			Url: ENV.WebhookURL,
		}
	default:
		return errors.New("error on webhook instance")
	}
	return webhookSender.Send(message)
}
