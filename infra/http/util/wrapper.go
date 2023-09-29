package util

import (
	"encoding/json"
	"flu/infra/common"
	"flu/infra/common/bresp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Wrapper[M any, N any](f func(a *M) *bresp.BaseRespone[*N]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rsp *bresp.BaseRespone[*N]
		req := new(M)
		if err := ctx.ShouldBindJSON(req); err != nil {
			log.Errorf("[flu][http] bind failed, err: %+v", err)
			rsp = bresp.OfErr[*N](common.PARAM_ERR, "请确认参数没有问题")
		} else {
			rsp = f(req)
			reqBytes, _ := json.Marshal(req)
			rspBytes, _ := json.Marshal(rsp)
			log.Infof("method: %s, req: %s, rsp: %s", ctx.Request.RequestURI, string(reqBytes), string(rspBytes))
		}
		ctx.JSON(http.StatusOK, rsp)
	}
}
