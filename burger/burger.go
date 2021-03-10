package burger

import "context"

// Burger entity
type Burger struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Repository interface
type Repository interface {
	CreateBurger(ctx context.Context, burger Burger) error
	GetBurger(ctx context.Context, id string) (string, error)
}
