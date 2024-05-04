package main

import (
	"embed"
	"go-kitchen-sink/handler"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEnvironment(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	port := os.Getenv("SERVER_PORT")
	slog.Info("Server is running on: ", "port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func initEnvironment() error {
	return godotenv.Load()
}
