package notifier

import "context"

type Notifier interface {
	AlertNotifier(ctx context.Context, alertChan <-chan string, workers int) error
}
