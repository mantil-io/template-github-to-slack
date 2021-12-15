package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	SlackWebhookEnv = "SLACK_WEBHOOK"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Default(ctx context.Context, e *Event) error {
	slackWebhook, ok := os.LookupEnv(SlackWebhookEnv)
	if !ok {
		return fmt.Errorf("slack webhook not set")
	}
	return postToSlack(slackWebhook, e)
}

func postToSlack(url string, e *Event) error {
	msg := struct {
		Text string `json:"text"`
	}{
		fmt.Sprintf(`There is new Github star for %s which now has %d stars!
New :star: was made by <%s|%s>.`, e.Repository.Name, e.Repository.Stars, e.Sender.URL, e.Sender.Login),
	}

	buf, _ := json.Marshal(msg)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(buf))
	if err != nil {
		return fmt.Errorf("new request failed: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response failed: %s", err)
	}
	if string(body) != "ok" {
		return fmt.Errorf("non-ok response: %s", string(body))
	}
	return nil
}

type Event struct {
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

type Repository struct {
	Name  string `json:"name"`
	Stars int    `json:"stargazers_count"`
}

type Sender struct {
	Login string `json:"login"`
	URL   string `json:"html_url"`
}
