package adapter

import (
	"gitgub.com/gna69/bot-auth-service/internal/entity"
	"gitgub.com/gna69/bot-auth-service/proto"
)

func ToUserEntity(user *proto.User) *entity.User {
	return &entity.User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.UserName,
		LanguageCode: user.LanguageCode,
		IsBot:        user.IsBot,
		ChatId:       user.ChatId,
	}
}

func ToGroupEntity(group *proto.Group) *entity.Group {
	return &entity.Group{
		Id:      group.Id,
		OwnerId: group.OwnerId,
		Name:    group.Name,
		Members: group.Members,
	}
}
