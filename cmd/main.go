package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/catninpo/gophi/handlers"
	"github.com/catninpo/gophi/middleware"
	"github.com/catninpo/gophi/postgres"
)

var done = make(chan struct{})

func main() {
	port := flag.Int("port", 8888,
		"specify the port to run the http server on, defaults to 8888")
	flag.Parse()

	db, err := postgres.Open(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("unable to connect to postgres database: %v", err)
	}
	defer db.Close()

	router := registerHandlers(db)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: middleware.Logging(router),

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go interruptListener(&server)

	log.Printf("starting HTTP server on port :%d", *port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("unable to start http server: %v", err)
	}

	<-done
}

// registerHandlers will link the given router to the handlers specified in the
// /handlers directory. Routes can be specified with HTTP verb and path
// variables.
func registerHandlers(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	userHandler := handlers.UserHandler{
		UserService: postgres.NewUserService(db),
	}

	router.HandleFunc("GET /user/{id}", userHandler.GetUser)

	return router
}

// interruptListener will listen for OS interrupt signals and upon receiving an
// interrupt signal, it will gracefully shutdown the HTTP server, waiting for
// idle connections to be finished before closing the done channel to signal
// shutdown completion.
func interruptListener(server *http.Server) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown: %v", err)
	}

	close(done)
}
