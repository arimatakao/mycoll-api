package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/arimatakao/mycoll-api/internal/app/apiserver"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config-path", "configs/config.toml", "Path to config toml file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(cfgPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
