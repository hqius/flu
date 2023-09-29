package config

import (
	"github.com/jmoiron/sqlx"
)

var ShortLinkDB *sqlx.DB

func Init() {
	InitDB()
}
