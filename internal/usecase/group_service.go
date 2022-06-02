package usecase

import "context"

type GroupService interface {
	Create(ctx context.Context, group *GroupService) error
	Delete(ctx context.Context, groupId int32) error
	AddUser(ctx context.Context, addingUserId, ownerId, groupId int32) error
	RemoveUser(ctx context.Context, deletingUserId, ownerId, groupId int32) error
}
