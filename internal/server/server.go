package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/to-the-moshpit/sandpit-turtle/internal/database"
	"github.com/to-the-moshpit/sandpit-turtle/internal/environment"
	"github.com/to-the-moshpit/sandpit-turtle/internal/logger"
)

type Server struct {
	db  database.Service
	log logger.Logger
}

func NewServer(env *environment.Env) *http.Server {
	NewServer := &Server{
		db:  database.New(env),
		log: logger.New(env),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", env.Port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
