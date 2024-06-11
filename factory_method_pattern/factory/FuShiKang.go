package factory

import "design_pattern/factory_method_pattern/product"

type FuShiKang struct{}

func (f FuShiKang) NewXiaoMi() product.IPhone {
	return product.XiaoMi{Factory: "FuShiKang"}
}

func (f FuShiKang) NewHuaWei() product.IPhone {
	return product.HuaWei{Factory: "FuShiKang"}
}
