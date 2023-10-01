package eventbus

import (
	"flu/infra/contract/eventbus"
	bus "github.com/werbenhu/eventbus"
)

const (
	SHOP_PACKAGE_PUBLISH = "/shop/package/publish"
)

var PackagePublish *bus.EventBus

func init() {
	PackagePublish = bus.New()
	PackagePublish.PublishSync(SHOP_PACKAGE_PUBLISH, &eventbus.PackagePublish{
		Data: &eventbus.GoodPk{},
	})
}
