package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type BaseInterceptor struct{}

func (hw *BaseInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Unary Method: ", info.FullMethod)
	resp, err := handler(ctx, req)
	return resp, err
}

// The stream interceptor function.
func (hw *BaseInterceptor) Stream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("Stream Method: ", info.FullMethod)
	err := handler(srv, ss)
	return err
}
