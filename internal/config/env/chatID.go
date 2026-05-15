package env

import (
	"errors"
	"os"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
)

var _ config.ChatIDConfig = (*ChatID)(nil)

const (
	chatIDEnvName = "ChatID"
)

type ChatID struct {
	chatID string
}

func NewChatID() (*ChatID, error) {
	chatID := os.Getenv(chatIDEnvName)
	if len(chatID) == 0 {
		return nil, errors.New("udp port not found")
	}

	return &ChatID{
		chatID: chatID,
	}, nil
}

func (c ChatID) GetChatID() string {
	return c.chatID
}
