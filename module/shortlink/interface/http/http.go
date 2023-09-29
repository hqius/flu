package http

import (
	"flu/infra/contract/request"
	"flu/infra/contract/response"
	"flu/infra/http/util"
	"flu/module/shortlink/internal/service"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.POST("/gen", util.Wrapper[request.GenReq, response.GenResp](service.ShortLinkService.Gen))
	r.POST("/query", util.Wrapper[request.QueryReq, response.QueryResp](service.ShortLinkService.Query))
}
