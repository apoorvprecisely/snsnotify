# snsnotify

A golang implementation to send notification to slack/mail via AWS Simple notification Service

Example
* Slack
    ###### Requirements
    * Target ARN for slack SNS
    * Webhook Url from your slack and sns integration
    * Name of your slack channel
```
cred := snsnotify.SlackCredential{
		SlackArn:     slackarn,
		SlackWebhook: webhook,
		SlackCh:      slackchName,
	}

err := snsnotify.NotifySlack(awsRegion, cred, subject, message)

```

* Mail

```
err := snsnotify.NotifyMail(awsRegion, mailArn, subject, message)
```