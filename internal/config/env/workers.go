package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
)

var _ config.WorkersConfig = (*Workers)(nil)

const (
	workersNumberEnvName = "WorkersNumber"
)

type Workers struct {
	workers int
}

func NewWorkers() (*Workers, error) {
	sworkers := os.Getenv(workersNumberEnvName)
	if len(sworkers) == 0 {
		return nil, errors.New("udp port not found")
	}

	workers, err := strconv.Atoi(sworkers)
	if err != nil {
		return nil, errors.New("invalid parse udport in integer")
	}

	return &Workers{
		workers: workers,
	}, nil
}

func (w Workers) GetWorkers() int {
	return w.workers
}
