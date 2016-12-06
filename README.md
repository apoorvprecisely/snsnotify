# snsnotify

A golang implementation to send notification to slack/mail via AWS Simple notification Service

Example
* Slack
    ###### Requirements
    * Target ARN for slack SNS
    * Webhook Url from your slack and sns integration
    * Name of your slack channel
```
sn := snsnotify.NewSlackNotifier(awsRegion,slackarn,slackchName,webhook)
err := sn.NotifySlack("Alert", message)

```

* Mail

```
mn := snsnotify.NewMailNotifier(awsRegion,mailarn)
err := mn.NotifyMail(subject, message)
```