package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"log/slog"
	"os"

	// Import pgx driver anonymously (with underscore) so it registers itself
	// with database/sql but we don't use its name directly
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zabdiel-bknd/devtracker/internal/config"
	"github.com/zabdiel-bknd/devtracker/internal/database"
	"github.com/zabdiel-bknd/devtracker/internal/handlers"
	"github.com/zabdiel-bknd/devtracker/internal/middlewares"
)

func main() {

	//Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// DB Connection
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("pgx", connStr)
	if err != nil { log.Fatal(err) }
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("Could not connect to DB:", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to Postgres!", "db_name", cfg.DBName)

	// Service Layer (Dependency Injection manual)
	dbService := database.NewService(db)
	projectHandler := handlers.NewProjectHandler(dbService)
	taksHandler := handlers.NewTaskHandler(dbService)
	dashboardHandler := handlers.NewDashboardHandler(dbService)

	// Router
	mux := http.NewServeMux()

	// Middle ware
	wrapperHandler := middlewares.RequestLogger(mux)


	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /projects", projectHandler.Create)

	mux.HandleFunc("GET /projects/{id}", projectHandler.GetById)

	mux.HandleFunc("POST /projects/{id}/tasks", taksHandler.Create)

	mux.HandleFunc("GET /projects/{id}/tasks", taksHandler.List)

	mux.HandleFunc("GET /dashboard", dashboardHandler.GetStats)

	addr := ":" + cfg.ServerPort
	slog.Info("Server running on port", "db_port", cfg.ServerPort)
	
	srv := &http.Server{
		Addr:         addr,
		Handler:      wrapperHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
