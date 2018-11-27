package main

import (
	"log"
	"generic-op/singletons"
	"generic-op/utils"
	"github.com/valyala/fasthttp"
	"strconv"
	"service"
)

func init() {
	// Init configuration
	service.ConfigInitAndFill("uhe.toml", "local", "request_type")

	// Init IP and port
	ip, err := service.Config.GetString("local", "ip")
	if err != nil {
		log.Fatal(err)
	}
	if ip != "" {
		singletons.IP = ip
	} else {
		singletons.IP, err = utils.GetLocalIP()
		if err != nil {
			log.Fatal(err)
		}
	}
	port, err := service.Config.GetInt("local", "port")
	if err != nil {
		log.Fatal(err)
	}
	singletons.Port = port
}

func main() {
	err := fasthttp.ListenAndServe(singletons.IP+":"+strconv.Itoa(singletons.Port), service.HttpHandle)
	if err != nil {
		log.Fatal(err)
	}
}
