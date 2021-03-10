package burger

import "context"

// Service interface
type Service interface {
	CreateBurger(ctx context.Context, name string) (string, error)
	GetBurger(ctx context.Context, id string) (string, error)
}
