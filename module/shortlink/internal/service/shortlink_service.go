package service

import (
	"flu/infra/common"
	"flu/infra/common/bresp"
	"flu/infra/contract"
	"flu/infra/contract/request"
	"flu/infra/contract/response"
	db2 "flu/module/shortlink/internal/repository/db"
	"flu/module/shortlink/internal/repository/idgen"
	log "github.com/sirupsen/logrus"
	"strings"
)

var ShortLinkService *shortLinkService

type shortLinkService struct {
	contract.ShortLinkService
}

func (s *shortLinkService) Gen(req *request.GenReq) *bresp.BaseRespone[*response.GenResp] {
	// 校验参数
	if req == nil || req.OLink == "" {
		return bresp.OfErr[*response.GenResp](common.PARAM_ERR, "参数输入有误")
	}
	// 检查链接是否是以http或者https开头
	if !strings.HasPrefix(req.OLink, "http://") && !strings.HasPrefix(req.OLink, "https://") {
		return bresp.OfErr[*response.GenResp](common.PARAM_ERR, "链接需要以http://或https://开头")
	}
	// 查询oLink是否存在，存在就直接报错返回
	exist, err := db2.Slink.OLinkExist(req.OLink)
	if err != nil {
		log.Errorf("[ShortLinkService][Gen], 查询DB失败 req: %+v, err: %+v", req, err)
		return bresp.OfErr[*response.GenResp](common.INTERNAL_BUSINESS_ERR, "生成短链接失败")
	}
	if exist {
		log.Warnf("[ShortLinkService][Gen], 链接已存在，不可重复生成 req: %+v", req)
		return bresp.OfErr[*response.GenResp](common.PARAM_ERR, "链接已存在，不可重复生成")
	}
	// oLink存在尝试插入
	sLink := idgen.NextId()
	_, err = db2.Slink.Insert(sLink, req.OLink)
	if err != nil {
		if db2.IsDuplicate(err) {
			for i := 0; i < 3; i++ {
				sLink = idgen.NextId()
				_, err = db2.Slink.Insert(sLink, req.OLink)
				if err == nil || !db2.IsDuplicate(err) {
					break
				}
			}
		}
		if err != nil {
			log.Errorf("[ShortLinkService][Generate], 生成短链接失败, req: %+v, oLink: %s", req, sLink)
			return bresp.OfErr[*response.GenResp](common.INTERNAL_BUSINESS_ERR, "生成短链接失败")
		}
	}
	log.Infof("[ShortLinkService][Generate], 生成短链接成功, req: %+v, sLink: %s\n", req, sLink)
	return bresp.OfSuccess[*response.GenResp](&response.GenResp{
		OLink: req.OLink,
		SLink: sLink,
	})
}

func (s *shortLinkService) Query(req *request.QueryReq) *bresp.BaseRespone[*response.QueryResp] {
	// 检测参数
	if (req.OLink == "" && req.SLink == "") || (req.OLink != "" && req.SLink != "") {
		return bresp.OfErr[*response.QueryResp](common.PARAM_ERR, "参数输入有误")
	}
	// 先查oLink
	if req.OLink != "" {
		sLink, err := db2.Slink.GetSLinkByOLink(req.OLink)
		if err != nil {
			log.Errorf("[ShortLinkService][Query], 查询DB失败 req: %+v, err: %+v", req, err)
			return bresp.OfErr[*response.QueryResp](common.INTERNAL_BUSINESS_ERR, "查询失败")
		}
		return bresp.OfSuccess[*response.QueryResp](&response.QueryResp{
			OLink: req.OLink,
			SLink: sLink,
		})
	}
	// 再查slink
	if req.SLink != "" {
		oLink, err := db2.Slink.GetOLinkBySLink(req.SLink)
		if err != nil {
			log.Errorf("[ShortLinkService][Query], 查询DB失败 req: %+v, err: %+v", req, err)
			return bresp.OfErr[*response.QueryResp](common.INTERNAL_BUSINESS_ERR, "查询失败")
		}
		return bresp.OfSuccess[*response.QueryResp](&response.QueryResp{
			OLink: oLink,
			SLink: req.SLink,
		})
	}
	return bresp.OfErr[*response.QueryResp](common.INTERNAL_BUSINESS_ERR, "内部服务出错")
}
