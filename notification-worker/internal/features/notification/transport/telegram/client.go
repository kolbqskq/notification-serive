package transport_telegram

import (
	"context"

	"github.com/go-telegram/bot"
)

type TelegramClient struct {
	bot    *bot.Bot
	chatID int64
}

type TelegramClientDeps struct {
	Token  string
	ChatID int64
}

func NewTelegramClient(deps TelegramClientDeps) (*TelegramClient, error) {
	bot, err := bot.New(deps.Token)
	if err != nil {
		return nil, err
	}

	return &TelegramClient{
		bot:    bot,
		chatID: deps.ChatID,
	}, nil
}

func (t *TelegramClient) Send(ctx context.Context, userID string, message string) error {
	_, err := t.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: t.chatID,
		Text:   message,
	})

	return err
}
