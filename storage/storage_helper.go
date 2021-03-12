package storage

import (
	"context"
	"fmt"

	"github.com/TeodorStamenov/burgers_api/helpers"
	"github.com/go-kit/kit/log/level"
	"github.com/lib/pq"
)

////////////////////////////////////////
func (s *storage) insertPlace(ctx context.Context, placeName string) error {
	_, err := s.db.ExecContext(ctx, insertPlaceQuery, placeName)

	if err != nil {
		level.Error(s.logger).Log("err", err.Error())
		return err
	}

	return nil
}

////////////////////////////////////////
func (s *storage) createTable(ctx context.Context, placeName string) error {
	createTable := fmt.Sprintf(createTableQuery, placeName)

	_, err := s.db.ExecContext(ctx, createTable)
	if err, ok := err.(*pq.Error); ok {
		level.Error(s.logger).Log("err", err.Error())
		return err
	}

	return nil
}

////////////////////////////////////////
func (s *storage) insertBurger(ctx context.Context, burgerName string, placeName string, burger helpers.Info) (*pq.Error, bool) {
	insert := fmt.Sprintf(insertBurgerQuery, placeName)

	_, err := s.db.ExecContext(ctx, insert, burgerName, burger.Price, burger.Supply, burger.Rating, burger.Date)

	pqError, ok := err.(*pq.Error)

	return pqError, ok
}

////////////////////////////////////////
func (s *storage) updateBurger(ctx context.Context, burgerName string, placeName string, burger helpers.Info) error {
	update := fmt.Sprintf(updateBurgerQuery, placeName, burger.Price, burger.Supply, burger.Date, burger.Rating, burgerName)

	_, err := s.db.ExecContext(ctx, update)

	if err != nil {
		level.Error(s.logger).Log("err", err.Error())
		return err
	}

	return nil
}
