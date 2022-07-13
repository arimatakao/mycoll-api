package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	cfg    *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
		router: mux.NewRouter(),
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
}

func (src *APIServer) handleHello() http.HandlerFunc {
	// can add vars witch use 1 time

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "api v1")
	}
}
