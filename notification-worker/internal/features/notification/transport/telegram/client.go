package transport_telegram

import (
	"context"
	"net/http"
	"time"

	"github.com/go-telegram/bot"
	"github.com/rs/zerolog"
)

type TelegramClient struct {
	bot    *bot.Bot
	chatID int64
	logger zerolog.Logger
}

type TelegramClientDeps struct {
	Token  string
	ChatID int64
	Logger zerolog.Logger
}

func NewTelegramClient(deps TelegramClientDeps) (*TelegramClient, error) {

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	b, err := bot.New(deps.Token, bot.WithHTTPClient(5*time.Second, httpClient))
	if err != nil {
		return nil, err
	}

	return &TelegramClient{
		bot:    b,
		chatID: deps.ChatID,
		logger: deps.Logger,
	}, nil
}

func (t *TelegramClient) Send(ctx context.Context, userID string, message string) error {
	_, err := t.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: t.chatID,
		Text:   message,
	})
	if err == nil {
		t.logger.Info().Msg("notification sent")
	}
	return err
}
