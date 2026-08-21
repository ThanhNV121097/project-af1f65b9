package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ThanhNV121097/project-af1f65b9/backend/internal/migrations"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := migrations.Apply(ctx, db); err != nil {
		log.Fatalf("apply migrations: %v", err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("ping database: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz(db))
	mux.HandleFunc("GET /api/greeting", greeting(db))

	server := &http.Server{
		Addr:              ":" + listenPort(),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %v", err)
	}
}

func healthz(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			writeError(w, http.StatusServiceUnavailable, "service_unavailable", "Service is unavailable.")
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	}
}

func greeting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		var text string
		switch err := db.QueryRowContext(ctx, `select text from greetings where id = 1`).Scan(&text); {
		case errors.Is(err, sql.ErrNoRows):
			writeError(w, http.StatusNotFound, "greeting_not_found", "Greeting is not available.")
		case err != nil:
			log.Printf("load greeting: %v", err)
			writeError(w, http.StatusInternalServerError, "internal_error", "Greeting could not be loaded.")
		case text == "":
			writeError(w, http.StatusUnprocessableEntity, "greeting_empty", "Greeting is empty.")
		default:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{"text": text})
		}
	}
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{"code": code, "message": message}})
}

func listenPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	if port := os.Getenv("APP_PORT"); port != "" {
		return port
	}
	return "8080"
}
