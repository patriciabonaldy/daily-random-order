package bootstrap

import (
	"daily-random-order/cmd/config"
	daily_order "daily-random-order/internal/daily-order"
	"daily-random-order/internal/platform/slack"
	"daily-random-order/internal/platform/storage"
)

// Run application
func Run() error {
	repository := storage.NewMemoryRepository(config.ENV.DataPath)
	service := daily_order.NewService(repository, slack.New())
	return service.FindAllCoWorker()
}
