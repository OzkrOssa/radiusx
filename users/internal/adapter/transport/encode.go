package transport

import (
	"context"
	"strings"

	usersv1 "github.com/OzkrOssa/radiusx-users/gen/users/v1"
	"github.com/OzkrOssa/radiusx-users/internal/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func encodeRegisterResponse(_ context.Context, request interface{}) (response interface{}, err error) {
	req, ok := request.(*domain.User)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid type from endpoint")
	}

	role := strings.Join([]string{"ROLE", strings.ToUpper(string(req.Role))}, "_")

	registerResponse := &usersv1.RegisterResponse{
		User: &usersv1.User{
			Id:        req.ID,
			Name:      req.Name,
			Email:     req.Email,
			Role:      usersv1.Role(usersv1.Role_value[role]),
			CreatedAt: timestamppb.New(req.CreatedAt),
			UpdatedAt: timestamppb.New(req.UpdatedAt),
		},
	}

	return registerResponse, nil

}

func encodeGetUserResponse(_ context.Context, request interface{}) (response interface{}, err error) {
	req, ok := request.(*domain.User)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid type from endpoint")
	}

	role := strings.Join([]string{"ROLE", strings.ToUpper(string(req.Role))}, "_")

	registerResponse := &usersv1.GetUserResponse{
		User: &usersv1.User{
			Id:        req.ID,
			Name:      req.Name,
			Email:     req.Email,
			Role:      usersv1.Role(usersv1.Role_value[role]),
			CreatedAt: timestamppb.New(req.CreatedAt),
			UpdatedAt: timestamppb.New(req.UpdatedAt),
		},
	}

	return registerResponse, nil

}

func encodeListUsersResponse(_ context.Context, request interface{}) (response interface{}, err error) {
	req, ok := request.([]domain.User)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid type from endpoint")
	}

	var pbUsers []*usersv1.User

	for _, du := range req {
		role := strings.Join([]string{"ROLE", strings.ToUpper(string(du.Role))}, "_")
		u := &usersv1.User{
			Id:        du.ID,
			Name:      du.Email,
			Email:     du.Email,
			Role:      usersv1.Role(usersv1.Role_value[role]),
			CreatedAt: timestamppb.New(du.CreatedAt),
			UpdatedAt: timestamppb.New(du.UpdatedAt),
		}

		pbUsers = append(pbUsers, u)
	}

	registerResponse := &usersv1.ListUsersResponse{
		User: pbUsers,
	}

	return registerResponse, nil

}

func encodeUpdateUserResponse(_ context.Context, request interface{}) (response interface{}, err error) {
	req, ok := request.(*domain.User)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid type from endpoint")
	}

	role := strings.Join([]string{"ROLE", strings.ToUpper(string(req.Role))}, "_")

	updateUserResponse := &usersv1.RegisterResponse{
		User: &usersv1.User{
			Id:        req.ID,
			Name:      req.Name,
			Email:     req.Email,
			Role:      usersv1.Role(usersv1.Role_value[role]),
			CreatedAt: timestamppb.New(req.CreatedAt),
			UpdatedAt: timestamppb.New(req.UpdatedAt),
		},
	}

	return updateUserResponse, nil

}
