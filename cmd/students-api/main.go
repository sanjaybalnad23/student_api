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
)

func main() {
	// load config
	cfg := config.LoadConfiguration()
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("<h1>Welcome to students api</h1>"))

	})
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

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown gracefully", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")

}
