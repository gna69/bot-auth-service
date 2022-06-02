package services

import (
	"context"
	"gitgub.com/gna69/bot-auth-service/internal/entity"

	"github.com/jackc/pgx/v4"
)

type UserService struct {
	conn *pgx.Conn
}

func NewUserService(conn *pgx.Conn) *UserService {
	return &UserService{conn: conn}
}

func (s *UserService) Add(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, firstName, lastName, userName, langCode, isBot, chatId) VALUES ($1, $2, $3, $4, $5, $6, $7);`
	if _, err := s.conn.Exec(ctx, query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Username,
		user.LanguageCode,
		user.IsBot,
		user.IsBot,
	); err != nil {
		return nil
	}
	return nil
}

func (s *UserService) Get(ctx context.Context, userId int32) (*entity.User, error) {
	var user *entity.User
	query := `SELECT * FROM users WHERE id = $1;`

	row := s.conn.QueryRow(ctx, query, userId)

	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.LanguageCode,
		&user.IsBot,
		&user.ChatId,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetGroups(ctx context.Context, userId int32) ([]*entity.Group, error) {
	query := `SELECT * FROM groups WHERE $1 = ANY(members);`

	rows, err := s.conn.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	groups, err := s.toGroups(rows)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *UserService) toGroups(rows pgx.Rows) ([]*entity.Group, error) {
	var groups []*entity.Group

	for rows.Next() {
		var group entity.Group

		err := rows.Scan(
			&group.Id,
			&group.OwnerId,
			&group.Name,
			&group.Members,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, &group)
	}

	return groups, nil
}
