package http

import (
	"flu/infra/http/middleware"
	shortLink "flu/shortlink/interface/http"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var serverPort = 25637

func initRoute(r *gin.Engine) {
	shortLink.Route(r.Group("shortlink"))
}

func initMiddleware(r *gin.Engine) {
	r.Use(middleware.LoggerFileMiddleware())
}

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initRoute(r)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("http running failed, err:%+v", err)
			}
		}()
		if err := r.Run(fmt.Sprintf("127.0.0.1:%d", serverPort)); err != nil {
			log.Errorf("http start failed, err:%+v", err)
			panic(err)
		}
	}()
}
