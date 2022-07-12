package apiserver

import "github.com/sirupsen/logrus"

type APIServer struct {
	cfg    *Config
	logger *logrus.Logger
}

func New(config *Config) *APIServer {
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server - ", "Port ", s.cfg.BindAddr)

	return nil
}

func (s *APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(s.cfg.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(lvl)

	return nil
}
