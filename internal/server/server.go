package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/to-the-moshpit/sandpit-turtle/internal/database"
	"github.com/to-the-moshpit/sandpit-turtle/internal/environment"
)

type Server struct {
	port int

	db database.Service
}

func NewServer(env *environment.Env) *http.Server {
	NewServer := &Server{
		port: env.Port,

		db: database.New(env),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
