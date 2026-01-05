package main

import (
	"flag"
	"log"

	"t03/internal/app/t03"

	"github.com/BurntSushi/toml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/t03.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := t03.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := t03.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
