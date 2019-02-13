package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App define application structure
type App struct {
	DB     *sql.DB
	Router *mux.Router
}

// Initialize setup database and router
func (app *App) Initialize(host string, port int, user, password, databaseName string) {
	var err error
	// Create data source name
	dsn := fmt.Sprintf(
		"host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		databaseName,
	)
	// Setup database
	app.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	// Ping database
	if err = app.DB.Ping(); err != nil {
		panic(err)
	}
	// Defer database
	defer app.DB.Close()
	// Setup router
	app.Router = mux.NewRouter()
}

// Run setup and serve web server
func (app *App) Run(address string) {
	// Create web server config
	server := &http.Server{
		Addr: address,
	}
	// Setup file server
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Serve
	log.Printf("Gootstrap server is started at %v", server.Addr)
	panic(server.ListenAndServe())
}
