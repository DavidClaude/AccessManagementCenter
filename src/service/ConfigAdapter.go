package service

import (
	"generic-op/codex"
	"log"
)

var Config *codex.TomlConfig

func ConfigInitAndFill(path string, secs ... string) {
	Config = &codex.TomlConfig{}
	err := Config.Init(path)
	if err != nil {
		log.Fatal(err)
	}
	for _,s := range secs{
		err0 := Config.Fill(s)
		if err0 != nil {
			log.Fatal(err0)
		}
	}
}
