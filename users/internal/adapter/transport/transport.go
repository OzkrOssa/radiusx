package transport

import (
	"context"

	usersv1 "github.com/OzkrOssa/radiusx-users/gen/users/v1"
	"github.com/OzkrOssa/radiusx-users/internal/adapter/endpoint"
	"github.com/OzkrOssa/radiusx-users/internal/core/domain"
	gt "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcTransport struct {
	RegisterHandler   gt.Handler
	GetUserHandler    gt.Handler
	ListUsersHandler  gt.Handler
	UpdateUserHandler gt.Handler
	DeleteUserHandler gt.Handler
	usersv1.UnimplementedUserServiceServer
}

func MakeGrpcTransport(endpoint endpoint.Endpoints) usersv1.UserServiceServer {
	return &grpcTransport{
		RegisterHandler:   gt.NewServer(endpoint.RegisterEndopoint, decodeRegisterRequest, encodeRegisterResponse),
		GetUserHandler:    gt.NewServer(endpoint.GetUserEndopoint, decodeGetUserRequest, encodeGetUserResponse),
		ListUsersHandler:  gt.NewServer(endpoint.ListUsersEndopoint, decodeListUsersRequest, encodeListUsersResponse),
		UpdateUserHandler: gt.NewServer(endpoint.UpdateUserEndopoint, decodeUpdateUserRequest, encodeUpdateUserResponse),
	}
}

func (g *grpcTransport) Register(ctx context.Context, request *usersv1.RegisterRequest) (*usersv1.RegisterResponse, error) {

	_, resp, err := g.RegisterHandler.ServeGRPC(ctx, request)
	if err != nil {
		switch err {
		case domain.ErrorConflictData:
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		case domain.ErrorInternal:
			return nil, status.Errorf(codes.Internal, err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.RegisterResponse), nil
}
func (g *grpcTransport) GetUser(ctx context.Context, request *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	_, resp, err := g.GetUserHandler.ServeGRPC(ctx, request)

	if err != nil {
		switch err {
		case domain.ErrorDataNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case domain.ErrorInternal:
			return nil, status.Errorf(codes.Internal, err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.GetUserResponse), nil

}
func (g *grpcTransport) ListUsers(ctx context.Context, request *usersv1.ListUsersRequest) (*usersv1.ListUsersResponse, error) {
	_, resp, err := g.ListUsersHandler.ServeGRPC(ctx, request)

	if err != nil {
		switch err {
		case domain.ErrorInternal:
			return nil, status.Errorf(codes.Internal, err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.ListUsersResponse), nil
}
func (g *grpcTransport) UpdateUser(ctx context.Context, request *usersv1.UpdateUserRequest) (*usersv1.UpdateUserResponse, error) {
	_, resp, err := g.UpdateUserHandler.ServeGRPC(ctx, request)

	if err != nil {
		switch err {
		case domain.ErrorDataNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case domain.ErrorInternal:
			return nil, status.Errorf(codes.Internal, err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return resp.(*usersv1.UpdateUserResponse), nil
}
func (g *grpcTransport) DeleteUser(ctx context.Context, request *usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
