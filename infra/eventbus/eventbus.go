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
}

func ShopPackagePublishSync(data *eventbus.PackagePublish) error {
	return PackagePublish.PublishSync(SHOP_PACKAGE_PUBLISH, data)
}

func ShopPackageSubscribeSync(handler any) error {
	return PackagePublish.Subscribe(SHOP_PACKAGE_PUBLISH, handler)
}
