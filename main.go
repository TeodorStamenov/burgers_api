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

const dbsource = "postgresql://postgres:admin@localhost:5432/burgers?sslmode=disable"

func main() {
	// var httpAddr = flag.String("https", ":10433", "http listen address")
	var httpAddr = flag.String("http", ":8080", "http listen address")
	port := os.Getenv("PORT")
	fmt.Println("PORT: ", port)
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
		fmt.Println("listening on port", *httpAddr)
		handler := burger.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
		// errs <- http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
