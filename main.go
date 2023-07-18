package main

import (
	"os/exec"
	"time"

	"github.com/nlopes/slack"
)

func checkGPUStatus() error {
	cmd := exec.Command("nvidia-smi")
	err := cmd.Run()
	if err != nil {
		return err
	}
	// Parse the output of nvidia-smi and check for error state
	// Return an error if error state detected, nil otherwise
	return nil
}

func sendSlackNotification(webhookURL, message string) error {
	api := slack.New(webhookURL)
	attachment := slack.Attachment{
		Text: message,
	}
	_, _, err := api.PostMessage("", slack.MsgOptionAttachments(attachment))
	if err != nil {
		return err
	}
	// Optionally, log the channelID and timestamp for reference
	return nil
}

func main() {
	webhookURL := "https://hooks.slack.com/services/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"
	checkInterval := 5 * time.Minute // Adjust the interval as per your needs

	for {
		err := checkGPUStatus()
		if err != nil {
			errorMessage := "Error state detected on the GPU: " + err.Error()
			err := sendSlackNotification(webhookURL, errorMessage)
			if err != nil {
				// Handle the error if Slack notification fails
			}
		}

		time.Sleep(checkInterval)
	}
}
