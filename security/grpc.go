package security

import (
	"context"
	"moises-ba/ms-criptcoin-vote/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TokenValidatorIf interface {
	Validate(token string) (*UserClaims, error)
	UnaryInterceptor() func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error)
}

type jwtValidator struct {
	jwtManager *JWTManager
}

func NewJWTValidator(pJWTManager *JWTManager) TokenValidatorIf {
	return &jwtValidator{jwtManager: pJWTManager}
}

func (v *jwtValidator) Validate(token string) (*UserClaims, error) {
	return v.jwtManager.Verify(token)
}

func (v *jwtValidator) UnaryInterceptor() func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	return func(ctx context.Context,
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

		userclaim, err := v.Validate(tokenValue[0])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "JWT Inválido")
		}

		ctx = metadata.AppendToOutgoingContext(ctx, "username", userclaim.Username)

		log.Logger().Printf("--> claim: ", userclaim)
		return handler(ctx, req)
	}

}
