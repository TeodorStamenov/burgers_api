package storage

import (
	"context"

	"github.com/TeodorStamenov/burgers_api/helpers"
)

// Storage interface
type Storage interface {
	CreateBurger(ctx context.Context, burgerName string, placeName string, burger helpers.Info) error
	GetBurger(ctx context.Context, id string) (string, error)
	GetBurgerRandom(ctx context.Context) (helpers.Info, error)
	GetBurgerPagination(ctx context.Context, place string, limit int, offset int) ([]helpers.Info, error)
}
