package burger

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService method
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// CreateBurger method
func (s service) CreateBurger(ctx context.Context, name string) (string, error) {
	logger := log.With(s.logger, "method", "CreateBurger")

	genid, _ := uuid.NewV4()
	id := genid.String()
	burger := Burger{
		ID:   id,
		Name: name,
	}

	if err := s.repository.CreateBurger(ctx, burger); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create burger", id)

	return "Success", nil
}

// GetBurger method
func (s service) GetBurger(ctx context.Context, id string) (string, error) {

	return "", nil
}
