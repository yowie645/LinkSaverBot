package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yowie645/ReadItLaterBot/storage"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, p *storage.Page) error {
	q := `INSERT INFO pages (url, user_name) VALUES (?, ?)`

	if _, err := s.db.ExecContext(ctx, q, p.URL, p.UserName); err != nil {
		return fmt.Errorf("can't save page: %w", err)
	}
	return nil
}
