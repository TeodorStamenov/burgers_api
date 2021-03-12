package service

import (
	"context"

	"github.com/TeodorStamenov/burgers_api/helpers"
)

// Service interface
type Service interface {
	CreateBurger(ctx context.Context, burgerName string, places []helpers.Place) (string, error)
	GetBurger(ctx context.Context, id string) (string, error)
	GetBurgerRadnom(ctx context.Context) (helpers.Info, error)
	GetBurgerPagination(ctx context.Context, place string, limit int, offset int) ([]helpers.Info, error)
}
