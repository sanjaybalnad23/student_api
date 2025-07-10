package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Drstrange373/student-api/internal/config"
	"github.com/Drstrange373/student-api/internal/http/handlers/student"
	"github.com/Drstrange373/student-api/internal/storage/database"
)

func main() {
	// load config
	cfg := config.LoadConfiguration()
	// database setup
	storage, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Database initialized", slog.String("env", cfg.Env))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/student", student.New(storage))
	// setup server

	var server = http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	var done = make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("Server running on", slog.String("address", server.Addr))
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server:", err)
		}

	}()

	<-done

	slog.Info("Shutting down the server")

	err = storage.Db.Close()
	if err != nil {
		slog.Error("Error closing database:" + err.Error())
	} else {
		slog.Info("Database closed successfully")
	}

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown gracefully", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
