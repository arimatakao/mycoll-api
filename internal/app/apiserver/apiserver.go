package apiserver

import (
	"context"
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
	srv.router.HandleFunc("/api/v1/signup", srv.handleSignup()).Methods("POST")
	srv.router.HandleFunc("/api/v1/signin", srv.handleSignin()).Methods("POST")
	srv.router.HandleFunc("/api/v1/deleteme", srv.handleDeleteUser()).Methods("DELETE")
	srv.router.HandleFunc("/api/v1/createlink", srv.handleCreateLinks()).Methods("POST")
	srv.router.HandleFunc("/api/v1/findlinks", srv.handleFindLinks()).Methods("POST")
	srv.router.HandleFunc("/api/v1/updatelinks", srv.handleUpdateLinks()).Methods("PUT")
	srv.router.HandleFunc("/api/v1/deletelinks", srv.handleDeleteLinks()).Methods("DELETE")
}

func (srv *APIServer) Shutdown() error {
	srv.logger.Warningln("Server shutdown")
	err := srv.db.Shutdown()
	if err != nil {
		srv.logger.Fatalln("Cannot shutdown connection to db: ", err)
		return err
	}
	srv.logger.Warningln("Success shutdown")
	return nil
}
