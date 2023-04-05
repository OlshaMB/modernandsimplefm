package main

import (
	"fmt"
	"io/ioutil"
	"github.com/OlshaMB/modernandsimplefm/cmd/config"
	"github.com/OlshaMB/modernandsimplefm/pkg/server"
	"github.com/pelletier/go-toml/v2"
)

func main() {
	file, err := ioutil.ReadFile("config.toml");
	if err!=nil {
		fmt.Println("Config is not accessable")
		return
	}

	if toml.Unmarshal([]byte(file), &config.Cfg)!=nil {
		panic("Config invalid")
	}

	server.Server()
}