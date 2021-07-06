package security

import (
	"context"

	"moises-ba/ms-criptcoin-vote/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Metadado nao provido")
	}

	tokenValue := md["authorization_jwt_token"]
	if len(tokenValue) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "JWT nao fornecido")
	}

	//	metadata.FromOutgoingContextRaw(ctx, tokenValue[0])

	log.Logger().Println("--> unary interceptor: ")
	return handler(ctx, req)
}
