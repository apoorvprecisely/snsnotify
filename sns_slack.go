package snsnotify

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	ew "github.com/hashicorp/errwrap"
)

type SlackMessage struct {
	Content     string `json:"content"`
	Channel     string `json:"channel"`
	Webhookpath string `json:"webhookpath"`
}

func NotifySlack(awsRegion string, slackWebhook string, slackCh string, slackArn string, subject string, message string) error {
	svc := sns.New(session.New(aws.NewConfig().WithRegion(awsRegion)))
	smsg := &SlackMessage{Channel: slackCh, Webhookpath: slackWebhook, Content: message}
	b, err := json.Marshal(smsg)

	if err != nil {
		return ew.Wrapf("Error while generating message : {{err}}",
			err)
	}

	params := &sns.PublishInput{
		Message:  aws.String(string(b)),
		TopicArn: aws.String(slackArn),
		Subject:  aws.String(subject),
	}

	_, err = svc.Publish(params)

	if err != nil {
		return ew.Wrapf("Error while notifying slack : {{err}}",
			err)
	}
	return nil
}
