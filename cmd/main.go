package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/dimiro1/tudu/internal/api"
	"github.com/dimiro1/tudu/internal/config"
	"github.com/dimiro1/tudu/internal/localstorage"
	"github.com/dimiro1/tudu/internal/logging"
	"github.com/dimiro1/tudu/internal/toolkit/graceful"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		// Log is not configured
		// Lets just call the standard panic function
		panic(err)
	}

	var (
		store   = localstorage.NewInMemory()
		logger  = logging.NewLogger()
		router  = mux.NewRouter()
		address = fmt.Sprintf(":%d", cfg.Port)
	)

	{
		httpHandler, err := api.NewGetSingleTodoHandler(store, logger)
		if err != nil {
			logger.PanicCouldNotInstantiateHandler(err, "GetSingleTodoHandler")
		}

		router.Handle("/todos/{id}", httpHandler).Methods("GET")
	}

	{
		httpHandler, err := api.NewGetTodosHandler(store, logger)
		if err != nil {
			logger.PanicCouldNotInstantiateHandler(err, "GetTodosHandler")
		}

		router.Handle("/todos", httpHandler).Methods("GET")
	}

	// This is the only way to safely start a http server
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	logger.ListeningHTTP(address)
	if err := graceful.ListenAndServe(server, 5*time.Second); err != nil {
		logger.FatalListeningHTTP(err, address)
	}
}
