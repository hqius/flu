package config

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func initMySQL(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Errorf("connect DB failed, err:%+v\n", err)
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db
}

func InitDB() {
	// 短链DB
	ShortLinkDB = initMySQL("root:12345@tcp(127.0.0.1:3306)/short_link?charset=utf8mb4&parseTime=True")
}
