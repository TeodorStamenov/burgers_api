package burger

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

// ErrRepo value
var ErrRepo error = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateBurger(ctx context.Context, burger Burger) error {
	sql := `
		INSERT INTO burgers(id, name)
		VALUES ($1, $2)`

	if burger.Name == "" {
		return ErrRepo
	}

	_, err := repo.db.ExecContext(ctx, sql, burger.ID, burger.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetBurger(ctx context.Context, id string) (string, error) {
	var name string
	err := repo.db.QueryRow("SELECT name FROM burgers WHERE id=$1", id).Scan(&name)
	if err != nil {
		return "", ErrRepo
	}

	return name, nil
}
