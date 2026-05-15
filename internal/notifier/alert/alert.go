package alert

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/notifier"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/telegram"
	"golang.org/x/sync/errgroup"
)

var _ notifier.Notifier = (*AlertNotify)(nil)

type AlertNotify struct {
	telegramNotifier telegram.TelegramNotifier
}

func NewAlertNotify(telegramNotifier telegram.TelegramNotifier) AlertNotify {
	return AlertNotify{
		telegramNotifier: telegramNotifier,
	}
}

func (a AlertNotify) AlertNotifier(ctx context.Context, alertChan <-chan string, workers int) error {
	if a.telegramNotifier == nil {
		return fmt.Errorf("telegram notifier is nil")
	}
	
	g, gctx := errgroup.WithContext(ctx)

	if workers <= 0 {
		workers = 1
	}

	for i := 0; i < workers; i++ {
		g.Go(func() error {
			for {
				select {
				case <-gctx.Done():
					return gctx.Err()
				case alert, ok := <-alertChan:
					if !ok {
						return nil
					}
					fmt.Println(strings.Repeat("-", 50))
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
					fmt.Println(alert)
					fmt.Println(strings.Repeat("-", 50))

					if err := a.telegramNotifier.SendToMessageTelegram(alert); err != nil {
						log.Printf("failed to send alert to telegram: %v", err)
						continue
					}

				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("alert notifier stopped: %w", err)
	}

	return nil
}
