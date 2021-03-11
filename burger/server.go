package burger

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/TeodorStamenov/burgers_api/helpers"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer function
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
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := helpers.GetVisitor(ip)

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
