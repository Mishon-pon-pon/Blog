package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Mishon-pon-pon/Blog/app/serverblog"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/blog.toml", "this is path to config")
}

func main() {
	flag.Parse()
	config := serverblog.NewConfig()
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	serverblog.Start(config)
}
