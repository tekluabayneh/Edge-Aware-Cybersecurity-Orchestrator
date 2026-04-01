package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	router http.Handler
	pool   *pgxpool.Pool
	db     *db.Queries
}

// create db sqlc instance and check load the db to the router
func New() *App {
	//
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	//
	// if host == "" || port == "" || user == "" || password == "" || dbname == "" {
	// 	log.Fatal("failed to get env files")
	// }
	//
	// DBURL := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=migrations",
	// 	user, password, host, port, dbname,
	// )

	DBURL := os.Getenv("DBURL")
	if DBURL == "" {
		log.Fatal("failed to get env files")
	}

	config, err := pgxpool.ParseConfig(DBURL)
	if err != nil {
		log.Fatal("Failed to parse DB config:", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Failed to create connection pool:", err)
	}

	//Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Database connection successful")
	db := db.New(pool)

	app := &App{
		router: router.LoadRouter(db),
		pool:   pool,
		db:     db,
	}

	return app
}

func (app *App) Start() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		json.Marshal(map[string]string{
			"message": "failed to load env file",
		})
	}

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      app.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Server is running on post %v", PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to run ", err.Error())
	}

}
