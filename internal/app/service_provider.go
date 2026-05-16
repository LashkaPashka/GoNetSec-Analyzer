package app

import (
	"log"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/config/env"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/notifier"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/notifier/alert"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/parser"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/parser/rules"
	"github.com/lashkapashka/GoNetSec_Analyzer/internal/telegram"
	sendmessage "github.com/lashkapashka/GoNetSec_Analyzer/internal/telegram/sendMessage"
)

type serviceProvider struct {
	portConfig       config.PortConfig
	bufferSizeConfig config.BufferSizeConfig
	telegramToken    config.TelegramTokenConfig
	chatID           config.ChatIDConfig
	workersNumber    config.WorkersConfig

	telegramBot telegram.TelegramNotifier
	notify      notifier.Notifier
	parser      parser.Parser
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) UDPPortConfig() config.PortConfig {
	if s.portConfig == nil {
		cfg, err := env.NewUDPPort()
		if err != nil {
			log.Fatalf("failed to get udp port config: %s", err.Error())
		}

		s.portConfig = cfg
	}

	return s.portConfig
}

func (s *serviceProvider) TelegramTokenConfig() config.TelegramTokenConfig {
	if s.telegramToken == nil {
		cfg, err := env.NewTelegramToken()
		if err != nil {
			log.Fatalf("failed to get telegram token config: %s", err.Error())
		}

		s.telegramToken = cfg
	}

	return s.telegramToken
}

func (s *serviceProvider) ChatIDConfig() config.ChatIDConfig {
	if s.chatID == nil {
		cfg, err := env.NewChatID()
		if err != nil {
			log.Fatalf("failed to get chatID config: %s", err.Error())
		}

		s.chatID = cfg
	}

	return s.chatID
}

func (s *serviceProvider) BufferSizeConfig() config.BufferSizeConfig {
	if s.bufferSizeConfig == nil {
		cfg, err := env.NewBufferSize()
		if err != nil {
			log.Fatalf("failed to get buffer size config: %s", err.Error())
		}

		s.bufferSizeConfig = cfg
	}

	return s.bufferSizeConfig
}

func (s *serviceProvider) WorkersConfig() config.WorkersConfig {
	if s.workersNumber == nil {
		cfg, err := env.NewWorkers()
		if err != nil {
			log.Fatalf("failed to get workersNumber config: %s", err.Error())
		}

		s.workersNumber = cfg
	}

	return s.workersNumber
}

func (s *serviceProvider) TelegramBot() telegram.TelegramNotifier {
	if s.telegramBot == nil {
		s.telegramBot = sendmessage.NewTelegramNotifier(s.TelegramTokenConfig().GetToken(), s.ChatIDConfig().GetChatID())
	}

	return s.telegramBot
}

func (s *serviceProvider) Notifier() notifier.Notifier {
	if s.notify == nil {
		s.notify = alert.NewAlertNotify(s.TelegramBot())
	}

	return s.notify
}

func (s *serviceProvider) Parser() parser.Parser {
	if s.parser == nil {
		s.parser = rules.NewAnalyzer()
	}

	return s.parser
}
