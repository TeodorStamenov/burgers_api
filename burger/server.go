package burger

import (
	"context"
	"io"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Heroku")
	})

	r.Methods("POST").Path("/create_burger").Handler(httptransport.NewServer(
		endpoints.CreateBurger,
		decodeCreateBurgerRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/burger/{id}").Handler(httptransport.NewServer(
		endpoints.GetBurger,
		decodeGetBurgerRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
