package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TeodorStamenov/burgers_api/burger"
)

// const dbsource = "postgresql://postgres:admin@localhost:5432/burgers?sslmode=disable"
// const dbsource = "postgres://nlbfkiinwxjhhf:37ec1e6ac38312ef28cab753ee63ac1527d236fdc612839b4d88dd92a638b2f0@ec2-108-128-104-50.eu-west-1.compute.amazonaws.com:5432/d964uh4q48k8cq"

func main() {
	port := os.Getenv("PORT")
	dbsource := "postgres://nlbfkiinwxjhhf:37ec1e6ac38312ef28cab753ee63ac1527d236fdc612839b4d88dd92a638b2f0@ec2-108-128-104-50.eu-west-1.compute.amazonaws.com:5432/d964uh4q48k8cq"
	if port == "" {
		port = "8080"
		dbsource = "postgresql://postgres:admin@localhost:5432/burgers?sslmode=disable"
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	flag.Parse()
	ctx := context.Background()
	var srv burger.Service
	{
		repository := burger.NewRepo(db, logger)

		srv = burger.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := burger.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port ", port)
		handler := burger.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(":"+port, handler)
		// errs <- http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
