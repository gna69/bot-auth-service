package services

import (
	"context"
	"gitgub.com/gna69/bot-auth-service/internal/entity"
	"github.com/jackc/pgx/v4"
)

type GroupService struct {
	conn *pgx.Conn
}

func NewGroupService(conn *pgx.Conn) *GroupService {
	return &GroupService{conn: conn}
}

func (s *GroupService) Create(ctx context.Context, group *entity.Group) error {
	query := `INSERT INTO groups(ownerId, "name", members) values ($1, $2, $3);`
	if _, err := s.conn.Exec(ctx, query, group.OwnerId, group.Name, group.Members); err != nil {
		return err
	}
	return nil
}

func (s *GroupService) Delete(ctx context.Context, groupId int32) error {
	query := `DELETE FROM groups WHERE id = $1;`
	if _, err := s.conn.Exec(ctx, query, groupId); err != nil {
		return err
	}
	return nil
}

func (s *GroupService) AddUser(ctx context.Context, addingUserId, ownerId, groupId int32) error {
	groupMembers, err := s.getGroupMembers(ctx, groupId, ownerId)
	if err != nil {
		return err
	}

	groupMembers = append(groupMembers, addingUserId)

	return s.setGroupMembers(ctx, groupMembers, groupId)
}

func (s *GroupService) RemoveUser(ctx context.Context, deletingUserId, ownerId, groupId int32) error {
	groupMembers, err := s.getGroupMembers(ctx, groupId, ownerId)
	if err != nil {
		return err
	}

	for idx, member := range groupMembers {
		if member == deletingUserId {
			groupMembers[idx] = groupMembers[len(groupMembers)-1]
			groupMembers = groupMembers[:len(groupMembers)-1]
			break
		}
	}

	return s.setGroupMembers(ctx, groupMembers, ownerId)
}

func (s *GroupService) getGroupMembers(ctx context.Context, groupId, ownerId int32) ([]int32, error) {
	query := `SELECT members FROM groups WHERE id = $1 AND ownerId = $2;`
	var groupMembers []int32

	row := s.conn.QueryRow(ctx, query, groupId, ownerId)

	if err := row.Scan(&groupMembers); err != nil {
		return nil, err
	}

	return groupMembers, nil
}

func (s *GroupService) setGroupMembers(ctx context.Context, members []int32, groupId int32) error {
	query := `UPDATE groups SET members = $1 WHERE groupId = $2;`
	if _, err := s.conn.Exec(ctx, query, members, groupId); err != nil {
		return err
	}

	return nil
}
