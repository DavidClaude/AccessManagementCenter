package main

import (
	"log"
	"singletons"
	"generic-op/utils"
	"github.com/valyala/fasthttp"
	"strconv"
	"service"
	"fmt"
)

func init() {
	// Init configuration
	service.ConfigInitAndFill("main.toml", "local", "req_tp", "cnt_tp")

	// Init IP and port
	ip, err := service.ConfigIns.GetString("local", "ip")
	if err != nil {
		log.Fatal(err)
	}
	if ip != "" {
		singletons.LocalIP = ip
	} else {
		singletons.LocalIP, err = utils.GetLocalIP()
		if err != nil {
			log.Fatal(err)
		}
	}
	port, err := service.ConfigIns.GetInt("local", "port")
	if err != nil {
		log.Fatal(err)
	}
	singletons.Port = port
}

func main() {
	fmt.Println("AMC server starts...")
	fmt.Printf("IP: %s, port: %d\n", singletons.LocalIP, singletons.Port)
	err := fasthttp.ListenAndServe(singletons.LocalIP+":"+strconv.Itoa(singletons.Port), service.HttpHandle)
	if err != nil {
		log.Fatal(err)
	}
}
