package dbi

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Driver            string `default:"postgres"`
	Host              string `required:"true"`
	User              string `required:"true"`
	Password          string `required:"true"`
	Port              int    `default:"5432"`
	MaxIdleConns      int    `split_words:"true" default:"20"`
	MaxOpenConns      int    `split_words:"true" default:"30"`
	MaxConnLifetimeMs int    `split_words:"true" default:"1000"`
	Name              string `split_words:"true" required:"true"`
	SslMode           string `split_words:"true" default:"disable"`
	AesKey            string `split_words:"true" required:"true"`
}

func (db Config) MaxConnLifetime() time.Duration {
	return time.Millisecond * time.Duration(db.MaxConnLifetimeMs)
}

func (db Config) URL() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", db.User, db.Password, db.Host, db.Port, db.Name, db.SslMode)
}

func NewDB(cfg Config) (*sqlx.DB, error) {
	var err error
	db, err := sqlx.Open(cfg.Driver, cfg.URL())
	if err != nil {
		return nil, fmt.Errorf("error opening conn to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxConnLifetime())
	return db, nil
}
