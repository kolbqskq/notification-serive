package notification_service

import (
	"context"
	"time"

	"github.com/kolbqskq/notification-service/notification-worker/internal/core/domain"
)

func (s *NotificationService) HandleNotification(ctx context.Context, n *domain.NotificationRecord) error {
	if err := s.notificationRepository.Create(ctx, n); err != nil {
		return err
	}
	if err := s.sendWithRetry(ctx, n.UserID.String(), n.Message); err != nil {
		n.MarkFailed(err.Error())
		s.notificationRepository.UpdateStatus(ctx, n)
		return err
	}
	n.MarkSend()
	if err := s.notificationRepository.UpdateStatus(ctx, n); err != nil {
		return err
	}
	s.logger.Info().Str("user_id", n.UserID.String()).Str("notification_id", n.ID.String()).Msg("notification_sent")
	return nil
}

func (s *NotificationService) sendWithRetry(ctx context.Context, userID, message string) error {
	var err error
	for range 3 {
		err = s.telegramSender.Send(ctx, userID, message)
		if err == nil {
			return nil
		}
		time.Sleep(time.Second * 3)
	}
	return nil
}
