package users

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// Ideally we would map struct in to another struct while decoding, so you've better control on which fields to ignore on different APIs
type User struct {
	ID        string         `json:"id"     db:"id"`
	Name      string         `json:"name"   db:"name"`
	Bio       sql.NullString `json:"-"      db:"bio"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
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
