package app

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
	"golang.org/x/sync/errgroup"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	workers := 5

	logChan := make(chan string, 1000)
	alertChan := make(chan string, 100)

	fmt.Println("Запуск GoNetSec Analyzer")

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.serviceProvider.Parser().LogAnalyzer(logChan, alertChan)
		return nil
	})

	g.Go(func() error {
		return a.serviceProvider.Notifier().AlertNotifier(gctx, alertChan, workers)
	})

	g.Go(func() error {
		defer close(logChan)
		return a.startUDPServer(gctx, logChan)
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("app stopped with error: %w", err)
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(".env"); err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) startUDPServer(ctx context.Context, logChan chan<- string) error {
	addr := net.UDPAddr{
		Port: a.serviceProvider.UDPPortConfig().GetUDPPort(),
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return fmt.Errorf("listen udp: %w", err)
	}
	defer conn.Close()

	fmt.Printf("Сервер слушает UDP порт %d...\n", a.serviceProvider.UDPPortConfig().GetUDPPort())

	buffer := make([]byte, 4096)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			if err := conn.SetReadDeadline(time.Now().Add(1 * time.Second)); err != nil {
				return fmt.Errorf("set read deadline: %w", err)
			}

			n, remoteAddr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				netErr, ok := err.(net.Error)
				if ok && netErr.Timeout() {
					continue
				}

				return fmt.Errorf("read udp: %w", err)
			}

			message := string(buffer[:n])

			select {
			case logChan <- fmt.Sprintf("[%s] %s", remoteAddr.IP.String(), message):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}
