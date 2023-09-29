package db

import (
	"database/sql"
	"flu/module/shortlink/internal/config"
)

type slink struct{}

func (s *slink) OLinkExist(oLink string) (bool, error) {
	var cnt int
	if err := config.ShortLinkDB.Get(&cnt, "select count(*) from tb_slink where o_link = ?", oLink); err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (s *slink) Insert(sLink, oLink string) (sql.Result, error) {
	return config.ShortLinkDB.DB.Exec("insert into tb_slink(s_link, o_link)values(?, ?)", sLink, oLink)
}

func (s *slink) GetSLinkByOLink(oLink string) (string, error) {
	var sLink string
	if err := config.ShortLinkDB.Get(&sLink, "select s_link from tb_slink where o_link = ?", oLink); err != nil {
		return "", err
	}
	return sLink, nil
}

func (s *slink) GetOLinkBySLink(sLink string) (string, error) {
	var oLink string
	if err := config.ShortLinkDB.Get(&oLink, "select o_link from tb_slink where s_link = ?", sLink); err != nil {
		return "", err
	}
	return oLink, nil
}
