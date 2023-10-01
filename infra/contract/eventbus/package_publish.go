package eventbus

import (
	"time"
)

type PkType int

const (
	NormalPkType PkType = iota + 1000
	GoodPkType
	MallPkType
)

type PackagePublish struct {
	PkId int
	PkType
	Data       interface{}
	PubTime    time.Time
	CreateTime time.Time
	UpdateTime time.Time
}

type GoodPk struct {
	Id    int
	Name  string
	Price float64
}

type MallPk struct {
	Id   int
	Name string
	Addr string
}

type PackagePublishService interface {
	PackagePublish(data *PackagePublish) error
}
