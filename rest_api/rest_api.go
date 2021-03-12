package rest_api

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer function
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods(http.MethodGet).Path("/v2").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello burger")
	})

	r.Methods(http.MethodPost).Path("/v2/burger/create_burger").Handler(httptransport.NewServer(
		endpoints.CreateBurger,
		decodeCreateBurgerRequest,
		encodeResponse,
	))

	r.Methods(http.MethodGet).Path("/v2/burger/random").Handler(httptransport.NewServer(
		endpoints.GetBurgerRandom,
		decodeGetBurgerRandomRequest,
		encodeResponseGetRandom,
	))

	r.Methods(http.MethodGet).Path("/v2/burger/{id}").Handler(httptransport.NewServer(
		endpoints.GetBurger,
		decodeGetBurgerRequest,
		encodeResponse,
	))

	r.Methods(http.MethodGet).Path("/v2/burger").
		Queries(
			"page", "{page:[0-9,]+}",
			"per_page", "{per_page:[0-9,]+}",
			"place", "{place}").
		Handler(httptransport.NewServer(
			endpoints.GetBurgerPagination,
			decodeGetBurgerPaginationRequest,
			encodeResponseGetPagination,
		))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := GetVisitor(ip)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("x-ratelimit-limit", strconv.Itoa(limiter.GetLimit()))
		w.Header().Add("x-ratelimit-remaining", strconv.Itoa(limiter.GetRestCalls()))

		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
