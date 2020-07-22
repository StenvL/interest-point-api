package apiserver

import (
	"net/http"

	"github.com/StenvL/interest-points-api/controllers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//APIServer type for creating and configuring server for API.
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//New method to create new API Server
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start method to start API server
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/api/points", controllers.GetAllPointsHandler()).Methods("GET")
	s.router.HandleFunc("/api/points/{id}", controllers.GetPointByIDHandler()).Methods("GET")
	s.router.HandleFunc("/api/points", controllers.CreatePoint()).Methods("POST")
	s.router.HandleFunc("/api/points", controllers.EditPoint()).Methods("PUT")
	s.router.HandleFunc("/api/closest-points", controllers.GetClosestPointsHandler()).Methods("GET")
}
