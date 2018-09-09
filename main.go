package main

import (
	"os"

	eventhandler "github.com/mtyurt/supervisor-event-handler"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

func sendSlackMessage(process string, prevStatus string, changeEvent string) error {
	channel := os.Getenv("SLACK_CHANNEL")
	token := os.Getenv("SLACK_TOKEN")

	api := slack.New(token)
	text := "Supervisor has a message for you!"
	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			slack.Attachment{
				MarkdownIn: []string{"pretext"},
				Title:      text,
				Fields: []slack.AttachmentField{
					slack.AttachmentField{
						Title: "Process",
						Value: process,
						Short: true,
					},
					slack.AttachmentField{
						Title: "Previous status",
						Value: prevStatus,
						Short: true,
					},
					slack.AttachmentField{
						Title: "Change event",
						Value: changeEvent,
						Short: true,
					},
				},
			},
		},
	}
	_, _, err := api.PostMessage(channel, "", params)
	if err != nil {
		return errors.Wrap(err, "Could not send the message to slack")
	}
	return nil

}
func main() {

	handler := eventhandler.EventHandler{}

	handler.HandleEvent("PROCESS_STATE", func(header eventhandler.HeaderTokens, payload map[string]string) {

		err := sendSlackMessage(payload["processname"], payload["from_state"], header.EventName)
		if err != nil {
			panic(err)
		}

	})
	handler.Start()
}
