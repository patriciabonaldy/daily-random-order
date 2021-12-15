package slack

import (
	"bytes"
	"daily-random-order/cmd/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type RequestBody struct {
	Text string `json:"text"`
}

type Client struct{}

func New() Slack {
	return &Client{}
}

func (c *Client) SendSlackNotification(msg string) error {
	body, _ := json.Marshal(RequestBody{Text: msg})
	//url := fmt.Sprintf(config.ENV.WebhookUrl, config.ENV.Token)
	req, err := http.NewRequest(http.MethodPost, config.ENV.WebhookUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
