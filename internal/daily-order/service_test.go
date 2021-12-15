package daily_order_test

import (
	daily_order "daily-random-order/internal/daily-order"
	"daily-random-order/internal/platform/slack/slackmocks"
	"daily-random-order/internal/platform/storage/storagemocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_FindAllCoWorker(t *testing.T) {
	testCases := []struct {
		name                string
		MockRepositoryError error
		MockSlackError      error
		expectedError       error
	}{
		{
			name:                "FindAllCoWorker-OK",
			MockRepositoryError: nil,
			MockSlackError:      nil,
			expectedError:       nil,
		},
		{
			name:                "FindAllCoWorker-Error-in-repository",
			MockRepositoryError: errors.New("unknown error"),
			expectedError:       errors.New("error getting team - unknown error"),
		},
		{
			name:           "FindAllCoWorker-Error-in-slack-client",
			MockSlackError: errors.New("unknown error"),
			expectedError:  errors.New("error sending the message"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repositoryMock := new(storagemocks.DailyRepository)
			repositoryMock.On("FindAllCoWorkers").Return([]string{"Maria", "Juan"}, testCase.MockRepositoryError)
			slackClient := new(slackmocks.Slack)
			slackClient.On("SendSlackNotification", mock.Anything).Return(testCase.MockSlackError)
			service := daily_order.NewService(repositoryMock, slackClient)
			err := service.FindAllCoWorker()
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
