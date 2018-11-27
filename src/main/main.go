package main

import (
	"log"
	"generic-op/codex"
	"fmt"
	"generic-op/utils"
)

var config *codex.TomlConfig

func init() {
	// Init configuration
	config = &codex.TomlConfig{}
	err := config.Init("uhe.toml")
	if err != nil {
		log.Fatal(err)
	}
	err = config.Fill("local")
	if err != nil {
		log.Fatal(err)
	}


}

func main() {
	fmt.Println(utils.GetLocalIP())
}
