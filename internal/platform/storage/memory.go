package storage

import (
	"bufio"
	"daily-random-order/internal"
	"errors"
	"os"
)

var ErrorFileDoesntExist = errors.New("file is empty")

type MemoryRepository struct {
	path string
}

func NewMemoryRepository(path string) internal.DailyRepository {
	return &MemoryRepository{path: path}
}

func (m *MemoryRepository) FindAllCoWorkers() ([]string, error) {
	var coWorkers = []string{}
	file, err := os.OpenFile(m.path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			coWorkers = append(coWorkers, text)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(coWorkers) == 0 {
		return nil, ErrorFileDoesntExist
	}

	return coWorkers, nil
}
