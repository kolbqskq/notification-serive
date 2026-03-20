package transport_grpc

import (
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHistoryClient(addr string) (pb.HistoryServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return pb.NewHistoryServiceClient(conn), conn, nil
}
