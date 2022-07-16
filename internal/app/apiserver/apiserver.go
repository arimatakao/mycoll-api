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
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		db:     database.NewConnection(context.Background(), config.DBURI),
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
	srv.router.HandleFunc("/api/v1/createlink", srv.handleCreateLinks()).Methods("POST")
	srv.router.HandleFunc("/api/v1/findalinks", srv.handleFindLinks()).Methods("POST")
	srv.router.HandleFunc("/api/v1/updatelinks", srv.handleUpdateLinks()).Methods("PUT")
	srv.router.HandleFunc("/api/v1/deletelinks", srv.handleDeleteLinks()).Methods("DELETE")
	srv.router.HandleFunc("/api/v1/signup", srv.handleSignup()).Methods("POST")
	srv.router.HandleFunc("/api/v1/signin", srv.handleSignin()).Methods("POST")
	srv.router.HandleFunc("/api/v1/refresh", srv.handleRefreshToken()).Methods("POST")
	srv.router.HandleFunc("/api/v1/deleteme", srv.handleDeleteUser()).Methods("DELETE")
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

func (srv *APIServer) handleFindLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Find links")
	}
}

func (srv *APIServer) handleUpdateLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Update links")
	}
}

func (srv *APIServer) handleDeleteLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Delete links")
	}
}

func (srv *APIServer) handleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Singin user")
	}
}

func (srv *APIServer) handleSignin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Singin user")
	}
}

func (srv *APIServer) handleRefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Refresh token for user")
	}
}

func (srv *APIServer) handleDeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Delete user")
	}
}
