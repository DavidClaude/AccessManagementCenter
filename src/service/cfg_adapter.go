package service

import (
	"generic-op/codex"
	"log"
)

var Config *codex.TomlConfig

func ConfigInitAndFill(path string, secs ... string) {
	Config = &codex.TomlConfig{}
	for _,s := range secs{
		err := Config.Fill(s)
		if err != nil {
			log.Fatal(err)
		}
	}
}
