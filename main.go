package main

import (
	"net/http"
	r "go-services/app/http"
	"log"
	"fmt"
	"go-services/global/common/server/conf"
	"go-services/global/common/server/app"
)

// @title Go api services
// @description This is a pet project docs, service api provide.
// @contact.name API Support
// @contact.url https://s.a.iri9.cn
// @host 127.0.0.1:7869
func main() {
	//start
	conf.Setup()
	server := &http.Server{
		Addr:			fmt.Sprintf(":%d",app.Server.HttpPort),
		Handler:		r.Routes(),
		ReadTimeout:	app.Server.ReadTimeout,
		WriteTimeout:	app.Server.WriteTimeout,
		MaxHeaderBytes:	1 << 20,
	}
	log.Printf("[info] start http server listening %v ", app.Server.HttpPort)
	server.ListenAndServe()
}
