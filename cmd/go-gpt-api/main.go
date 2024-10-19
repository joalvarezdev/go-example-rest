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

	"github.com/joalvarezdev/go-gpt/internal/config"
	"github.com/joalvarezdev/go-gpt/internal/http/handler/product"
	"github.com/joalvarezdev/go-gpt/storage/postgresql"
)

func main() {
  // load config
  cfg := config.MustLoad()

  // database setup
  storage, err := postgresql.New(cfg)
  if err != nil {
    log.Fatal(err)
  }

  slog.Info("storage initialized")

  // setup router
  router := http.NewServeMux()

  router.HandleFunc("POST /products", product.Create(storage))
  router.HandleFunc("GET /products/{id}", product.GetById(storage))
  router.HandleFunc("GET /products", product.GetAll(storage))

  // setup server
  server := http.Server{
    Addr: cfg.Port,
    Handler: router,
  }

  slog.Info("Server started in", slog.String("address", server.Addr))

  done := make(chan os.Signal, 1)

  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  go func() {

    err := server.ListenAndServe()

    if err != nil {
      log.Fatal("failed to start server")
    }
  } ()

  <-done

  slog.Info("shutting down the server")

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

  defer cancel()

  if err := server.Shutdown(ctx); err != nil {
    slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
  }

  slog.Info("server shutdown successfully")
}