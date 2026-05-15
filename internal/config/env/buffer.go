package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
)

var _ config.BufferSizeConfig = (*BufferSize)(nil)

const (
	bufferSizeEnvName = "BufferSize"
)

type BufferSize struct {
	bufferSize int
}

func NewBufferSize() (*BufferSize, error) {
	sbufferSize := os.Getenv(bufferSizeEnvName)
	if len(sbufferSize) == 0 {
		return nil, errors.New("udp port not found")
	}

	udpport, err := strconv.Atoi(sbufferSize)
	if err != nil {
		return nil, errors.New("invalid parse udport in integer")
	}

	return &BufferSize{
		bufferSize: udpport,
	}, nil
}

func (b BufferSize) GetBufferSize() int {
	return b.bufferSize
}
