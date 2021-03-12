package rest_api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/TeodorStamenov/burgers_api/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// Endpoints public struct
type Endpoints struct {
	CreateBurger        endpoint.Endpoint
	GetBurger           endpoint.Endpoint
	GetBurgerRandom     endpoint.Endpoint
	GetBurgerPagination endpoint.Endpoint
}

// MakeEndpoints function
func MakeEndpoints(srvc service.Service) Endpoints {
	return Endpoints{
		CreateBurger:        makeCreateBurgerEndpoint(srvc),
		GetBurger:           makeGetBurgerEndpoint(srvc),
		GetBurgerRandom:     makeGetBurgerRandomEndpoint(srvc),
		GetBurgerPagination: makeGetBurgerPaginationEndpoint(srvc),
	}
}

func makeCreateBurgerEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBurgerRequest)

		if err := CheckCreateBurgerParams(req); err != nil {
			return CreateBurgerResponse{Description: err.Error()}, nil
		}

		ok, err := srvc.CreateBurger(ctx, req.Name, req.Places)
		if err != nil {
			return CreateBurgerResponse{Description: err.Error()}, nil
		}
		return CreateBurgerResponse{Ok: ok}, err
	}
}

func makeGetBurgerEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBurgerRequest)
		name, err := srvc.GetBurger(ctx, req.ID)
		return GetBurgerResponse{Name: name}, err
	}
}

func makeGetBurgerRandomEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		burger, err := srvc.GetBurgerRadnom(ctx)
		return burger, err
	}
}

func makeGetBurgerPaginationEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*http.Request)

		page_str := mux.Vars(req)["page"]
		page_limit_str := mux.Vars(req)["per_page"]
		place := mux.Vars(req)["place"]

		page_limit, err := strconv.Atoi(page_limit_str)
		if err != nil {
			return nil, err
		}

		page, err := strconv.Atoi(page_str)
		if err != nil {
			return nil, err
		}

		burgers, err := srvc.GetBurgerPagination(ctx, place, page_limit, page)
		return burgers, err
	}
}
