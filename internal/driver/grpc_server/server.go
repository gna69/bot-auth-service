package grpc_server

import (
	"context"
	"gitgub.com/gna69/bot-auth-service/internal/adapter"

	"gitgub.com/gna69/bot-auth-service/internal/usecase"
	proto "gitgub.com/gna69/bot-auth-service/proto"
)

type GrpcServer struct {
	proto.UnimplementedAuthServiceServer
	UsersService usecase.UserService
	GroupService usecase.GroupService
}

func NewGrpcServer(us usecase.UserService, gs usecase.GroupService) *GrpcServer {
	return &GrpcServer{
		UsersService: us,
		GroupService: gs,
	}
}

func (s *GrpcServer) AuthUser(ctx context.Context, user *proto.User) (*proto.AuthResponse, error) {
	existed, err := s.UsersService.Get(ctx, user.GetId())
	if err != nil {
		return nil, err
	}

	if existed != nil {
		groups, err := s.UsersService.GetMyGroups(ctx, user.Id)
		if err != nil {
			return nil, err
		}

		groupsWhichUserMember, err := s.UsersService.GetConsistsGroups(ctx, user.FirstName)
		if err != nil {
			return nil, err
		}

		groups = append(groups, groupsWhichUserMember...)

		userGroups := make([]int32, len(groups))
		for idx, group := range groups {
			userGroups[idx] = group.Id
		}

		return &proto.AuthResponse{
			Status:     proto.Status_SUCCESS,
			UserGroups: userGroups,
		}, nil
	}

	err = s.UsersService.Add(ctx, adapter.ToUserEntity(user))
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{Status: proto.Status_SUCCESS}, nil
}

func (s *GrpcServer) GetUserGroups(ctx context.Context, req *proto.GroupsRequest) (*proto.GroupsResponse, error) {
	groups, err := s.UsersService.GetMyGroups(ctx, req.OwnerId)
	if err != nil {
		return nil, err
	}
	return &proto.GroupsResponse{Groups: adapter.ToGroupsProto(groups)}, nil
}

func (s *GrpcServer) CreateGroup(ctx context.Context, group *proto.Group) (*proto.Result, error) {
	err := s.GroupService.Create(ctx, adapter.ToGroupEntity(group))
	if err != nil {
		return nil, err
	}
	return &proto.Result{Status: proto.Status_SUCCESS}, nil
}

func (s *GrpcServer) RemoveGroup(ctx context.Context, group *proto.Group) (*proto.Result, error) {
	err := s.GroupService.Delete(ctx, group.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Result{Status: proto.Status_SUCCESS}, nil
}

func (s *GrpcServer) AddToGroup(ctx context.Context, req *proto.GroupRequest) (*proto.Result, error) {
	err := s.GroupService.AddUser(ctx, req.AddingUser, req.InitiatorId, req.GroupId)
	if err != nil {
		return nil, err
	}
	return &proto.Result{Status: proto.Status_SUCCESS}, nil
}

func (s *GrpcServer) DeleteFromGroup(ctx context.Context, req *proto.GroupRequest) (*proto.Result, error) {
	err := s.GroupService.RemoveUser(ctx, req.AddingUser, req.InitiatorId, req.GroupId)
	if err != nil {
		return nil, err
	}
	return &proto.Result{Status: proto.Status_SUCCESS}, nil
}
