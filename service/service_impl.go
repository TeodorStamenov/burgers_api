package service

import (
	"context"

	"github.com/TeodorStamenov/burgers_api/helpers"
	"github.com/TeodorStamenov/burgers_api/storage"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	storage storage.Storage
	logger  log.Logger
}

// NewService method
func NewService(storage storage.Storage, logger log.Logger) Service {
	return service{
		storage: storage,
		logger:  logger,
	}
}

// CreateBurger method
func (srvc service) CreateBurger(ctx context.Context, burgerName string, places []helpers.Place) (string, error) {
	logger := log.With(srvc.logger, "method", "CreateBurger")

	// genid, _ := uuid.NewV4()
	// id := genid.String()

	for _, place := range places {
		if err := srvc.storage.CreateBurger(ctx, burgerName, place.Name, place.BurgerInfo); err != nil {
			level.Error(logger).Log("err", err)
			return "", err
		}

		logger.Log("create burger for ", place.Name)
	}

	return "Success", nil
}

// GetBurger method
func (srvc service) GetBurger(ctx context.Context, id string) (string, error) {
	return "", nil
}

func (srvc service) GetBurgerRadnom(ctx context.Context) (helpers.Info, error) {
	logger := log.With(srvc.logger, "method", "GetBurgerRandom")

	burger, err := srvc.storage.GetBurgerRandom(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return helpers.Info{}, err
	}

	return burger, nil
}

func (srvc service) GetBurgerPagination(ctx context.Context, place string, limit int, offset int) ([]helpers.Info, error) {
	logger := log.With(srvc.logger, "method", "GetBurgerPagination")

	burgers, err := srvc.storage.GetBurgerPagination(ctx, place, limit, offset)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return burgers, nil
}
