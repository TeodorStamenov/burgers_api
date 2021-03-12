package rest_api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TeodorStamenov/burgers_api/helpers"
	"github.com/gorilla/mux"
)

type (
	CreateBurgerRequest struct {
		Name   string          `json:"burger_name"`
		Places []helpers.Place `json:"places"`
	}
	CreateBurgerResponse struct {
		Ok          string `json:"ok,omitempty"`
		Description string `json:"description,omitempty"`
	}

	GetBurgerRequest struct {
		ID string `json:"id"`
	}
	GetBurgerResponse struct {
		Name string `json:"name"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(CreateBurgerResponse)
	if resp.Description != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeResponseGetRandom(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeResponseGetPagination(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateBurgerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateBurgerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetBurgerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetBurgerRequest
	vars := mux.Vars(r)

	req = GetBurgerRequest{
		ID: vars["id"],
	}

	return req, nil
}

func decodeGetBurgerRandomRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func decodeGetBurgerPaginationRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
