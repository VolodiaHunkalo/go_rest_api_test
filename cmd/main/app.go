package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	config2 "rest_api/internal/config"
	"rest_api/internal/user"
	"rest_api/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("router register...")
	router := httprouter.New()
	cfg := config2.GetConfig()
	logger.Info("register user handler...")

	userHandler := user.HandlerImpl(logger)
	userHandler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config2.Config) {
	logger := logging.GetLogger()

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		abs, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		logger.Info("start sock server at " + abs)
		socketPath := path.Join(abs, "app.sock")
		logger.Debug("socket path: " + socketPath)

		listener, listenErr = net.Listen("unix", socketPath)
		if listenErr != nil {
			logger.Fatal(err)
		}
		logger.Infof("server is listening unix socket %s", socketPath)

	} else {
		logger.Info("create tcp socket")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s: %s", cfg.Listen.BindId, cfg.Listen.Port))
		if listenErr != nil {
			logger.Fatal(listenErr)
		}
		logger.Infof("server is listening port %s: %s", cfg.Listen.BindId, cfg.Listen.Port)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
