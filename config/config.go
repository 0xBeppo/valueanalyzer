package config

import "os"

type TelegramConfig struct {
	TelegramBotToken string
	TelegramChatId   string
}

func GetConfig() *TelegramConfig {
	return &TelegramConfig{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChatId:   os.Getenv("TELEGRAM_CHAT_ID"),
	}
}
