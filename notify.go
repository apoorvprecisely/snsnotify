package snsnotify

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	ew "github.com/hashicorp/errwrap"
)

type SlackMessage struct {
	Content     string `json:"content"`
	Channel     string `json:"channel"`
	Webhookpath string `json:"webhookpath"`
}

type SlackNotifier struct {
	SlackArn     string
	SlackCh      string
	SlackWebhook string
	AwsRegion    string
}

type MailNotifier struct {
	MailArn   string
	AwsRegion string
}

func NewSlackNotifier(awsRegion string, slackArn string, slackCh string, slackWebhook string) *SlackNotifier {
	sn := &SlackNotifier{
		SlackArn:     slackArn,
		SlackCh:      slackCh,
		SlackWebhook: slackWebhook,
		AwsRegion:    awsRegion,
	}
	return sn
}

func NewMailNotifier(awsRegion string, mailArn string) *MailNotifier {
	mn := &MailNotifier{
		AwsRegion: awsRegion,
		MailArn:   mailArn,
	}
	return mn
}

func (sn *SlackNotifier) NotifySlack(subject string, message string) error {
	svc := sns.New(session.New(aws.NewConfig().WithRegion(sn.AwsRegion)))
	return sn.call(svc, subject, message)
}

func (sn *SlackNotifier) call(svc *sns.SNS, subject string, message string) error {
	smsg := &SlackMessage{Channel: sn.SlackCh, Webhookpath: sn.SlackWebhook, Content: message}
	b, err := json.Marshal(smsg)

	if err != nil {
		return ew.Wrapf("Error while generating message : {{err}}",
			err)
	}

	params := &sns.PublishInput{
		Message:  aws.String(string(b)),
		TopicArn: aws.String(sn.SlackArn),
		Subject:  aws.String(subject),
	}

	_, err = svc.Publish(params)

	if err != nil {
		return ew.Wrapf("Error while notifying slack : {{err}}",
			err)
	}
	return nil
}

func (sn *SlackNotifier) NotifySlackWKey(accessKey string, secretKey string, subject string, message string) error {
	svc := sns.New(session.New(
		aws.NewConfig().WithRegion(sn.AwsRegion).
			WithCredentials(credentials.NewStaticCredentials(accessKey, secretKey, "nil"))))
	return sn.call(svc, subject, message)
}

func (mn *MailNotifier) NotifyMail(subject string, message string) error {
	svc := sns.New(session.New(aws.NewConfig().WithRegion(mn.AwsRegion)))
	params := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(mn.MailArn),
		Subject:  aws.String(subject),
	}

	_, err := svc.Publish(params)

	if err != nil {
		return ew.Wrapf("Error while notifying by mail : {{err}}",
			err)
	}
	return nil
}
