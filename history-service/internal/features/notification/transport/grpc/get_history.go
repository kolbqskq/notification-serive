package transport_grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetHistory(ctx context.Context, req *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, s.toGRPCError(err)
	}

	records, total, err := s.notificationService.GetHistory(ctx, userID, req.Limit, req.Offset)
	if err != nil {
		return nil, s.toGRPCError(err)
	}

	return &pb.GetHistoryResponse{
		Records: toProtoRecords(records),
		Total:   total,
	}, nil
}

func toProtoRecords(records []*domain.NotificationRecord) []*pb.NotificationRecord {
	res := make([]*pb.NotificationRecord, 0, len(records))
	for _, r := range records {
		res = append(res, toProtoRecord(r))
	}
	return res
}

func toProtoRecord(r *domain.NotificationRecord) *pb.NotificationRecord {
	record := &pb.NotificationRecord{
		Id:            r.ID.String(),
		UserId:        r.UserID.String(),
		Type:          r.Type,
		Message:       r.Message,
		SourceService: r.SourceService,
		Status:        toProtoStatus(r.Status),
		CreatedAt:     timestamppb.New(r.CreatedAt),
	}
	if r.SentAt != nil {
		record.SentAt = timestamppb.New(*r.SentAt)
	}
	if r.ErrorMessage != nil {
		record.ErrorMessage = *r.ErrorMessage
	}

	return record
}

func toProtoStatus(s domain.NotificationStatus) pb.NotificationStatus {
	switch s {
	case domain.StatusPending:
		return pb.NotificationStatus_STATUS_PENDING
	case domain.StatusSent:
		return pb.NotificationStatus_STATUS_SENT
	case domain.StatusFailed:
		return pb.NotificationStatus_STATUS_FAILED
	default:
		return pb.NotificationStatus_STATUS_UNSPECIFIED
	}
}
