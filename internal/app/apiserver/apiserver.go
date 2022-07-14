package apiserver

import (
	"context"
	"io"
	"net/http"

	"github.com/arimatakao/mycoll-api/internal/app/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	cfg    *Config
	logger *logrus.Logger
	router *mux.Router
	db     *database.Connection
}

func New(config *Config) *APIServer {
	ctx := context.Background()
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		db:     database.NewDatabase(ctx, config.DBURI),
	}
}

func (srv *APIServer) Start() error {
	if err := srv.configureLogger(); err != nil {
		return err
	}
	srv.logger.Info("Configure routes")
	srv.configureRouter()

	srv.logger.Info("Starting listening port ", srv.cfg.BindAddr)
	return http.ListenAndServe(srv.cfg.BindAddr, srv.router)
}

func (srv *APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(srv.cfg.LogLevel)
	if err != nil {
		return err
	}

	srv.logger.SetLevel(lvl)

	srv.logger.Info("Logger configure is success")
	return nil
}

func (srv *APIServer) configureRouter() {
	srv.router.HandleFunc("/api/v1", srv.handleHello())
	srv.router.HandleFunc("/api/v1/createlink", srv.handleCreateLinks())
	srv.router.HandleFunc("/api/v1/readalllinks", srv.handleReadAllLinks())
	srv.router.HandleFunc("/api/v1/updatealllinks", srv.handleUpdateAllLinks())
	srv.router.HandleFunc("/api/v1/deletealllinks", srv.handleDeleteAllLinks())
}

func (srv *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "api v1")
	}
}

func (srv *APIServer) handleCreateLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Create link")
	}
}

func (srv *APIServer) handleReadAllLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Read all links")
	}
}

func (srv *APIServer) handleUpdateAllLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Update all links")
	}
}

func (srv *APIServer) handleDeleteAllLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Delete all links")
	}
}
