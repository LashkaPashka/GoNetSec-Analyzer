package config

import (
	"github.com/joho/godotenv"
)

type PortConfig interface {
	GetUDPPort() int
}

type BufferSizeConfig interface {
	GetBufferSize() int
}

type TelegramTokenConfig interface {
	GetToken() string
}

type ChatIDConfig interface {
	GetChatID() string
}

type WorkersConfig interface {
	GetWorkers() int
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
