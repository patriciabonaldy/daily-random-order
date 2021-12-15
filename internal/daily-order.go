package internal

// DailyRepository defines the expected behaviour from a storage.
type DailyRepository interface {
	FindAllCoWorkers() ([]string, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=DailyRepository
