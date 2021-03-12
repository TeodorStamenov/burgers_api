package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/TeodorStamenov/burgers_api/helpers"
	"github.com/go-kit/kit/log"
)

// ErrRepo value
var ErrRepo error = errors.New("Unable to handle Repo Request")

type storage struct {
	db     *sql.DB
	logger log.Logger
}

// NewStorage function
func NewStorage(db *sql.DB, logger log.Logger) Storage {
	return &storage{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (s *storage) CreateBurger(ctx context.Context, burgerName string, placeName string, burger helpers.Info) error {
	if err := s.insertPlace(ctx, placeName); err != nil {
		return err
	}

	if err := s.createTable(ctx, placeName); err != nil {
		return err
	}

	if err, ok := s.insertBurger(ctx, burgerName, placeName, burger); ok {
		// if the row is already in the table update it
		if UNIQUE_VIOLATION_ERR == err.Code.Name() {
			if err := s.updateBurger(ctx, burgerName, placeName, burger); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (s *storage) GetBurger(ctx context.Context, id string) (string, error) {
	var name string
	err := s.db.QueryRow("SELECT name FROM burgers WHERE id=$1", id).Scan(&name)
	if err != nil {
		return "", ErrRepo
	}

	return name, nil
}

func (s *storage) GetBurgerRandom(ctx context.Context) (helpers.Info, error) {
	var place string
	err := s.db.QueryRow(selectRandomPlace).Scan(&place)
	if err != nil {
		return helpers.Info{}, err
	}

	var burger helpers.Info
	selectBurger := fmt.Sprintf(selectRandomBurger, place)

	err = s.db.QueryRow(selectBurger).Scan(&burger.Name, &burger.Price, &burger.Supply, &burger.Date, &burger.Rating)
	if err != nil {
		return helpers.Info{}, err
	}

	return burger, nil
}

func (s *storage) GetBurgerPagination(ctx context.Context, place string, limit int, offset int) ([]helpers.Info, error) {
	page := limit * offset
	selectPagination := fmt.Sprintf(selectPageLimit, place, limit, page)
	rows, err := s.db.Query(selectPagination)
	if err != nil {
		return nil, err
	}

	if rows != nil {
		var burgers []helpers.Info
		for rows.Next() {
			var burger helpers.Info
			if err := rows.Scan(&burger.Name, &burger.Price, &burger.Supply, &burger.Date, &burger.Rating); err != nil {
				return nil, err
			}
			burgers = append(burgers, burger)
		}
		return burgers, nil
	}

	return nil, fmt.Errorf("Empty rows")
}
