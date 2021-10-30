package users

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type Service struct {
	store userStore
}

func (s *Service) TotalUsers(ctx context.Context) (int, error) {
	cnt, err := s.store.FetchUsersCount(ctx)
	if err != nil {
		return -1, err
	}
	return cnt, nil
}

func (s *Service) ListUsers(ctx context.Context) ([]User, error) {
	users, err := s.store.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewService(db *sqlx.DB) Service {
	return Service{
		store: userStore{db},
	}
}
