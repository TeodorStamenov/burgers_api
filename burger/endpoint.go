package burger

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateBurger endpoint.Endpoint
	GetBurger    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateBurger: makeCreateBurgerEndpoint(s),
		GetBurger:    makeGetBurgerEndpoint(s),
	}
}

func makeCreateBurgerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBurgerRequest)
		ok, err := s.CreateBurger(ctx, req.Name)
		return CreateBurgerResponse{Ok: ok}, err
	}
}

func makeGetBurgerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBurgerRequest)
		name, err := s.GetBurger(ctx, req.ID)
		return GetBurgerResponse{Name: name}, err
	}
}
