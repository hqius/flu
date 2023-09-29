package cmd

import (
	"flu/infra/config"
	"flu/infra/http"
	shortlink "flu/module/shortlink/cmd"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	// infra run
	config.Init()
	http.Run()

	// service run
	shortlink.Run()

	// waiting 3s before close
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	time.Sleep(3 * time.Second)
	log.Info("server stopped running...")
}
