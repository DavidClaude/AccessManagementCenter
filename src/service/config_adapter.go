package service

import (
	"generic-op/codex"
	"log"
)

var ConfigIns *codex.TomlConfig

func ConfigInitAndFill(path string, secs ... string) {
	ConfigIns = &codex.TomlConfig{}
	err := ConfigIns.Init(path)
	if err != nil {
		log.Fatal(err)
	}
	for _,s := range secs{
		err0 := ConfigIns.Fill(s)
		if err0 != nil {
			log.Fatal(err0)
		}
	}
}
