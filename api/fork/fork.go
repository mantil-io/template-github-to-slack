package fork

import (
	"context"
	"fmt"
	"os"

	"github.com/mantil-io/template-github-to-slack/slack"
)

const (
	SlackWebhookEnv = "SLACK_WEBHOOK"
)

type Fork struct{}

func New() *Fork {
	return &Fork{}
}

func (f *Fork) Default(ctx context.Context, e *Event) error {
	slackWebhook, ok := os.LookupEnv(SlackWebhookEnv)
	if !ok {
		return fmt.Errorf("slack webhook not set")
	}
	return postToSlack(slackWebhook, e)
}

func postToSlack(url string, e *Event) error {
	text := fmt.Sprintf("Repository %s was forked by <%s|%s>!", e.Repository.Name, e.Sender.URL, e.Sender.Login)
	return slack.Post(url, text)
}

type Event struct {
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

type Repository struct {
	Name string `json:"name"`
}

type Sender struct {
	Login string `json:"login"`
	URL   string `json:"html_url"`
}
