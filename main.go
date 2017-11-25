package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/sirupsen/logrus"

	"github.com/JieJhih/golang-imgur/config"
	"github.com/JieJhih/golang-imgur/imgur"
	"github.com/gin-gonic/gin"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.Parse()

	imgur.Conf = config.BuildDefaultPushConf()

	var err error

	if configFile != "" {
		imgur.Conf, err = config.LoadConfYaml(configFile)

		if err != nil {
			log.Printf("Load yaml config file error: '%v'", err)

			return
		}
	}

	server := &http.Server{
		Addr:         ":" + imgur.Conf.Server.Port,
		Handler:      routerEngine(),
		ReadTimeout:  time.Duration(imgur.Conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(imgur.Conf.Server.WriteTimeout) * time.Second,
	}

	if err := startServer(server); err != nil {
		logrus.Fatal(err)
	}
}
func routerEngine() *gin.Engine {

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST(imgur.Conf.API.UploadImage, imgur.UploadImage)

	return r
}

func startServer(s *http.Server) error {
	return gracehttp.Serve(s)
}
