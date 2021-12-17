package star

import (
	"context"
	"fmt"
	"os"

	"github.com/mantil-io/template-github-to-slack/slack"
)

const (
	SlackWebhookEnv = "SLACK_WEBHOOK"
)

type Star struct{}

func New() *Star {
	return &Star{}
}

func (s *Star) Default(ctx context.Context, e *Event) error {
	slackWebhook, ok := os.LookupEnv(SlackWebhookEnv)
	if !ok {
		return fmt.Errorf("slack webhook not set")
	}
	return postToSlack(slackWebhook, e)
}

func postToSlack(url string, e *Event) error {
	text := fmt.Sprintf(`There is new Github star for %s which now has %d stars!
New :star: was made by <%s|%s>.`, e.Repository.Name, e.Repository.Stars, e.Sender.URL, e.Sender.Login)
	return slack.Post(url, text)
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
