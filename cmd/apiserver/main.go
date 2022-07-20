package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	go func() {
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := s.Shutdown(); err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}
