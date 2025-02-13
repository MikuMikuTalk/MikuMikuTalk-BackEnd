package zrpc_interceptor

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	clientIP := metadata.ValueFromIncomingContext(ctx, "clientIP")
	userID := metadata.ValueFromIncomingContext(ctx, "userID")
	logx.Debug("ServerUnaryInterceptor: ", clientIP, "  ", userID)
	if len(clientIP) > 0 {
		ctx = context.WithValue(ctx, "clientIP", clientIP[0])
	}
	if len(userID) > 0 {
		ctx = context.WithValue(ctx, "userID", userID[0])
	}
	return handler(ctx, req)
}
