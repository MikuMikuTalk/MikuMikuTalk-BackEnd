package zrpc_interceptor

import (
	"context"
	"fmt"
	"im_server/common/contexts"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientInfoInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 安全地提取 clientIP 和 userID
	clientIPVal := ctx.Value(contexts.ContextKeyClientIP)
	userIDVal := ctx.Value(contexts.ContextKeyUserID)

	var clientIP, userID string

	if val, ok := clientIPVal.(string); ok {
		clientIP = val // 直接使用字符串
	} else {
		return fmt.Errorf("invalid type for clientIP: %T", clientIPVal)
	}

	if val, ok := userIDVal.(uint); ok {
		userID = fmt.Sprintf("%d", val) // 将 uint 转换为字符串
	} else {
		return fmt.Errorf("invalid type for userID: %T", userIDVal)
	}

	// 创建元数据
	md := metadata.New(map[string]string{
		"clientIP": clientIP,
		"userID":   userID,
	})

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	//请求之后
	return err
}
