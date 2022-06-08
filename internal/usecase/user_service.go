package usecase

import (
	"context"

	"gitgub.com/gna69/bot-auth-service/internal/entity"
)

type UserService interface {
	Add(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, userId int32) (*entity.User, error)
	GetMyGroups(ctx context.Context, userId int32) ([]*entity.Group, error)
	GetConsistsGroups(ctx context.Context, firstName string) ([]*entity.Group, error)
}
