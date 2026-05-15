package telegram

type TelegramNotifier interface {
	SendToMessageTelegram(message string) error
}