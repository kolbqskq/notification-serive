package transport_grpc

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetStatus(ctx context.Context, req *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	id, err := uuid.Parse(req.NotificationId)
	if err != nil {
		return nil, toGRPCError(err)
	}

	record, err := s.notificationService.GetStatus(ctx, id)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := &pb.GetStatusResponse{
		Status: string(record.Status),
	}
	if record.SentAt != nil {
		res.SentAt = timestamppb.New(*record.SentAt)
	}
	if record.ErrorMessage != nil {
		res.ErrorMessage = *record.ErrorMessage
	}
	return res, nil
}
