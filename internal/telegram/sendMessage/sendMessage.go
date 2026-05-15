package sendmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/telegram"
)

var _ telegram.TelegramNotifier = (*TelegramNotify)(nil)

type TelegramNotify struct {
	token  string
	chatID string
	client *http.Client
}

func NewTelegramNotifier(token, chatID string) *TelegramNotify {
	return &TelegramNotify{
		token:  token,
		chatID: chatID,
		client: &http.Client{},
	}
}

func (t *TelegramNotify) SendToMessageTelegram(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.token)

	body := map[string]string{
		"chat_id": t.chatID,
		"text":    message,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal telegram message: %w", err)
	}

	resp, err := t.client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("send telegram message: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram api returned status: %s", resp.Status)
	}

	return nil
}
