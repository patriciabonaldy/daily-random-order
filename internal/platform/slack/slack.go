package slack

type Slack interface {
	SendSlackNotification(msg string) error
}

//go:generate mockery --case=snake --outpkg=slackmocks --output=slackmocks --name=Slack
