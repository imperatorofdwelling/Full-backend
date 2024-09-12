package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	/*
		Idle timeout is a period of time during which
		the server or connection waits for any action from the client.
	*/
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

func LoadConfig() (*Config, error) {
	//TODO load vars from .env
	return &Config{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 5,
	}, nil
}

type Server struct {
	srv *http.Server
}

func New(cfg *Config, h http.Handler) *Server {
	srv := &http.Server{
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		Handler:      h,
	}
	return &Server{
		srv: srv,
	}
}

// NewRouter Creating chi router
func NewRouter(routerGroup chi.Router) http.Handler {
	r := chi.NewRouter()
	// There we need to write endpoints and middlewares

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)

	// We need db instance to work with it
	//TODO routes

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/", routerGroup)
	})

	return r
}

func (s *Server) Run() {
	// Logger print need
	log.Println("Server started...")
	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Disconnect() error {
	return s.srv.Close()
}
