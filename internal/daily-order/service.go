package daily_order

import (
	"daily-random-order/internal"
	"daily-random-order/internal/platform/slack"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	ErrorGetTeam = errors.New("error getting team")
	ErrorSendMsg = errors.New("error sending the message")
)

type Service struct {
	slackClient slack.Slack
	repository  internal.DailyRepository
}

func (s Service) FindAllCoWorker() error {
	coWorkers, err := s.repository.FindAllCoWorkers()
	if err != nil {
		return fmt.Errorf("%v - %s", ErrorGetTeam, err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(coWorkers), func(i, j int) { coWorkers[i], coWorkers[j] = coWorkers[j], coWorkers[i] })

	message := strings.Join(coWorkers, "\n")
	err = s.slackClient.SendSlackNotification(message)
	if err != nil {
		return ErrorSendMsg
	}

	return nil
}

func NewService(repository internal.DailyRepository, slackClient slack.Slack) Service {
	return Service{repository: repository, slackClient: slackClient}
}
