package contract

import (
	"flu/infra/common/bresp"
	"flu/infra/contract/request"
	"flu/infra/contract/response"
)

type ShortLinkService interface {
	Gen(req *request.GenReq) *bresp.BaseRespone[response.GenResp]
	Query(req *request.GenReq) *bresp.BaseRespone[response.GenResp]
}
