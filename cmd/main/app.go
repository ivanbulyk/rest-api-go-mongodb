package main

import (
	"fmt"
	"github.com/ivanbulyk/rest-api-go-mongodb/internal/config"
	"github.com/ivanbulyk/rest-api-go-mongodb/internal/user"
	"github.com/ivanbulyk/rest-api-go-mongodb/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main(){
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	//cfgMongo := cfg.MongoDB
	//mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	//if err != nil {
	//	panic(err)
	//}
	//storage := db.NewStorage(mongoDBClient, "users", logger)



	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config){
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")


		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}


	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}


	logger.Fatal(server.Serve(listener))

}