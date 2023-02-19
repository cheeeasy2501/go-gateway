package service

import (
	"context"

	authID "github.com/cheeeasy2501/auth-id/gen/authorization"
	"google.golang.org/grpc"
)

type Services struct {
	Auth *AuthService
}

func NewServices(conn grpc.ClientConnInterface) *Services {
	return &Services{
		Auth: NewAuthService(conn),
	}
}

type AuthService struct {
	conn   *grpc.ClientConn
	client authID.AuthorizationServiceClient
}

func NewAuthService(conn *grpc.ClientConn) *AuthService {
	return &AuthService{
		conn: conn,
	}
}

func (s *AuthService) GetUserInformation(ctx context.Context, id uint64) {
	response := &authID.GetUserByIdResponse{}
	response, err := s.client.GetUserInformation(
		ctx,
		authID.GetUserByIdRequest{
			UserId: id,
		},
	)
}
