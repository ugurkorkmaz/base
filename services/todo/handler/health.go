package handler

import (
	"context"
	"time"

	pb "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	pb.HealthServer
}

// TODO
func (hs *HealthServer) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

// TODO
func (hs *HealthServer) Watch(req *pb.HealthCheckRequest, stream pb.Health_WatchServer) error {
	for {
		time.Sleep(2 * time.Second)
		resp := &pb.HealthCheckResponse{
			Status: pb.HealthCheckResponse_SERVING,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
