package env

import (
	"errors"
	"os"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
)

var _ config.TelegramTokenConfig = (*Token)(nil)

const (
	telegramTokenEnvName = "TelegramToken"
)

type Token struct {
	telegramToken string
}

func NewTelegramToken() (*Token, error) {
	telegramToken := os.Getenv(telegramTokenEnvName)
	if len(telegramToken) == 0 {
		return nil, errors.New("udp port not found")
	}

	return &Token{
		telegramToken: telegramToken,
	}, nil
}

func (t Token) GetToken() string {
	return t.telegramToken
}
