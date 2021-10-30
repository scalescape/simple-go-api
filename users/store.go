package users

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userStore struct {
	db *sqlx.DB
}

func (s userStore) FetchUsersCount(ctx context.Context) (int, error) {
	query := `select count(*) from users`
	query = s.db.Rebind(query)
	var count int
	err := s.db.Get(&count, query)
	if err != nil {
		return -1, fmt.Errorf("error fetching count: %w", err)
	}
	return count, nil
}

func (s userStore) FetchUsers(ctx context.Context) ([]User, error) {
	query := `select * from users`
	query = s.db.Rebind(query)
	var users []User
	err := s.db.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching count: %w", err)
	}
	return users, nil
}
