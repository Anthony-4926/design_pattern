package factory

import "design_pattern/factory_method_pattern/product"

type BiYaDi struct{}

func (b BiYaDi) NewXiaoMi() product.IPhone {
	return &product.XiaoMi{Factory: "BiYaDi"}
}

func (b BiYaDi) NewHuaWei() product.IPhone {
	return &product.HuaWei{Factory: "BiYaDi"}
}
