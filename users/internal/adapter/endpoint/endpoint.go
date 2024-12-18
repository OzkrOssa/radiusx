package endpoint

import (
	"context"

	usersv1 "github.com/OzkrOssa/radiusx-users/gen/users/v1"
	"github.com/OzkrOssa/radiusx-users/internal/core/domain"
	"github.com/OzkrOssa/radiusx-users/internal/core/port"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	RegisterEndopoint   endpoint.Endpoint
	GetUserEndopoint    endpoint.Endpoint
	ListUsersEndopoint  endpoint.Endpoint
	UpdateUserEndopoint endpoint.Endpoint
	DeleteEndopoint     endpoint.Endpoint
}

func MakeServerEndpoints(us port.UserService) *Endpoints {
	return &Endpoints{
		RegisterEndopoint:   MakeRegisterEndpoint(us),
		GetUserEndopoint:    MakeGetUserEndopoint(us),
		ListUsersEndopoint:  MakeListUsersEndopoint(us),
		UpdateUserEndopoint: MakeUpdateUserEndopoint(us),
		DeleteEndopoint:     MakeDeleteEndopoint(us),
	}
}

func MakeRegisterEndpoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userReq, ok := request.(*usersv1.RegisterRequest)
		if !ok {
			return nil, err
		}

		user := &domain.User{
			Name:     userReq.Name,
			Email:    userReq.Email,
			Password: userReq.Password,
		}

		userRes, err := us.Register(ctx, user)
		if err != nil {
			return nil, err
		}

		return userRes, nil
	}
}

func MakeGetUserEndopoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(*usersv1.GetUserRequest)
		if !ok {
			return nil, err
		}

		userRes, err := us.GetUser(ctx, user.Id)
		if err != nil {
			return nil, err
		}

		return userRes, err
	}
}

func MakeListUsersEndopoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		listReq, ok := request.(*usersv1.ListUsersRequest)
		if !ok {
			return nil, err
		}

		userRes, err := us.ListUsers(ctx, listReq.Skip, listReq.Limit)
		if err != nil {
			return nil, err
		}

		return userRes, err
	}
}

func MakeUpdateUserEndopoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*usersv1.UpdateUserRequest)
		if !ok {
			return nil, err
		}

		updateUser := &domain.User{
			ID:       req.Id,
			Name:     *req.Name,
			Email:    *req.Email,
			Password: *req.Password,
			Role:     domain.Role(string(req.Role.String())),
		}

		user, err := us.UpdateUser(ctx, updateUser)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func MakeDeleteEndopoint(us port.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*usersv1.DeleteUserRequest)
		if !ok {
			return nil, err
		}

		return nil, us.DeleteUser(ctx, req.Id)
	}
}
