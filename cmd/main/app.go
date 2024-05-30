package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"rest_api/internal/user"
	"rest_api/pkg/logging"
	"time"
)

func main() {
	log := logging.GetLogger()
	log.Info("router register...")
	router := httprouter.New()

	log.Info("register user handler...")
	userHandler := user.HandlerImpl()
	userHandler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
