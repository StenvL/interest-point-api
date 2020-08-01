package apiserver

import (
	"net/http"

	"github.com/StenvL/interest-points-api/controllers"
	"github.com/StenvL/interest-points-api/middleware"
	"github.com/StenvL/interest-points-api/services"
	"github.com/StenvL/interest-points-api/store"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

// APIServer is a type for creating and configuring server for API.
type APIServer struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

// New creates new API Server.
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

// Start starts API server.
func (s *APIServer) Start() error {
	if err := s.configurateStore(); err != nil {
		return err
	}

	handler := s.configurateRouter()

	return http.ListenAndServe(s.config.BindAddr, handler)
}

func (s *APIServer) configurateStore() error {
	store := store.New(s.config.Store)

	if err := store.Open(); err != nil {
		return err
	}

	s.store = store

	return nil
}

func (s *APIServer) configurateRouter() http.Handler {
	pointService := services.NewPointService(s.store)
	healthService := services.NewHealthService(s.store)
	authService := services.NewAuthenticationService(s.store)

	s.router.HandleFunc("/api/points", controllers.GetPointsByCityHandler(pointService)).Methods("GET")
	s.router.HandleFunc("/api/points/{id}", controllers.GetPointByIDHandler(pointService)).Methods("GET")
	s.router.HandleFunc("/api/points", controllers.CreatePoint(pointService)).Methods("POST")
	s.router.HandleFunc("/api/points/{id}", controllers.EditPoint(pointService)).Methods("PUT")
	s.router.HandleFunc("/api/nearest-points", controllers.GetNearestPointsHandler(pointService)).Methods("GET")
	s.router.HandleFunc("/api/health", controllers.CheckHealth(healthService)).Methods("GET")
	s.router.HandleFunc("/api/authenticate", controllers.Authenticate(authService)).Methods("POST")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins:   s.config.CorsAllowedOrigins,
		AllowedMethods:   s.config.CorsAllowedMethods,
		AllowedHeaders:   s.config.CorsAllowedHeaders,
		AllowCredentials: true,
	})

	s.router.Use(middleware.Authorization)
	return corsOpts.Handler(s.router)
}
