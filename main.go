package main

import (
	"go-chat-demo/conf"
	"go-chat-demo/router"
	"go-chat-demo/service"
)

func main() {
	conf.Init()
	go service.Manager.Start()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
