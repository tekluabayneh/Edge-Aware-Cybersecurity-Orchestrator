package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
	DBURL := os.Getenv("DBURL")

	if DBURL == "" {
		log.Fatal("failed to get env files")
	}
	pool, err := pgxpool.New(context.Background(), DBURL)
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
		Addr:    ":" + PORT,
		Handler: app.router,
	}

	fmt.Printf("Server is running on post %v", PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to run ", err.Error())
	}

}
