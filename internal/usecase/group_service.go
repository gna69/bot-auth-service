package usecase

import (
	"context"

	"gitgub.com/gna69/bot-auth-service/internal/entity"
)

type GroupService interface {
	Create(ctx context.Context, group *entity.Group) error
	Delete(ctx context.Context, groupId int32) error
	AddUser(ctx context.Context, addingUserId, ownerId, groupId int32) error
	RemoveUser(ctx context.Context, deletingUserId, ownerId, groupId int32) error
}
