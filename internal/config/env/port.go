package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
)

var _ config.PortConfig = (*UDPPort)(nil)

const (
	udpporEnvName = "UDPPort"
)

type UDPPort struct {
	udpport int
}

func NewUDPPort() (*UDPPort, error) {
	sudpport := os.Getenv(udpporEnvName)
	if len(sudpport) == 0 {
		return nil, errors.New("udp port not found")
	}

	udpport, err := strconv.Atoi(sudpport)
	if err != nil {
		return nil, errors.New("invalid parse udport in integer")
	}

	return &UDPPort{
		udpport: udpport,
	}, nil
}

func (u UDPPort) GetUDPPort() int {
	return u.udpport
}
