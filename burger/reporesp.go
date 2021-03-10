package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateBurgerRequest struct {
		Name string `json:"name"`
	}
	CreateBurgerResponse struct {
		Ok string `json:"ok"`
	}

	GetBurgerRequest struct {
		ID string `json:"id"`
	}
	GetBurgerResponse struct {
		Name string `json:"name"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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
