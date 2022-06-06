package adapter

import (
	"gitgub.com/gna69/bot-auth-service/internal/entity"
	"gitgub.com/gna69/bot-auth-service/proto"
)

func ToGroupsProto(groups []*entity.Group) []*proto.Group {
	var protoGroups []*proto.Group
	for _, group := range groups {
		protoGroups = append(protoGroups, &proto.Group{
			Name:    group.Name,
			OwnerId: group.OwnerId,
			Id:      group.Id,
			Members: group.Members,
		})
	}

	return protoGroups
}
